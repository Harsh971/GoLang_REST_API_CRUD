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
)

func main() {
	// ------------------------------------ Load Config
	cfg := config.MustLoad()

	// ------------------------------------ DB Setup

	// ------------------------------------ Setup Router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome..."))
	})

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

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(context)
	if err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")
}
