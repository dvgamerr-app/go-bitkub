package bitkub

import (
	util "github.com/dvgamerr-app/go-bitkub/Util"
	"github.com/rs/zerolog"
)

func init() {
	// Disable zerolog during tests
	zerolog.SetGlobalLevel(zerolog.Disabled)

	util.LoadDotEnv("../.env")
	Initlizer()
}
