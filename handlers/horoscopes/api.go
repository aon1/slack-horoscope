package horoscopes

import "net/http"

type HoroscopeRoutes interface {
	DailyHoroscope(w http.ResponseWriter, r *http.Request)
}
