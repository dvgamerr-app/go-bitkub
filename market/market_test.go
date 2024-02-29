package market

import (
	"log"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
	"github.com/touno-io/go-bitkub/bitkub"
	"github.com/touno-io/go-bitkub/helper"
)

type BitkubSuite struct {
	suite.Suite
}

func (t *BitkubSuite) SetupSuite() {
	log.Println("Setup before all tests")

}

func (t *BitkubSuite) TearDownSuite() {
	log.Println("Tear down after all tests")
}

func (t *BitkubSuite) SetupTest() {
	log.Println("Setup before each test")
}

func (t *BitkubSuite) TearDownTest() {
	log.Println("Tear down after each test")
}

func TestBitkubAPISuite(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.Disabled)

	if err := helper.LoadDotEnv("../.env"); err != nil {
		log.Fatalln(err)
	}

	if err := bitkub.Initlizer("", ""); err != nil {
		log.Fatalln(err)
	}
	suite.Run(t, new(BitkubSuite))
}
