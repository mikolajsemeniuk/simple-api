package main

import (
	"ex/configuration"
	"ex/router"
	"log"
)

func main() {
	envConfiguration := configuration.EnvConfiguration{}
	if err := envConfiguration.Configure(); err != nil {
		log.Fatal(err)
	}

	httpRouter := router.HTTPRouter{Configuration: envConfiguration}
	if err := httpRouter.Route(); err != nil {
		log.Fatal(err)
	}
}
