package util

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// LoadEnvFile Load environment variables from .env file, and return a map containing containing key-value pairs
func LoadEnvFile(filePath string) (map[string]string, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not determine absolute path: %v", err)
	}
	file, err := os.Open(absPath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	envMap := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// skip empty lines and comment lines
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Split the key-value pairs
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		envMap[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}

	return envMap, nil
}
