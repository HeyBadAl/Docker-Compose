package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	router := mux.NewRouter()

	db, err := sql.Open("mysql", "example_user:example_password@tcp(mysql-db:3306)/example_database")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users, err := getAllUsers(db)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error retrieving users: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}).Methods("GET")

	log.Println("Server is running on :8080")
	http.ListenAndServe(":8080", router)
}

func getAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, username, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

