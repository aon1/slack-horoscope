package services

import "github.com/aon1/slack-horoscope-bot/models"

type Services interface {
	GetDailyHoroscope(sunsign string) (models.HoroscopeDaily, error)
	GetWeeklyHoroscope(sunsign string) (models.HoroscopeWeek, error)
}