package common

import "os"

// IsFile checks that the given path is a file or not
func IsFile(path string) bool {
	f, e := os.Stat(path)
	if e != nil {
		return false
	}
	return !f.IsDir()
}
