package config

import "os"

type SupabaseConfig struct {
	URL        string
	Key        string
	BucketName string
}

func NewSupabaseConfig() *SupabaseConfig {
	return &SupabaseConfig{
		URL:        os.Getenv("SUPABASE_URL"),
		Key:        os.Getenv("SUPABASE_KEY"),
		BucketName: os.Getenv("SUPABASE_BUCKETNAME"),
	}
}
