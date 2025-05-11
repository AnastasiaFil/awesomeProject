package main

import (
	_ "awasomeProject/docs"
	"awasomeProject/internal/config"
	"awasomeProject/internal/repository/psql"
	"awasomeProject/internal/service"
	"awasomeProject/pkg/rest"
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	var cfg config.Config
	err = envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal("Error processing envconfig:", err)
	}

	// init db
	db, err := NewPostgresConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// init deps
	usersRepo := psql.NewUsers(db)
	usersService := service.NewUsers(usersRepo)
	handler := rest.NewHandler(usersService)
	router := handler.InitRouter()

	// Add Swagger UI route
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// init & run server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
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
	log.Println("Server starting on :8080")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func NewPostgresConnection(config config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		config.DBHost, config.DBPort, config.DBUsername, config.DBName, config.DBSSLMode, config.DBPassword))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
