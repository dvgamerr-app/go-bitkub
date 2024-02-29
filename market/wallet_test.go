package market

import (
	"github.com/stretchr/testify/assert"
)

func (t *BitkubSuite) TestGetWallet() {
	_, err := GetWallet()
	assert.Equal(t.T(), err, nil)
}
