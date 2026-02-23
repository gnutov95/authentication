package main

import (
	"auth/internal/handlers"
	"auth/internal/repo"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./front"))
	http.Handle("/", fs)
	DB, err := repo.StartBD()
	if err != nil {
		log.Fatal("Error connections to db!")
	}
	log.Println("Success connections to db")
	http.HandleFunc("/registr", handlers.RegistrHandler(DB))
	http.HandleFunc("/login", handlers.AuthHandler(DB))
	http.HandleFunc("/success", handlers.SuccessHandler)

	log.Printf("🚀 %s запущен на http://localhost:%s", "Satl.Courses", "8080")

	log.Fatal(http.ListenAndServe(":"+"8080", nil))
}
