package server

import (
	"github.com/aon1/slack-horoscope-bot/handlers"
	"github.com/aon1/slack-horoscope-bot/server/wrappers"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

// New creates a new server with routes from the provided api.
func New(api handlers.API) (*Server, error) {
	s := &Server{Router: mux.NewRouter()}

	s.Router.HandleFunc("/daily/{sunsign}", api.GetDailyHoroscope).Methods(http.MethodPost)

	s.Router.Use(wrappers.Log)
	s.Router.Use(wrappers.JSONResponse) // All of our routes should return JSON

	return s, nil
}

func (s *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin) // TODO: Restrict this to proper origins
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	// Lets Gorilla work
	s.Router.ServeHTTP(rw, req)
}
