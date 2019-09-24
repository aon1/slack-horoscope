package handlers

import (
	"github.com/aon1/slack-horoscope-bot/handlers/horoscopes"
	"github.com/aon1/slack-horoscope-bot/services"
)

type Handler struct {
	horoscopes.HoroscopeRoutes
}

func New(service services.Services) (*Handler, error) {
	horoscopeHandler, err := horoscopes.New(service)
	if err != nil {
		return nil, err
	}

	handler := &Handler{
		HoroscopeRoutes: horoscopeHandler,
	}

	return handler, nil
}
