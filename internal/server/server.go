package server

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Import pprof package to register pprof handlers

	"github.com/gorilla/mux"
	"github.com/pasipiya/go-boilerplate/config"
	"github.com/pasipiya/go-boilerplate/internal/handler"
	"github.com/pasipiya/go-boilerplate/internal/repository"
	"github.com/pasipiya/go-boilerplate/internal/service"
	"github.com/pasipiya/go-boilerplate/profiling"
)

func Start(cfg *config.Config) error {
	repo := repository.NewUserRepository()
	svc := service.NewUserService(repo)
	handler := handler.NewUserHandler(svc)

	router := mux.NewRouter()

	// Start profiling (pprof and trace)
	profiling.StartPProf()
	profiling.StartTrace()

	// Register pprof handlers on a separate goroutine or using default mux
	go func() {
		fmt.Printf("Starting pprof server on :%s\n", cfg.PprofPort)
		log.Fatal(http.ListenAndServe(":"+cfg.PprofPort, nil)) // nil uses the default mux
	}()

	// Set up API routes
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/users", handler.GetUsers).Methods(http.MethodGet)
	api.HandleFunc("/users/{id}", handler.GetUser).Methods(http.MethodGet)
	api.HandleFunc("/users", handler.CreateUser).Methods(http.MethodPost)

	// Start the main server
	fmt.Printf("Server is running on port %s\n", cfg.ServerPort)
	return http.ListenAndServe(":"+cfg.ServerPort, router)
}
