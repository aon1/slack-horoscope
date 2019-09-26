package horoscopes

import (
	"encoding/json"
	"fmt"
	"github.com/aon1/slack-horoscope/errors"
	"github.com/aon1/slack-horoscope/helpers/http"
	"github.com/aon1/slack-horoscope/models"
	"github.com/aon1/slack-horoscope/services"
	"github.com/aon1/slack-horoscope/services/redis"
	"log"
	"net/http"
	"strings"
	"time"
)

type Handler struct {
	service services.Services
	redis *redis.Redis
}

func New(service services.Services, redis *redis.Redis) (*Handler, error) {
	handler := &Handler{
		service: service,
		redis: redis,
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

	log.Println("Sending 200")
	//we should tell slack asap that we received the request to avoid timeout error
	w.Write(helpers.JSON(helpers.IDResponse{ID:200}))

	now := time.Now().Format("2006-01-02")
	redisKey := fmt.Sprintf("%s:%s:%s", sunsign, period, now)
	val, err := h.redis.Get(redisKey)

	if err != nil {
		helpers.HandleError(w, errors.ResourceNotFoundError, http.StatusInternalServerError)
		return
	}

	//cache miss
	if len(val) == 0 {
		log.Println("cache miss")
		if period == "today" {
			result, err = h.service.GetDailyHoroscope(sunsign)
		} else if period == "week" {
			result, err = h.service.GetWeeklyHoroscope(sunsign)
		}

		if err != nil {
			helpers.HandleError(w, errors.ResourceNotFoundError, http.StatusBadRequest)
			return
		}

		serialized, _ := json.Marshal(&result)
		key := fmt.Sprintf("%s:%s:%s", sunsign, period, now)
		_, err = h.redis.Set(key, string(serialized))
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		log.Println("cache hit")
		err = json.Unmarshal([]byte(val), &result)
	}

	text := strings.Title(result.Horoscope)

	response = models.SlackResponse{
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
