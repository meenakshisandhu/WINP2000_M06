package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// TestMain runs before the tests to set up the environment.
func TestMain(m *testing.M) {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Initialize the DB
	initDB()

	// Run tests
	code := m.Run()

	// Optionally, clean up resources (like closing DB connection)
	defer db.Close()

	// Exit with the result of the tests
	os.Exit(code)
}

// Test currentTimeHandler for successful response
func TestCurrentTimeHandler(t *testing.T) {
	// Create a new request for the /current-time endpoint
	req, err := http.NewRequest("GET", "/current-time", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder to record the response from the handler
	rr := httptest.NewRecorder()

	// Call the currentTimeHandler function with the request and recorder
	handler := http.HandlerFunc(currentTimeHandler)
	handler.ServeHTTP(rr, req)

	// Check that the status code is 200 OK
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check if the response body contains the expected JSON structure
	assert.Contains(t, rr.Body.String(), `"current_time"`)
}

// Test loggedTimesHandler for successful response
func TestLoggedTimesHandler(t *testing.T) {
	// Create a new request for the /logged-times endpoint
	req, err := http.NewRequest("GET", "/logged-times", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder to record the response from the handler
	rr := httptest.NewRecorder()

	// Call the loggedTimesHandler function with the request and recorder
	handler := http.HandlerFunc(loggedTimesHandler)
	handler.ServeHTTP(rr, req)

	// Check that the status code is 200 OK
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check if the response body contains an array of times
	assert.Contains(t, rr.Body.String(), `"id"`)
	assert.Contains(t, rr.Body.String(), `"timestamp"`)
}
