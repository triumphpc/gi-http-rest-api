package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/triumphpc/go-http-rest-api/internal/app/storage"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *storage.Store
}

// New instance for APIServer
func New(c *Config) *APIServer {
	return &APIServer{
		config: c,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start APIServer
func (s *APIServer) Start() error {
	// Init logger
	if err := s.configureLogger(); err != nil {
		return err
	}

	// Init router
	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Server has started")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// configureLogger ...
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.Loglevel)
	if err != nil {
		return err
	}

	// Set level from config
	s.logger.SetLevel(level)

	return nil
}

// configureStorage ...
func (s *APIServer) configureStore() error {
	stg := storage.New(s.config.DatabaseURL)
	if err := stg.Open(); err != nil {
		return err
	}

	s.store = stg

	return nil
}

// configureRouter ...
func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	// There code for run first before func handler and will be exec only once
	return func(writer http.ResponseWriter, request *http.Request) {
		// ...
		io.WriteString(writer, "Hello")
	}
}
