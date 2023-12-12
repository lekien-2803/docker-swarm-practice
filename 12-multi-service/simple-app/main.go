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

	// http.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
	// 	rows, err := db.Query("SELECT * FROM people")
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	defer rows.Close()

	// 	var people []Person
	// 	for rows.Next() {
	// 		var p Person
	// 		if err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Gender, &p.DateOfBirth); err != nil {
	// 			http.Error(w, err.Error(), http.StatusInternalServerError)
	// 			return
	// 		}
	// 		people = append(people, p)
	// 	}

	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(people)
	// })

	http.HandleFunc("/people", enableCORS(func(w http.ResponseWriter, r *http.Request) {
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
	}))
	

	log.Fatal(http.ListenAndServe(":8088", nil))
}

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Cho phép tất cả các origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Nếu là OPTIONS request, trả về 200 OK mà không cần gọi handler khác
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Gọi handler tiếp theo trong chuỗi
		next(w, r)
	}
}

