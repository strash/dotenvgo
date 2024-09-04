package dotenvgo

import (
	"log"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	envContent := `
HOST='localhost'
PORT=8080
SUPPORT_EMAIL="test@test.ok"
LUCKY_NUMBER=420
SALT=496d3c88-e4f5-4ff1-8eb8-e4c0a8ac12f2
`
	tmpFile, err := os.CreateTemp("./", ".env_testfile")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	tmpFileName := tmpFile.Name()
	defer os.Remove(tmpFileName)

	if _, err := tmpFile.Write([]byte(envContent)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	if err := Load(&tmpFileName); err != nil {
		log.Fatal(err.Error())
	}

	tests := []struct {
		varName     string
		expectedVal string
	}{
		{"HOST", "localhost"},
		{"PORT", "8080"},
		{"SUPPORT_EMAIL", "test@test.ok"},
		{"LUCKY_NUMBER", "420"},
		{"SALT", "496d3c88-e4f5-4ff1-8eb8-e4c0a8ac12f2"},
	}

	for _, test := range tests {
		val, err := Get(test.varName)
		if err != nil {
			t.Errorf("Error getting variable %s: %v", test.varName, err)
		}
		if val != test.expectedVal {
			t.Errorf("Expected %s for variable %s, got %s", test.expectedVal, test.varName, val)
		}
	}
}

func TestGetNonExistentVariable(t *testing.T) {
	_, err := Get("NON_EXISTENT")
	if err == nil {
		t.Error("Expected an error for non-existent variable, got nil")
	}
}
