package handlers

import "net/http"

type API interface {
	GetDailyHoroscope(w http.ResponseWriter, r *http.Request)
}
