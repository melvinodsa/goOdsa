package utils

import (
	"bufio"
	"os"
)

// WriteFile writes the data to the given file.
func WriteFile(data []byte, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	w.Write(data)
	return w.Flush()
}
