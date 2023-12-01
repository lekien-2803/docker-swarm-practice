package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"

    _ "github.com/lib/pq"
)

type Person struct {
    ID         int    `json:"id"`
    FirstName  string `json:"first_name"`
    LastName   string `json:"last_name"`
    Gender     string `json:"gender"`
    DateOfBirth string `json:"date_of_birth"`
}

func main() {
    connStr := "postgres://postgres:abc123-@postgres:5432/demo?sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
        rows, err := db.Query("SELECT * FROM people")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        var people []Person
        for rows.Next() {
            var p Person
            if err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Gender, &p.DateOfBirth); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            people = append(people, p)
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(people)
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}