package utils

import (
	"log"
	"os"
)

func LogDebug(s string) {
	if os.Getenv("DEBUG") == "true" {
		log.Printf("DEBUG => %s\n", s)
	}
}
