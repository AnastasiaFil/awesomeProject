package main

import (
	"awasomeProject/internal/domain"
	"awasomeProject/internal/repository/psql"
	"awasomeProject/internal/service"
	"awasomeProject/internal/transport/rest"
	"context"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// init db
	db, err := domain.NewPostgresConnection(domain.ConnectionInfo{
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
	usersRepo := psql.NewUsers(db)
	usersService := service.NewUsers(usersRepo)
	handler := rest.NewHandler(usersService)

	// init & run server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler.InitRouter(),
	}

	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start a goroutine to handle shutdown
	go func() {
		sig := <-sigChan
		log.Printf("Received signal: %v. Shutting down.", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("Server shutdown failed: %v", err)
		} else {
			log.Println("Server shut down gracefully")
		}
	}()

	// Start the server
	log.Println("SERVER STARTED AT", time.Now().Format(time.RFC3339))
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
