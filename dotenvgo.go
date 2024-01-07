package dotenvgo

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Env struct {
	Env           string
	Domain        string
	DBHost        string
	DB            string
	DBUser        string
	DBPassword    string
	DBPort        int
	APIHost       string
	AuthHost      string
	AuthClientID  string
	Secret        string
	Salt          string
	Port          int
	EmailSMTP     string
	EmailAddress  string
	EmailUsername string
	EmailPassword string
	EmailPort     int
}

func NewEnv() *Env {
	env := new(Env)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}

	pathToEnv := filepath.Join(dir, ".env")
	file, err := os.ReadFile(pathToEnv)
	if err != nil {
		log.Fatal(err.Error())
	}

	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if len(line) != 0 && !strings.HasPrefix(line, "#") {
			pair := strings.Split(line, "=")
			trimmed := strings.Trim(strings.Trim(pair[1], `"`), `'`)

			switch strings.ToLower(pair[0]) {
			case "env":
				if len(trimmed) != 0 {
					env.Env = trimmed
				} else {
					env.Env = "release"
				}
			case "domain":
				env.Domain = trimmed
			case "db":
				env.DB = trimmed
			case "db_host":
				env.DBHost = trimmed
			case "db_user":
				env.DBUser = trimmed
			case "db_password":
				env.DBPassword = trimmed
			case "port":
				if port, err := strconv.Atoi(trimmed); err == nil {
					env.Port = port
				} else {
					env.Port = 8080
				}
			case "db_port":
				if port, err := strconv.Atoi(trimmed); err == nil {
					env.DBPort = port
				} else {
					env.DBPort = 5432
				}
			case "api_host":
				env.APIHost = trimmed
			case "auth_host":
				env.AuthHost = trimmed
			case "auth_client_id":
				env.AuthClientID = trimmed
			case "secret":
				env.Secret = trimmed
			case "salt":
				env.Salt = trimmed
			case "email_smtp":
				env.EmailSMTP = trimmed
			case "email_address":
				env.EmailAddress = trimmed
			case "email_user":
				env.EmailUsername = trimmed
			case "email_password":
				env.EmailPassword = trimmed
			case "email_port":
				if port, err := strconv.Atoi(trimmed); err == nil {
					env.EmailPort = port
				} else {
					env.EmailPort = 587
				}
			}
		}
	}
	return env
}
