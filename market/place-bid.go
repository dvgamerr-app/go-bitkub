package market

import (
	"github.com/dvgamerr-app/go-bitkub/bitkub"
)

type PlaceBidRequest struct {
	Sym      string  `json:"sym"`
	Amt      float64 `json:"amt"`
	Rat      float64 `json:"rat"`
	Typ      string  `json:"typ"`
	ClientID string  `json:"client_id,omitempty"`
	PostOnly bool    `json:"post_only,omitempty"`
}

type PlaceBidResult struct {
	ID       string  `json:"id"`
	Typ      string  `json:"typ"`
	Amt      float64 `json:"amt"`
	Rat      float64 `json:"rat"`
	Fee      float64 `json:"fee"`
	Cre      float64 `json:"cre"`
	Rec      float64 `json:"rec"`
	Ts       string  `json:"ts"`
	ClientID string  `json:"ci"`
}

type PlaceBidResponse struct {
	Error  int            `json:"error"`
	Result PlaceBidResult `json:"result"`
}

func PlaceBid(req PlaceBidRequest) (*PlaceBidResult, error) {
	var response bitkub.ResponseAPI

	if err := bitkub.FetchSecure("POST", "/v3/market/place-bid", req, &response); err != nil {
		return nil, err
	}

	if err := response.CheckResponseError(); err != nil {
		return nil, err
	}

	byteData, err := stdJson.Marshal(response.Result)
	if err != nil {
		return nil, err
	}

	var result PlaceBidResult
	if err = stdJson.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
