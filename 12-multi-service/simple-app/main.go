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
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth"`
}

func main() {

    //Lấy mật khẩu từ secret
	password, err := os.ReadFile("/run/secrets/postgres_password")

    //Kiểm tra xem có mật khẩu hay không
	if err != nil {
		log.Fatal("Không thể đọc secret: ", err)
	}

    //Tạo kết nối tới database
	connStr := fmt.Sprintf("user=postgres dbname=demo sslmode=disable host=postgres password=%s", string(password))

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
