package handlers

import "net/http"

type API interface {
	GetHoroscope(w http.ResponseWriter, r *http.Request)
}
