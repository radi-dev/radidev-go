package util

import (
	"log"
	"os"
	"path/filepath"
)

func GetExecutableDir() string {
	execPath, err := os.Executable()
	if err != nil {
		log.Fatalf("Could not get executable path: %v", err)
	}
	return filepath.Dir(execPath)
}
