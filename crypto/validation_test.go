package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAddressValidationOffline(t *testing.T) {
	req := CreateAddressRequest{
		Network: "KUB",
	}
	_, err := CreateAddress(req)
	assert.NotNil(t, err)

	req = CreateAddressRequest{
		Symbol: "KUB",
	}
	_, err = CreateAddress(req)
	assert.NotNil(t, err)
}

func TestCreateWithdrawValidationOffline(t *testing.T) {
	req := CreateWithdrawRequest{
		Amount:  "2.00000000",
		Address: "0xDaCd17d1E77604aaFB6e47F5Ffa1F7E35F83fDa7",
		Network: "ARB",
	}
	_, err := CreateWithdraw(req)
	assert.NotNil(t, err)

	req = CreateWithdrawRequest{
		Symbol:  "RDNT",
		Address: "0xDaCd17d1E77604aaFB6e47F5Ffa1F7E35F83fDa7",
		Network: "ARB",
	}
	_, err = CreateWithdraw(req)
	assert.NotNil(t, err)

	req = CreateWithdrawRequest{
		Symbol:  "RDNT",
		Amount:  "2.00000000",
		Network: "ARB",
	}
	_, err = CreateWithdraw(req)
	assert.NotNil(t, err)

	req = CreateWithdrawRequest{
		Symbol:  "RDNT",
		Amount:  "2.00000000",
		Address: "0xDaCd17d1E77604aaFB6e47F5Ffa1F7E35F83fDa7",
	}
	_, err = CreateWithdraw(req)
	assert.NotNil(t, err)
}
