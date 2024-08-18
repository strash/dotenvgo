package dotenvgo

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var vars map[string]string = map[string]string{}

func Get(variable string) *string {
	v, ok := vars[variable]
	if ok {
		return &v
	}
	return nil
}

func Load(path *string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	p := ".env"
	if path != nil {
		p = *path
	}
	pathToEnv := filepath.Join(dir, p)
	file, err := os.ReadFile(pathToEnv)
	if err != nil {
		log.Fatal(err.Error())
	}
	regex := regexp.MustCompile(`^(\w*)=['"]?([^'"]+?)['"]?$`)

	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if len(line) != 0 {
			match := regex.FindStringSubmatch(line)
			if len(match) == 3 {
				vars[match[1]] = string(match[2])
			}
		}
	}
}

