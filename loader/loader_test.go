package loader

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	// Assuming use of testify for assertions
)

func TestCheckEnvFile(t *testing.T) {
	// Test case 1: .env file exists
	file, err := os.Create(".env") // Create the file for testing
	if err != nil {
		fmt.Println(err)
	}
	assert.NoError(t, CheckEnvFile()) // Expect no error

	// Test case 2: .env file does not exist
	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = os.Remove(".env")
	if err != nil {
		fmt.Println(err)
	}
	err = CheckEnvFile()
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
