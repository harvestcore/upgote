package utils

import "os"

// RunningInDocker Returns if the code is running within a Docker container
func RunningInDocker() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}

	return false
}
