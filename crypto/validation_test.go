package crypto

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryHelpers(t *testing.T) {
	query := url.Values{}
	addPagination(query, Pagination{Page: 2, Limit: 25})
	addDateRange(query, DateRange{CreatedStart: "2026-01-01", CreatedEnd: "2026-01-31"})

	assert.Equal(
		t,
		"/items?created_end=2026-01-31&created_start=2026-01-01&limit=25&page=2",
		pathWithQuery("/items", query),
	)
	assert.Equal(t, "/items", pathWithQuery("/items", url.Values{}))
}

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
