package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Global db variable to hold the database connection
var db *sql.DB

// TimeResponse represents the JSON structure for the API response
type TimeResponse struct {
	CurrentTime string `json:"current_time"`
}

// LoggedTime represents the structure for each logged time in the /logged-times endpoint response
type LoggedTime struct {
	ID        int    `json:"id"`
	Timestamp string `json:"timestamp"`
}

// Initialize the database connection
func initDB() {
	var err error

	// Load environment variables from the .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Retrieve MySQL password from environment variable
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	if mysqlPassword == "" {
		log.Fatal("MYSQL_PASSWORD environment variable not set")
	}

	// Use the environment variable in the DSN string
	//dsn := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/toronto_time", mysqlPassword)
	dsn := fmt.Sprintf("root:%s@tcp(mysql-toronto-time:3307)/toronto_time", mysqlPassword)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Test the database connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	log.Println("Database connected successfully!")
}

// Handler for the /current-time endpoint
func currentTimeHandler(w http.ResponseWriter, r *http.Request) {
	// Load Toronto's timezone
	location, err := time.LoadLocation("America/Toronto")
	if err != nil {
		http.Error(w, "Unable to load Toronto timezone", http.StatusInternalServerError)
		log.Printf("[ERROR] Error loading timezone: %v", err)
		return
	}

	// Get current time in Toronto timezone
	currentTime := time.Now().In(location)

	// Insert the current time into the database
	if err := logTimeToDatabase(currentTime); err != nil {
		http.Error(w, "Failed to log time to database", http.StatusInternalServerError)
		log.Printf("[ERROR] Error logging time to database: %v", err)
		return
	}

	// Create a JSON response
	response := TimeResponse{
		CurrentTime: currentTime.Format("2006-01-02 15:04:05"),
	}

	// Set JSON header and encode the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("[ERROR] Error encoding response: %v", err)
	}
}

// Log the current time to the database
func logTimeToDatabase(currentTime time.Time) error {
	query := "INSERT INTO time_log (timestamp) VALUES (?)"
	_, err := db.Exec(query, currentTime)
	return err
}

// Handler for the /logged-times endpoint
func loggedTimesHandler(w http.ResponseWriter, r *http.Request) {
	// Query the database to get all logged times
	rows, err := db.Query("SELECT id, timestamp FROM time_log ORDER BY timestamp DESC")
	if err != nil {
		http.Error(w, "Failed to retrieve logged times from the database", http.StatusInternalServerError)
		log.Printf("[ERROR] Error retrieving logged times: %v", err)
		return
	}
	defer rows.Close()

	// Create a slice to store the logged times
	var times []LoggedTime

	// Loop through the rows and append them to the slice
	for rows.Next() {
		var id int
		var timestamp string
		if err := rows.Scan(&id, &timestamp); err != nil {
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
			log.Printf("[ERROR] Error scanning row: %v", err)
			return
		}

		// Add each row to the times slice
		times = append(times, LoggedTime{
			ID:        id,
			Timestamp: timestamp,
		})
	}

	// Check for errors after iterating through the rows
	if err := rows.Err(); err != nil {
		http.Error(w, "Error iterating over rows", http.StatusInternalServerError)
		log.Printf("[ERROR] Error iterating over rows: %v", err)
		return
	}

	// Create a JSON response and encode the times array
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(times); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("[ERROR] Error encoding response: %v", err)
	}
}

func main() {
	// Initialize the database
	initDB()
	defer db.Close()

	// Register the /current-time endpoint
	http.HandleFunc("/current-time", currentTimeHandler)

	// Register the /logged-times endpoint
	http.HandleFunc("/logged-times", loggedTimesHandler)

	// Start the server
	log.Println("[INFO] Server is running on http://localhost:8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("[ERROR] Error starting server: %v", err)
	}
}
