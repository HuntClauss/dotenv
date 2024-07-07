package dotenv

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

	parser := newParser(newTokenizer(string(data)).readAll())
	var groupErr error
	for k, v := range parser.parse() {
		if _, ok := os.LookupEnv(k); !ok {
			if err := os.Setenv(k, v); err != nil {
				groupErr = errors.Join(groupErr, err)
			}
		}
	}

	return groupErr
}

func Environ() []KeyValue {
	envs := os.Environ()
	result := make([]KeyValue, len(envs))
	for i, env := range envs {
		result[i] = keyValueFromString(env)
	}

	return result
}

func Get(key string) string {
	return os.Getenv(key)
}

func MustGet(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	panic("missing environment variable: " + key)
}

func ISSet(keys ...string) bool {
	for _, key := range keys {
		if _, ok := os.LookupEnv(key); !ok {
			return false
		}
	}

	return true
}

func GetDefault(key string, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return def
}
