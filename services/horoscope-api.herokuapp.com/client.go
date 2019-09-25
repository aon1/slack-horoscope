package horoscope_api_herokuapp_com

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
	url := fmt.Sprintf("%s/%s/%s", r.conf.ApiURL, r.conf.DailyEndpoint, sunsign)
	mapResult := r.client.Get(url, nil)

	horoscope := models.Horoscope{
		Horoscope: mapResult["horoscope"],
		Sunsign:   strings.Title(mapResult["sunsign"]),
	}

	return horoscope, nil
}

func (r *restClient) GetWeeklyHoroscope(sunsign string) (models.Horoscope, error) {
	url := fmt.Sprintf("%s/%s/%s", r.conf.ApiURL, r.conf.WeeklyEndpoint, sunsign)
	mapResult := r.client.Get(url, nil)

	horoscope := models.Horoscope{
		Horoscope: mapResult["horoscope"],
		Sunsign:   mapResult["sunsign"],
	}

	return horoscope, nil
}


