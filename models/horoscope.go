package models

type HoroscopeDaily struct {
	Horoscope string `json:"horoscope"`
	Sunsign	string `json:"sunsign"`
	Date string `json:"date"`
}

type HoroscopeWeek struct {
	Horoscope string `json:"horoscope"`
	Sunsign	string `json:"sunsign"`
	Week string `json:"week"`
}


