package main

import (
	"os"

	"git.kill0.net/chill9/beepboop/bot"
)

func main() {
	if err := bot.Run(); err != nil {
		os.Exit(1)
	}
}
