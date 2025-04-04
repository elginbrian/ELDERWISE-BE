package config

import "os"

type SupabaseConfig struct {
	URL        string
	Key        string
	BucketName string
}

func NewSupabaseConfig() *SupabaseConfig {
	return &SupabaseConfig{
		URL:        os.Getenv("https://ucayizjdmgxgwweshzag.supabase.co"),
		Key:        os.Getenv("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InVjYXlpempkbWd4Z3d3ZXNoemFnIiwicm9sZSI6ImFub24iLCJpYXQiOjE3NDM3NjQwMTgsImV4cCI6MjA1OTM0MDAxOH0.pwFn0khq87rxYRim1lQezFMbot34dSp1xq-8h6XFV0o"),
		BucketName: os.Getenv("elderwise-images"),
	}
}
