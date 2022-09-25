package handler

import (
	"git.kill0.net/chill9/beepboop/config"
)

type Handlers struct {
	config *config.Config
}

func NewHandlers(config *config.Config) *Handlers {
	return &Handlers{
		config: config,
	}
}
