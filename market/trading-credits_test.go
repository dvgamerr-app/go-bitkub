package market

import (
	"github.com/stretchr/testify/assert"
)

func (t *BitkubSuite) TestGetTradingCredits() {
	_, err := GetTradingCredits()
	assert.Equal(t.T(), err, nil)
}
