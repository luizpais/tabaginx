package main

import (
	"github.com/luizpais/tabaginx/internal/controllers"
	"github.com/luizpais/tabaginx/internal/models"
	"gopkg.in/yaml.v2"
	"log"
	"net/url"
	"os"
)

func main() {

	var config models.Config

	data, err := os.ReadFile("tabaginx.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}

	for _, destination := range config.Destinations {
		_, err := url.Parse(destination)
		if err != nil {
			log.Fatalf("Invalid URL: %v", err)
			return
		}
	}

	// Now you can use the config struct in your application
	// For example, print the destinations
	for _, destination := range config.Destinations {
		log.Println(destination)
	}

	c := &controllers.Controller{}
	c.Init(config)

}
