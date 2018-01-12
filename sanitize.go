// Package sanitize provides a single method, IsSane() which checks whether a path expands to be outside the allowed root directory.
package sanitize

import (
	"fmt"
	"path/filepath"
	"strings"
)

// PathNotAbsoluteError is returned if a path isn't absolute.
type PathNotAbsoluteError struct {
	path string
}

// Err serializes the PathNotAbsoluteError.
func (p PathNotAbsoluteError) Error() string {
	return fmt.Sprintf("the path %v is not absolute", p.path)
}

// IsSane checks whether a user-supplied path expands to be outside the provided absolute rootPath.
// The provided root path has to be an absolute path in order to avoid surprises during expansion.
// The function will return a PathNotAbsoluteError otherwise.
func IsSane(rootPath string, path string) (bool, error) {
	// Check if root path is absolute
	if filepath.IsAbs(rootPath) == false {
		return false, PathNotAbsoluteError{rootPath}
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return false, err
	}
	if strings.HasPrefix(absPath, rootPath) {
		return true, nil
	}
	return false, nil
}
