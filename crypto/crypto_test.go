package crypto

import (
	"testing"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
	"github.com/dvgamerr-app/go-bitkub/utils"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	utils.LoadDotEnv("../.env")
	bitkub.Initlizer()
}

func TestGetAddresses(t *testing.T) {
	params := Addresses{
		Pagination: Pagination{
			Page:  1,
			Limit: 10,
		},
	}
	result, err := GetAddresses(params)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetAddressesWithFilter(t *testing.T) {
	params := Addresses{
		Pagination: Pagination{
			Page:  1,
			Limit: 5,
		},
		SymbolNetwork: SymbolNetwork{
			Symbol: "KUB",
		},
	}
	result, err := GetAddresses(params)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestCreateAddress(t *testing.T) {
	req := CreateAddressRequest{
		Symbol:  "ETH",
		Network: "ETH",
	}
	result, err := CreateAddress(req)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestValidationCreateAddress(t *testing.T) {
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

func TestGetCoins(t *testing.T) {
	params := Coins{}
	result, err := GetCoins(params)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetCoinsWithSymbol(t *testing.T) {
	params := Coins{
		Symbol: "KUB",
	}
	result, err := GetCoins(params)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetCoinsMultipleSymbols(t *testing.T) {
	symbols := []string{"ETH", "USDT", "BNB"}
	for _, symbol := range symbols {
		params := Coins{
			Symbol: symbol,
		}
		result, err := GetCoins(params)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	}
}

func TestGetCompensations(t *testing.T) {
	params := Compensations{
		Pagination: Pagination{
			Page:  1,
			Limit: 10,
		},
	}
	result, err := GetCompensations(params)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetCompensationsByType(t *testing.T) {
	params := Compensations{
		Pagination: Pagination{
			Page:  1,
			Limit: 5,
		},
		Type: "COMPENSATE",
	}
	result, err := GetCompensations(params)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetCompensationsByStatus(t *testing.T) {
	params := Compensations{
		Pagination: Pagination{
			Page:  1,
			Limit: 5,
		},
		Status: "COMPLETED",
	}
	result, err := GetCompensations(params)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetCompensationsBySymbol(t *testing.T) {
	params := Compensations{
		Pagination: Pagination{
			Page:  1,
			Limit: 5,
		},
		Symbol: "XRP",
	}
	result, err := GetCompensations(params)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetDeposits(t *testing.T) {
	params := Deposits{
		Pagination: Pagination{
			Page:  1,
			Limit: 10,
		},
	}
	result, err := GetDeposits(params)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetDepositsWithFilters(t *testing.T) {
	params := Deposits{
		Pagination: Pagination{
			Page:  1,
			Limit: 5,
		},
		Status: "complete",
	}
	result, err := GetDeposits(params)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetDepositsBySymbol(t *testing.T) {
	params := Deposits{
		Pagination: Pagination{
			Page:  1,
			Limit: 5,
		},
		Symbol: "BTC",
	}
	result, err := GetDeposits(params)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetWithdraws(t *testing.T) {
	params := Withdraws{
		Pagination: Pagination{
			Page:  1,
			Limit: 10,
		},
	}
	result, err := GetWithdraws(params)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetWithdrawsWithFilters(t *testing.T) {
	params := Withdraws{
		Pagination: Pagination{
			Page:  1,
			Limit: 5,
		},
		Status: "complete",
	}
	result, err := GetWithdraws(params)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetWithdrawsBySymbol(t *testing.T) {
	params := Withdraws{
		Pagination: Pagination{
			Page:  1,
			Limit: 5,
		},
		Symbol: "THB",
	}
	result, err := GetWithdraws(params)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestCreateWithdraw(t *testing.T) {
	t.Skip("requires is_withdraw permission and real transaction")
	req := CreateWithdrawRequest{
		Symbol:  "RDNT",
		Amount:  "2.00000000",
		Address: "0xDaCd17d1E77604aaFB6e47F5Ffa1F7E35F83fDa7",
		Network: "ARB",
	}
	result, err := CreateWithdraw(req)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestCreateWithdrawValidation(t *testing.T) {
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
