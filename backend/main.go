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

func main() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})

	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name FROM items ORDER BY id")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer rows.Close()
		items := []map[string]interface{}{}
		for rows.Next() {
			var id int
			var name string
			rows.Scan(&id, &name)
			items = append(items, map[string]interface{}{"id": id, "name": name})
		}
		json.NewEncoder(w).Encode(items)
	})

	http.HandleFunc("/api/items/count", func(w http.ResponseWriter, r *http.Request) {
		var count int
		if err := db.QueryRow("SELECT COUNT(*) FROM items").Scan(&count); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprint(w, count)
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), nil))
}
