package market

import (
	"github.com/stretchr/testify/assert"
)

func (t *BitkubSuite) TestGetMarketTicker() {
	_, err := GetMarketTicker("btc")
	assert.Equal(t.T(), err, nil)
}
