package horoscopes

import "net/http"

type HoroscopeRoutes interface {
	GetHoroscope(w http.ResponseWriter, r *http.Request)
}
