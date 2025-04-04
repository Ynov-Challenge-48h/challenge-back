package managefiles

import (
	"os"
)

// FileExists checks if a file exists in a given directory.
// It constructs the full path using the directory and filename, then checks if the file exists.
// Returns true if the file exists, false otherwise.
func FileExists(directory, filename string) bool {
	path := directory + string(os.PathSeparator) + filename
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// CreateFile creates a file in the specified directory.
// It constructs the full path using the directory and filename, then attempts to create the file.
// Returns an error if the file creation fails.
func CreateFile(directory, filename string) error {
	path := directory + string(os.PathSeparator) + filename
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
