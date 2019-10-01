package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/aon1/slack-horoscope/config"
	"github.com/aon1/slack-horoscope/handlers"
	"github.com/aon1/slack-horoscope/server"
	"github.com/aon1/slack-horoscope/services"
	horoscopeService "github.com/aon1/slack-horoscope/services/babi.hefesto.io"
	"github.com/aon1/slack-horoscope/services/redis"
	"github.com/aon1/slack-horoscope/services/restclient"
	"log"
	"net/http"
)

func main() {
	var (
		api  handlers.API
		service services.Services
		restClient restclient.RestClient
	)

	confFile := flag.String("config", "config.yml", "specify the config file to use")
	conf, err := config.New(*confFile)
	if err != nil {
		log.Fatalf("unable to parse configuration file: %v", err)
	}

	service, err = horoscopeService.New(restClient, conf.HoroscopeServices.BabiHefestoIO)
	if err != nil {
		log.Fatalf("unable to start horoscope service: %v", err)
	}

	//redis, err := redis.New(conf)
	//if err != nil {
	//	log.Fatalf("unable to start redis service: %v", err)
	//}

	redis := &redis.Redis{}

	api, err = handlers.New(service, redis)
	if err != nil {
		log.Fatalf("unable to create handlers: %v", err)
	}

	s, err := server.New(api)
	if err != nil {
		log.Fatalf("error initializing server: %v", err)
	}

	srv := &http.Server{
		Addr:      fmt.Sprintf(":%v", conf.Port),
		Handler:   s,
		TLSConfig: &tls.Config{},
	}

	log.Printf("Starting server over http on port: %v", conf.Port)
	log.Fatal(srv.ListenAndServe())
}
