package market

import (
	"github.com/stretchr/testify/assert"
)

func (t *BitkubSuite) TestGetMyOpenOrders() {
	_, err := GetMyOpenOrders("btc")
	assert.Equal(t.T(), err, nil)
}
