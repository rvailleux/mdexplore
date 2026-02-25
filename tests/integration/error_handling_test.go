package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"mdexplore/internal/errors"
	"mdexplore/internal/parser"
)

func TestErrorHandling_FileNotFound(t *testing.T) {
	p := parser.NewGoldmarkParser()
	_, err := p.Parse("/path/that/does/not/exist.md")
	require.Error(t, err)

	fileNotFoundErr, ok := err.(errors.FileNotFoundError)
	assert.True(t, ok)
	assert.Contains(t, fileNotFoundErr.Error(), "file not found")
}

func TestErrorHandling_Directory(t *testing.T) {
	p := parser.NewGoldmarkParser()
	_, err := p.Parse("../fixtures") // Directory, not a file
	require.Error(t, err)

	invalidFileErr, ok := err.(errors.InvalidFileError)
	assert.True(t, ok)
	assert.Contains(t, invalidFileErr.Error(), "expected a file, got directory")
}

func TestErrorHandling_ErrorTypes(t *testing.T) {
	tests := []struct {
		name          string
		path          string
		expectedError interface{}
	}{
		{
			name:          "file not found",
			path:          "/nonexistent/file.md",
			expectedError: errors.FileNotFoundError{},
		},
		{
			name:          "directory instead of file",
			path:          "../fixtures",
			expectedError: errors.InvalidFileError{},
		},
	}

	p := parser.NewGoldmarkParser()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := p.Parse(tt.path)
			require.Error(t, err)
			assert.IsType(t, tt.expectedError, err)
		})
	}
}
