package main

import (
	"log"

	"github.com/begenov/tg-bot-02/config"
	"github.com/begenov/tg-bot-02/internal/app"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatalln(err)
	}

	if err := app.Run(cfg); err != nil {
		log.Fatalln(err)
	}

}
