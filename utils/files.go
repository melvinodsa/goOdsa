package utils

import (
	"bufio"
	"fmt"
	"os"
)

/*
This file has got functions required for the file operations
of the application
*/

// ReadFile reads a whole file into memory
// and returns a slice of its lines.
func ReadFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var lines string
	scanner := bufio.NewScanner(file)
	//Reading the file
	for scanner.Scan() {
		lines += scanner.Text() + "\n"
	}
	//Removing the trailing new line
	lines = lines[0 : len(lines)-1]
	return lines, scanner.Err()
}

// WriteFile writes the data to the given file.
func WriteFile(data string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	fmt.Fprintln(w, data)
	return w.Flush()
}
