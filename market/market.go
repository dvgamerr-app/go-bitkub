package market

import "github.com/dvgamerr-app/go-bitkub/bitkub"

func placeOrder[T any](path string, req any) (*T, error) {
	var response bitkub.ResponseAPI
	if err := bitkub.FetchSecure("POST", path, req, &response); err != nil {
		return nil, err
	}

	return bitkub.DecodeResult[T](response.Result)
}
