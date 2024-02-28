package market

import (
	"github.com/stretchr/testify/assert"
)

func (t *BitkubSuite) TestGetWallet() {
	wal, err := GetWallet()
	assert.Equal(t.T(), err, nil)
	assert.Greater(t.T(), len(wal), 0)
}
