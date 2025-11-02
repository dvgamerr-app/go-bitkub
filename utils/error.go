package utils

import (
	"github.com/rs/zerolog/log"
)

func FatalError(err error) {
	if err != nil {
		log.Fatal().Err(err).Stack().Send()
	}
}
