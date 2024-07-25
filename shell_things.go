package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func RunCommand(pwd string, name string, args ...string) (stdout, stderr string, err error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = pwd
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	err = cmd.Run()
	stdout = outBuf.String()
	stderr = errBuf.String()

	return
}

// fileExists checks if a file exists and is not a directory
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CreateDirectoryIfNotExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		// path already exists
		return false, nil
	}
	if os.IsNotExist(err) {
		// create directory
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return false, fmt.Errorf("failed to create directory: %v", err)
		}
		return true, nil
	}
	// Schrodinger: file may or may not exist. See err for details.
	return false, fmt.Errorf("failed to check if directory exists: %v", err)
}
