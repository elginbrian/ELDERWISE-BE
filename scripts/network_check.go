package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	testInternal := flag.Bool("internal", true, "Test internal Docker networking")
	testGmail := flag.Bool("gmail", true, "Test Gmail connectivity")
	testGoogle := flag.Bool("google", true, "Test general internet connectivity")

	flag.Parse()

	if *testInternal {
		testInternalNetworking()
	}

	if *testGmail {
		testGmailConnectivity()
	}

	if *testGoogle {
		testGeneralConnectivity()
	}

	os.Exit(0)
}

func testInternalNetworking() {
	fmt.Println("=== TESTING INTERNAL DOCKER NETWORKING ===")

	fmt.Println("Testing Docker internal DNS resolution...")
	ips, err := net.LookupIP("postgres")
	if err != nil {
		fmt.Printf("❌ Docker DNS lookup for postgres failed: %v\n", err)
	} else {
		fmt.Println("✅ Docker DNS lookup for postgres successful. Found IPs:")
		for _, ip := range ips {
			fmt.Printf("  %s\n", ip.String())
		}
	}

	fmt.Println("\nTesting connection to Postgres...")
	dialer := &net.Dialer{Timeout: 5 * time.Second}
	conn, err := dialer.Dial("tcp", "postgres:5432")
	if err != nil {
		fmt.Printf("❌ Connection to Postgres failed: %v\n", err)
	} else {
		fmt.Println("✅ Connection to Postgres successful!")
		conn.Close()
	}

	fmt.Println("")
}

func testGmailConnectivity() {
	fmt.Println("=== TESTING GMAIL CONNECTIVITY ===")

	fmt.Println("Testing DNS resolution for smtp.gmail.com...")
	ips, err := net.LookupIP("smtp.gmail.com")
	if err != nil {
		fmt.Printf("❌ DNS lookup failed: %v\n", err)
	} else {
		fmt.Println("✅ DNS lookup successful. Found IPs:")
		for _, ip := range ips {
			fmt.Printf("  %s\n", ip.String())
		}
	}

	fmt.Println("\nTesting connection to smtp.gmail.com:587...")
	dialer := &net.Dialer{Timeout: 10 * time.Second}
	conn, err := dialer.Dial("tcp", "smtp.gmail.com:587")
	if err != nil {
		fmt.Printf("❌ Connection to port 587 failed: %v\n", err)
	} else {
		fmt.Println("✅ Connection to port 587 successful!")
		conn.Close()
	}

	fmt.Println("\nTesting connection to smtp.gmail.com:465...")
	conn, err = tls.DialWithDialer(dialer, "tcp", "smtp.gmail.com:465", &tls.Config{
		ServerName: "smtp.gmail.com",
	})
	if err != nil {
		fmt.Printf("❌ TLS connection to port 465 failed: %v\n", err)
	} else {
		fmt.Println("✅ TLS connection to port 465 successful!")
		conn.Close()
	}

	fmt.Println("")
}

func testGeneralConnectivity() {
	fmt.Println("=== TESTING GENERAL INTERNET CONNECTIVITY ===")

	dialer := &net.Dialer{Timeout: 5 * time.Second}

	fmt.Println("Testing connection to Google DNS (8.8.8.8)...")
	conn, err := dialer.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Printf("❌ Connection to Google DNS failed: %v\n", err)
	} else {
		fmt.Println("✅ Connection to Google DNS successful!")
		conn.Close()
	}

	fmt.Println("\nTesting connection to Google web (google.com:443)...")
	conn, err = tls.DialWithDialer(dialer, "tcp", "google.com:443", &tls.Config{
		ServerName: "google.com",
	})
	if err != nil {
		fmt.Printf("❌ Connection to Google web failed: %v\n", err)
	} else {
		fmt.Println("✅ Connection to Google web successful!")
		conn.Close()
	}

	fmt.Println("")
}

