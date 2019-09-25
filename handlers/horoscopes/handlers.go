package horoscopes

import (
	"github.com/aon1/slack-horoscope/errors"
	"github.com/aon1/slack-horoscope/helpers/http"
	"github.com/aon1/slack-horoscope/models"
	"github.com/aon1/slack-horoscope/services"
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

func (h *Handler) GetHoroscope(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	params := strings.Split(r.FormValue("text"), " ")

	if len(params) != 2 {
		helpers.HandleError(w, errors.ResourceNotFoundError, http.StatusBadRequest)
		return
	}

	sunsign := params[0]
	period := params[1]

	var (
		err error
		result models.Horoscope
	)

	if period == "today" {
		result, err = h.service.GetDailyHoroscope(sunsign)
	} else if period == "week" {
		result, err = h.service.GetWeeklyHoroscope(sunsign)
	} else {
		helpers.HandleError(w, errors.InvalidPathParamError, http.StatusBadRequest)
		return
	}

	if err != nil {
		helpers.HandleError(w, errors.ResourceNotFoundError, http.StatusBadRequest)
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

	w.Write(helpers.JSON(response))
}
