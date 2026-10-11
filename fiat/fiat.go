package fiat

import (
	"net/url"
	"strconv"

	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

func fetchPage[T any](path string, page, limit int) (*T, error) {
	query := url.Values{}
	if page > 0 {
		query.Set("p", strconv.Itoa(page))
	}
	if limit > 0 {
		query.Set("lmt", strconv.Itoa(limit))
	}
	if len(query) > 0 {
		path += "?" + query.Encode()
	}

	var response bitkub.ResponseAPI
	if err := bitkub.FetchSecure("POST", path, nil, &response); err != nil {
		return nil, err
	}

	return bitkub.DecodeResult[T](response)
}
