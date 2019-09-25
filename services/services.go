package services

import "github.com/aon1/slack-horoscope/models"

type Services interface {
	GetDailyHoroscope(sunsign string) (models.Horoscope, error)
	GetWeeklyHoroscope(sunsign string) (models.Horoscope, error)
}
