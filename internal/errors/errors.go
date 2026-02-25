package errors

import (
	"fmt"
)

// FileNotFoundError represents an error when a file does not exist.
type FileNotFoundError struct {
	Path string
}

func (e FileNotFoundError) Error() string {
	return fmt.Sprintf("file not found: %s", e.Path)
}

// PermissionDeniedError represents an error when file permissions prevent access.
type PermissionDeniedError struct {
	Path string
}

func (e PermissionDeniedError) Error() string {
	return fmt.Sprintf("permission denied: %s", e.Path)
}

// InvalidFileError represents an error when a path is not a valid file (e.g., is a directory).
type InvalidFileError struct {
	Path   string
	Reason string
}

func (e InvalidFileError) Error() string {
	return fmt.Sprintf("invalid file: %s - %s", e.Path, e.Reason)
}

// ParseError represents an error during markdown parsing.
type ParseError struct {
	Message string
}

func (e ParseError) Error() string {
	return fmt.Sprintf("parse error: %s", e.Message)
}

// NewParseError creates a new ParseError with formatted message.
func NewParseError(format string, args ...interface{}) error {
	return ParseError{Message: fmt.Sprintf(format, args...)}
}

// FileTooLargeError represents an error when a file exceeds the size limit.
type FileTooLargeError struct {
	Path     string
	Size     int64
	MaxSize  int64
}

func (e FileTooLargeError) Error() string {
	return fmt.Sprintf("file too large: %s (%d bytes, max %d bytes)", e.Path, e.Size, e.MaxSize)
}
