package market

import (
	"github.com/dvgamerr-app/go-bitkub/bitkub"
	"github.com/dvgamerr-app/go-bitkub/utils"
)

type PlaceAskRequest struct {
	Symbol   string  `json:"sym"`
	Amount   float64 `json:"amt"`
	Rate     float64 `json:"rat"`
	Type     string  `json:"typ"`
	ClientID string  `json:"client_id,omitempty"`
	PostOnly bool    `json:"post_only,omitempty"`
}

type PlaceAskResult struct {
	ID        string  `json:"id"`
	Type      string  `json:"typ"`
	Amount    float64 `json:"amt"`
	Rate      float64 `json:"rat"`
	Fee       float64 `json:"fee"`
	Credit    float64 `json:"cre"`
	Receive   float64 `json:"rec"`
	Timestamp string  `json:"ts"`
	ClientID  string  `json:"ci"`
}

type PlaceAskResponse struct {
	bitkub.Error
	Result PlaceAskResult `json:"result"`
}

func PlaceAsk(req PlaceAskRequest) (*PlaceAskResult, error) {
	var response bitkub.ResponseAPI

	req.Symbol = utils.UppercaseSymbol(req.Symbol)

	if err := bitkub.FetchSecure("POST", "/api/v3/market/place-ask", req, &response); err != nil {
		return nil, err
	}

	if err := response.CheckResponseError(); err != nil {
		return nil, err
	}

	byteData, err := stdJson.Marshal(response.Result)
	if err != nil {
		return nil, err
	}

	var result PlaceAskResult
	if err = stdJson.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
