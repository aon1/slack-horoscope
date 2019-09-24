package handlers

import (
	"github.com/aon1/slack-horoscope-bot/config"
	"github.com/aon1/slack-horoscope-bot/handlers/horoscopes"
	"github.com/aon1/slack-horoscope-bot/services/restclient"
)

type Handler struct {
	horoscopes.HoroscopeRoutes
	restClient restclient.RestClient
}

func New(restClient restclient.RestClient, conf config.Config) (*Handler, error) {
	horoscopeHandler, err := horoscopes.New(restClient, conf)
	if err != nil {
		return nil, err
	}

	handler := &Handler{
		HoroscopeRoutes: horoscopeHandler,
		restClient: restClient,
	}

	return handler, nil
}
