package dotenvgo

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const lineRegEx string = `^(\w*)=['"]?([^'"]+?)['"]?$`

var vars map[string]string = map[string]string{}

// Load loads all variables from the specified file and stores them in memory.
// If path is nil, it will look for a '.env' file in the root of your project.
func Load(path *string) error {
	dir, err := os.Getwd()
	if err != nil {
		return errors.Join(fmt.Errorf("dotenvgo:"), err)
	}
	p := ".env"
	if path != nil {
		p = *path
	}
	pathToEnv := filepath.Join(dir, p)
	file, err := os.ReadFile(pathToEnv)
	if err != nil {
		return errors.Join(fmt.Errorf("dotenvgo:"), err)
	}
	regex := regexp.MustCompile(lineRegEx)
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if len(line) != 0 {
			match := regex.FindStringSubmatch(line)
			if len(match) == 3 {
				vars[match[1]] = string(match[2])
			}
		}
	}
	return nil
}

// Get returns the value for the provided variable.
func Get(variable string) (string, error) {
	v, ok := vars[variable]
	if !ok {
		return "", fmt.Errorf("dotenvgo: Variable \"%s\" does not exist", variable)
	}
	return v, nil
}
