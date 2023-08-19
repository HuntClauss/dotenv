package env

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var environments = make(map[string]string, 10)

func LoadEnv(files ...string) error {
	var group error
	for _, file := range files {
		err := loadEnv(environments, file)
		group = errors.Join(group, err)
	}

	return group
}

func loadEnv(result map[string]string, filepath string) error {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		return fmt.Errorf("cannot open '%s' file: %w", filepath, err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("cannot read '%s' content: %w", filepath, err)
	}

	parser := NewParser(NewTokenizer(string(data)).ReadAll())
	for k, v := range parser.Parse() {
		if _, ok := result[k]; !ok {
			result[k] = v
		}
	}

	return nil
}

func Get(key string) string {
	return environments[key]
}

func GetDefault(key string, def string) string {
	if value, ok := environments[key]; ok {
		return value
	}
	return def
}
