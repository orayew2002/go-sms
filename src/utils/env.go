package utils

import "os"

// Get os.Env key
func GetEnv(key string) string {
	return os.Getenv(key)
}

// Get os.Env key with default value
func GetEnvD(key, def string) string {
	val := GetEnv(key)
	if len(val) == 0 {
		return def
	}

	return val
}
