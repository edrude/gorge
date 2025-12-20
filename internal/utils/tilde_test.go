package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExpandTilde(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	require.NoError(t, err, "Could not get user home directory for test setup")

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "No tilde",
			input:    "/foo/bar",
			expected: "/foo/bar",
		},
		{
			name:     "Leading tilde",
			input:    "~/foo/bar",
			expected: homeDir + "/foo/bar",
		},
		{
			name:     "Tilde in middle",
			input:    "/foo/~/bar",
			expected: "/foo/" + homeDir + "/bar",
		},
		{
			name:     "Only tilde",
			input:    "~",
			expected: homeDir,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ExpandTilde(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
