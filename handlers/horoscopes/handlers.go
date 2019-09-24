package horoscopes

import (
	"github.com/JonathonGore/knowledge-base/util/httputil"
	"github.com/aon1/slack-horoscope-bot/config"
	"github.com/aon1/slack-horoscope-bot/services/restclient"
	"net/http"
	"strings"
)

type Handler struct {
	restClient restclient.RestClient
	conf config.Config
}

type ResponseAttachment struct {
	Text string `json:"text"`
}

type Response struct {
	ResponseType string `json:"response_type"`
	Text string `json:"text"`
	Attachments []ResponseAttachment `json:"attachments"`
}

func New(restClient restclient.RestClient, conf config.Config) (*Handler, error) {
	handler := &Handler{
		restClient: restClient,
		conf: conf,
	}

	return handler, nil
}

func (h *Handler) DailyHoroscope(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sunsign := r.FormValue("text")
	url := h.conf.ApiURL + h.conf.DailyEndpoint + sunsign
	result := h.restClient.Get(url, nil)

	text := strings.Title(result["sunsign"])

	response := Response{
		ResponseType: "in_channel",
		Text: text,
		Attachments: []ResponseAttachment{
			{
				Text:result["horoscope"],
			},
		},
	}

	w.Write(httputil.JSON(response))
}
