package hashing_test

import (
	"passgengo/internal/hashing"
	"testing"
)

var input string = "OhHiMark"

func TestHashSHA256(t *testing.T) {
	expectedOutput := "ee6b7ac68c71fd2eb629e7113695b706783b8543c57b75cf3bcc25925281767b"

	output, err := hashing.HashSHA256([]byte(input))
	if err != nil {
		t.Errorf("ERROR => %s", err.Error())
	}

	if output != expectedOutput {
		t.Errorf("wanted: %s, got: %s", expectedOutput, output)
	}
}
