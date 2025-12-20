package utils

import (
	"os"
	"strings"
)

// ExpandTilde replaces ~ with the homedir of the current user
func ExpandTilde(path string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return strings.Replace(path, "~", home, 1), nil
}
