package main

import (
	"database/sql"
	"encoding/json"
	"log"
    "net/http"
    _ "github.com/mattn/go-sqlite3"
)

type TimesheetEntry struct {
    ID int
    UserID int
    StartDateTime string
    Hours float64
}

type User struct {
    ID int
    Email string 
    FirstName string 
    LastName string 
}


var db *sql.DB

func main() {
	var err error

	db, err = sql.Open("sqlite3", "../sqlite-db/job-organize")

	if err != nil {
		log.Fatal(err)
	}

    // GET
    http.HandleFunc("/users", withCORS(getUsers))
    
    // POST
    http.HandleFunc("/user", withCORS(addUser))

    // POST
    http.HandleFunc("/timesheetEntry", withCORS(addTimesheetEntry))

    log.Println("Server running on http:localhost:8080")

    log.Fatal(http.ListenAndServe(":8080", nil))
}

func withCORS(handler http.HandlerFunc) http.HandlerFunc{
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")

        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }

        handler(w, r)
    }
}


func getUsers(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, email, firstName, lastName FROM users")

    if err != nil {
        http.Error(w, "Failed to query users", http.StatusInternalServerError)
        log.Printf("Query error: %v", err)
        return
    }

    defer rows.Close()

    var users []User

    for rows.Next() {
        var user User

        if err := rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName); err != nil {
            http.Error(w, "Failed to scan user", http.StatusInternalServerError)

            log.Printf("Scan error: %v", err)

            return
        }

        users = append(users, user)
    }

    if err := rows.Err(); err != nil {
        http.Error(w, "Row iteration error", http.StatusInternalServerError)

        log.Printf("Rows error: %v", err)

        return
    }

    w.Header().Set("Content-Type", "application/json")

    w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)

}

func addUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "invalid JSON", http.StatusBadRequest)
        return
    }

    _, err := db.Exec(
        "INSERT INTO users (email, firstName, lastName) VALUES (?, ?, ?)",
        user.Email, user.FirstName, user.LastName,
    )

    if err != nil {
        http.Error(w, "failed to insert user", http.StatusInternalServerError)
        log.Printf("Insert error: %v", err)
        return
    }

    w.Header().Set("content-Type", "application/json")

    w.WriteHeader(http.StatusCreated)

    json.NewEncoder(w).Encode(user)
}

func addTimesheetEntry(w http.ResponseWriter, r *http.Request) {
    var timesheetEntry TimesheetEntry
    if err := json.NewDecoder(r.Body).Decode(&timesheetEntry); err != nil {
        http.Error(w, "invalid JSON", http.StatusBadRequest)
        return
    }

    _, err := db.Exec(
        "INSERT INTO TimesheetEntrys (UserID, StartDateTime, Hours) VALUES (?, ?, ?)",
        timesheetEntry.UserID, timesheetEntry.StartDateTime, timesheetEntry.Hours,
    )

    if err != nil {
        http.Error(w, "failed to insert time sheet entry", http.StatusInternalServerError)
        log.Printf("Insert error: %v", err)
        return
    }

    w.Header().Set("content-Type", "application/json")

    w.WriteHeader(http.StatusCreated)

    json.NewEncoder(w).Encode(timesheetEntry)
}
