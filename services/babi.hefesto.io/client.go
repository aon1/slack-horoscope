package babi_hefesto_io

import (
	"fmt"
	"github.com/aon1/slack-horoscope/config"
	"github.com/aon1/slack-horoscope/models"
	"github.com/aon1/slack-horoscope/services/restclient"
	"strings"
)

type restClient struct {
	client restclient.RestClient
	conf config.HoroscopeService
}

func New(client restclient.RestClient, conf config.HoroscopeService) (*restClient, error) {
	return &restClient{
		client: client,
		conf: conf,
	}, nil
}

func (r *restClient) GetDailyHoroscope(sunsign string) (models.Horoscope, error) {
	url := fmt.Sprintf("%s/%s/%s", r.conf.ApiURL, sunsign, r.conf.DailyEndpoint)
	mapResult := r.client.Get(url, nil)

	horoscope := models.Horoscope{
		Horoscope: strings.TrimSpace(mapResult["texto"]),
		Sunsign:   strings.Title(mapResult["signo"]),
	}

	return horoscope, nil
}

func (r *restClient) GetWeeklyHoroscope(sunsign string) (models.Horoscope, error) {
	url := fmt.Sprintf("%s/%s/%s", r.conf.ApiURL, sunsign, r.conf.WeeklyEndpoint)
	fmt.Println(url)
	mapResult := r.client.Get(url, nil)

	horoscope := models.Horoscope{
		Horoscope: strings.TrimSpace(mapResult["texto"]),
		Sunsign:   strings.Title(mapResult["signo"]),
	}

	return horoscope, nil
}

