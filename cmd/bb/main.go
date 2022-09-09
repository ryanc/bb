package main

import (
	"git.kill0.net/chill9/beepboop/bot"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := bot.Run(); err != nil {
		log.Fatal(err)
	}
}
