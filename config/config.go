package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	DefaultAllowPublic      = false
	DefaultCookieName       = "knowledge_base"
	DefaultPublicCookieName = "kb-public"
	DefaultCookieDuration   = 3600 * 24 * 365
	DefaultDBHost           = "0.0.0.0"
	DefaultDBName           = "faq"
	DefaultDBUser           = "postgres"
	DefaultDBPassword       = "postgres"
	DefaultPort             = 3001
)

type DBConfig struct {
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
}

type Config struct {
	Port                 int      `yaml:"port"`
	ApiURL				 string	  `yaml:"apiurl"`
	DailyEndpoint	     string	  `yaml:"dailyEndpoint"`
}

// DefaultConfig builds a Config object using all the default values.
func DefaultConfig() Config {
	return Config{
		Port:             DefaultPort,
	}
}

// New creates a new config from the config file specified in the filename.
func New(filename string) (Config, error) {
	var c Config

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return c, err
	}

	c = DefaultConfig()

	err = yaml.Unmarshal(contents, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}
