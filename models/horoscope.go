package models

import "fmt"

var signs = []string{"aquario", "peixes", "aries", "touro", "gemeos", "cancer", "leao", "virgem", "libra", "escorpiao", "sagitario", "capricornio"}
var periods = []string{"today", "week"}

type Horoscope struct {
	Horoscope string `json:"horoscope"`
	Sunsign	string `json:"sunsign"`
}

type HoroscopeWeek struct {
	Horoscope string `json:"horoscope"`
	Sunsign	string `json:"sunsign"`
}

func contains(s string, list []string) bool {
	for _, item := range list {
		if item == s {
			return true
		}
	}

	return false
}

func Validate(sunsign, period string) error {
	err := ValidateSign(sunsign)
	if err != nil {
		return err
	}

	err = ValidatePeriod(period)
	if err != nil {
		return err
	}

	return nil
}

func ValidateSign(sunsign string) error {
	if contains(sunsign, signs) {
		return nil
	}

	return fmt.Errorf("Sign must be one of the following %v.", signs)
}

func ValidatePeriod(period string) error {
	if contains(period, periods) {
		return nil
	}

	return fmt.Errorf("Period must be one of the following %v.", periods)
}

