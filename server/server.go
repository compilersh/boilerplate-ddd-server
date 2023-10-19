package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/compilersh/boilerplate-ddd-server/user"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	addr string
	srv  *http.Server
}

type ServerOptions struct {
	Timeout time.Duration
}

// Option pattern is great for optional parameters.
type Option func(*ServerOptions)

func WithTimeout(n time.Duration) Option {
	return func(so *ServerOptions) {
		so.Timeout = n
	}
}

func NewServer(addr string, r chi.Router, opts ...Option) *Server {

	// default options
	so := &ServerOptions{
		Timeout: 15 * time.Second,
	}
	// apply options which may or may not override the defaults
	for _, opt := range opts {
		opt(so)
	}

	srv := &http.Server{
		Addr:         addr,
		ReadTimeout:  so.Timeout,
		WriteTimeout: so.Timeout,
		Handler:      r,
	}

	return &Server{
		addr: addr,
		srv:  srv,
	}
}

// ListenAndServe starts the server, listening on the configured address.
// This function is blocking.
func (s *Server) ListenAndServe() error {
	fmt.Println("Server listening on ", s.addr)
	return s.srv.ListenAndServe()
}

// NewRouter returns a new chi.Router with some basic middleware
// We separate router creation from server creation so that we can
// create test servers with a different router.
func NewRouter(userHandler *user.Handler) chi.Router {
	r := chi.NewRouter()

	// Global middleware
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello there!")

		w.Write([]byte("Hello there!"))
		w.WriteHeader(http.StatusOK)
	})

	r.Route("/users", func(r chi.Router) {
		// We can add middleware to specific routes
		r.Use(authMiddleWare)

		r.Post("/", userHandler.CreateUser)
		r.Get("/", userHandler.GetAllUsers)
		r.Get("/{id}", userHandler.GetUser)
	})

	return r
}

func authMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do auth stuff
		if true {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}
