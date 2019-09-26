package handlers

import (
	"github.com/aon1/slack-horoscope/handlers/horoscopes"
	"github.com/aon1/slack-horoscope/services"
	"github.com/aon1/slack-horoscope/services/redis"
)

type Handler struct {
	horoscopes.HoroscopeRoutes
}

func New(service services.Services, redis *redis.Redis) (*Handler, error) {
	horoscopeHandler, err := horoscopes.New(service, redis)
	if err != nil {
		return nil, err
	}

	handler := &Handler{
		HoroscopeRoutes: horoscopeHandler,
	}

	return handler, nil
}
