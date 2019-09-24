package horoscopes

import (
	"github.com/JonathonGore/knowledge-base/errors"
	"github.com/JonathonGore/knowledge-base/util/httputil"
	"github.com/aon1/slack-horoscope-bot/models"
	"github.com/aon1/slack-horoscope-bot/services"
	"net/http"
	"strings"
)

type Handler struct {
	service services.Services
}

func New(service services.Services) (*Handler, error) {
	handler := &Handler{
		service: service,
	}

	return handler, nil
}

func (h *Handler) DailyHoroscope(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sunsign := r.FormValue("text")
	result, err := h.service.GetDailyHoroscope(sunsign)

	if err != nil {
		httputil.HandleError(w, errors.BadIDError, http.StatusBadRequest)
		return
	}

	text := strings.Title(result.Horoscope)

	response := models.SlackResponse{
		ResponseType: "in_channel",
		Text: result.Sunsign,
		Attachments: []models.SlackResponseAttachment{
			{
				Text: text,
			},
		},
	}

	w.Write(httputil.JSON(response))
}
