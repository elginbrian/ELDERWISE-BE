package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// Test if we can resolve Gmail's SMTP server
	fmt.Println("Testing DNS resolution for smtp.gmail.com...")
	ips, err := net.LookupIP("smtp.gmail.com")
	if err != nil {
		fmt.Printf("DNS lookup failed: %v\n", err)
	} else {
		fmt.Println("DNS lookup successful. Found IPs:")
		for _, ip := range ips {
			fmt.Printf("  %s\n", ip.String())
		}
	}

	// Test if we can connect to Gmail's SMTP server on port 587
	fmt.Println("\nTesting connection to smtp.gmail.com:587...")
	dialer := &net.Dialer{Timeout: 10 * time.Second}
	conn, err := dialer.Dial("tcp", "smtp.gmail.com:587")
	if err != nil {
		fmt.Printf("Connection to port 587 failed: %v\n", err)
	} else {
		fmt.Println("Connection to port 587 successful!")
		conn.Close()
	}

	// Test if we can connect to Gmail's SMTP server on port 465
	fmt.Println("\nTesting connection to smtp.gmail.com:465...")
	conn, err = tls.DialWithDialer(dialer, "tcp", "smtp.gmail.com:465", &tls.Config{
		ServerName: "smtp.gmail.com",
	})
	if err != nil {
		fmt.Printf("TLS connection to port 465 failed: %v\n", err)
	} else {
		fmt.Println("TLS connection to port 465 successful!")
		conn.Close()
	}

	// Test if we can reach Google's DNS
	fmt.Println("\nTesting connection to Google DNS (8.8.8.8)...")
	conn, err = dialer.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Printf("Connection to Google DNS failed: %v\n", err)
	} else {
		fmt.Println("Connection to Google DNS successful!")
		conn.Close()
	}

	os.Exit(0)
}
