package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type HoroscopeService struct {
	ApiURL				 string	  `yaml:"apiurl"`
	DailyEndpoint	     string	  `yaml:"dailyEndpoint"`
	WeeklyEndpoint	     string	  `yaml:"weeklyEndpoint"`
}

type HoroscopeServices struct {
	HoroscopeAPIHeroku HoroscopeService `yaml:"horoscope-api-herokuapp-com"`
	BabiHefestoIO HoroscopeService `yaml:"babi-hefesto-io"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}

type Config struct {
	Port                 int      `yaml:"port"`
	Redis 			 	 Redis `yaml:"redis"`
	HoroscopeServices	 HoroscopeServices `yaml:"horoscope-services"`
}

// New creates a new config from the config file specified in the filename.
func New(filename string) (Config, error) {
	var c Config

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return c, err
	}

	err = yaml.Unmarshal(contents, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}
