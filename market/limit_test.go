package market

import (
	"github.com/stretchr/testify/assert"
)

func (t *BitkubSuite) TestGetUserLimits() {
	_, err := GetUserLimits()
	assert.Equal(t.T(), err, nil)
}
