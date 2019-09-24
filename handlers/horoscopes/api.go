package horoscopes

import "net/http"

type HoroscopeRoutes interface {
	GetDailyHoroscope(w http.ResponseWriter, r *http.Request)
}
