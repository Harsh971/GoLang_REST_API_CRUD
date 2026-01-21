package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Harsh971/GoLang_REST_API_CRUD/internal/config"
	"github.com/Harsh971/GoLang_REST_API_CRUD/internal/http/handlers/student"
	"github.com/Harsh971/GoLang_REST_API_CRUD/internal/storage/sqlite"
)

func main() {
	// ------------------------------------ Load Config
	cfg := config.MustLoad()

	// ------------------------------------ DB Setup
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal("Failed to initialize database")
	}
	slog.Info("Storage Initialized successfully", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	// ------------------------------------ Setup Router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))

	// ------------------------------------ Setup HTTP Server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server started successfully :", slog.String("address", cfg.Addr))

	// fmt.Printf("Server Started %s", cfg.Addr)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to Start Server")
		}
	}()

	<-done

	slog.Info("Shutting Down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")
}
