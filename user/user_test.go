package user

import (
	"github.com/dvgamerr-app/go-bitkub/bitkub"
	"github.com/dvgamerr-app/go-bitkub/util"
	"github.com/rs/zerolog"
)

func init() {
	// Disable zerolog during tests
	zerolog.SetGlobalLevel(zerolog.Disabled)

	util.LoadDotEnv("../.env")
	bitkub.Initlizer()
}
