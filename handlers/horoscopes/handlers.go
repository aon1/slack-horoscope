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

	var response models.SlackResponse

	if len(params) != 2 {
		response := models.SlackResponse{
			ResponseType: "ephemeral",
			Text: "Oops! you didn't provide enough parameters. Usage: /horoscopo [signo] [today|week]",
			Attachments: []models.SlackResponseAttachment{},
		}

		w.Write(helpers.JSON(response))

		return
	}

	sunsign := params[0]
	period := params[1]

	var (
		err error
		result models.Horoscope
	)

	err = models.Validate(sunsign, period)
	if err != nil {
		response = models.SlackResponse{
			ResponseType: "ephemeral",
			Text: err.Error(),
			Attachments: []models.SlackResponseAttachment{},
		}

		w.Write(helpers.JSON(response))

		return
	}

	//we should tell slack asap that we received the request to avoid timeout error
	w.WriteHeader(http.StatusOK)

	if err != nil {
		helpers.HandleError(w, errors.ResourceNotFoundError, http.StatusInternalServerError)
		return
	}

	if period == "today" {
		result, err = h.service.GetDailyHoroscope(sunsign)
	} else if period == "week" {
		result, err = h.service.GetWeeklyHoroscope(sunsign)
	}

	response = models.SlackResponse{
		ResponseType: "in_channel",
		Text: result.Sunsign,
		Attachments: []models.SlackResponseAttachment{
			{
				Text: result.Horoscope,
			},
		},
	}

	w.Write(helpers.JSON(response))
}
