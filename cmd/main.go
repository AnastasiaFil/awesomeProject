package main

import (
	"awasomeProject/internal/repository/psql"
	"awasomeProject/internal/service"
	"awasomeProject/internal/transport/rest"
	"awasomeProject/pkg/database"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

func main() {
	// init db
	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "qwerty123",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// init deps
	booksRepo := psql.NewUsers(db)
	booksService := service.NewUsers(booksRepo)
	handler := rest.NewHandler(booksService)

	// init & run server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler.InitRouter(),
	}

	log.Println("SERVER STARTED AT", time.Now().Format(time.RFC3339))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
