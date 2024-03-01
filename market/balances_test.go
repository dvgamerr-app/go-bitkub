package market

import (
	"github.com/stretchr/testify/assert"
)

func (t *BitkubSuite) TestGetBalances() {
	_, err := GetBalances()
	assert.Equal(t.T(), err, nil)
}
