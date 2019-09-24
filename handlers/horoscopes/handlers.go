package horoscopes

import (
	"fmt"
	"github.com/aon1/slack-horoscope-bot/config"
	"github.com/aon1/slack-horoscope-bot/services/restclient"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	restClient restclient.RestClient
	conf config.Config
}

func New(restClient restclient.RestClient, conf config.Config) (*Handler, error) {
	handler := &Handler{
		restClient: restClient,
		conf: conf,
	}

	return handler, nil
}

func (h *Handler) GetDailyHoroscope(w http.ResponseWriter, r *http.Request) {
	sunsign := mux.Vars(r)["sunsign"]
	url := h.conf.ApiURL + h.conf.DailyEndpoint + sunsign
	result := h.restClient.Get(url, nil)

	fmt.Println(url)

	w.Write([]byte(result["horoscope"]))
}
