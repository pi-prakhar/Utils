package loader

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckEnvFile(t *testing.T) {
	// Test case 1: .env file exists
	file, err := os.Create(".env") // Create the file for testing
	if err != nil {
		fmt.Println(err)
	}
	assert.NoError(t, checkEnvFile()) // Expect no error

	// Test case 2: .env file does not exist
	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = os.Remove(".env")
	if err != nil {
		fmt.Println(err)
	}
	err = checkEnvFile()
	assert.Error(t, err, "Expected error when .env file is missing")
}

func TestLoadEnv(t *testing.T) {
	// Test Case 1: .env file exists and loads successfully
	// Create the file for testing
	file, err := os.Create(".env") // Create `.env` file for testing
	if err != nil {
		fmt.Println(err)
	}

	assert.NoError(t, LoadEnv()) // Expect no error

	// Test Case 2: .env file does not exist
	// Remove the file to simulate non-existence
	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = os.Remove(".env")
	if err != nil {
		fmt.Println(err)
	}

	assert.Error(t, LoadEnv())
}

func TestGetValueFromEnv_ExistingKey(t *testing.T) {
	key := "TEST_KEY"
	value := "test_value"
	os.Setenv(key, value)

	result, err := GetValueFromEnv(key)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != value {
		t.Errorf("Expected value %s, got %s", value, result)
	}

	os.Unsetenv(key)
}

func TestGetValueFromEnv_NonExistentKey(t *testing.T) {
	key := "NON_EXISTENT_KEY"

	result, err := GetValueFromEnv(key)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if err.Error() != fmt.Sprintf("key '%s' not found in .env file", key) {
		t.Errorf("Expected error message: %s, got: %s", fmt.Sprintf("key '%s' not found in .env file", key), err.Error())
	}

	if result != "" {
		t.Errorf("Expected empty string, got: %s", result)
	}
}

func TestGetValueFromEnv_EmptyKey(t *testing.T) {
	key := ""

	result, err := GetValueFromEnv(key)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if err.Error() != "key cannot be empty" {
		t.Errorf("Expected error message: %s, got: %s", "key cannot be empty", err.Error())
	}

	if result != "" {
		t.Errorf("Expected empty string, got: %s", result)
	}
}
