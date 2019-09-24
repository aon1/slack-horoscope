package handlers

import "net/http"

type API interface {
	DailyHoroscope(w http.ResponseWriter, r *http.Request)
}
