package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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
    
    // Đọc mật khẩu từ biến môi trường
    password := os.Getenv("POSTGRES_PASSWORD")

    // Kiểm tra xem mật khẩu đã được đặt hay chưa
    if password == "" {
        log.Fatal("Missing POSTGRES_PASSWORD environment variable")
    }

    // Tạo chuỗi kết nối với mật khẩu
    connStr := fmt.Sprintf("user=root dbname=demo sslmode=disable host=postgres password=%s", password)
    


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

    log.Fatal(http.ListenAndServe(":8088", nil))
}
