package market

import "github.com/dvgamerr-app/go-bitkub/utils"

type (
	PlaceBidRequest  PlaceAskRequest
	PlaceBidResult   PlaceAskResult
	PlaceBidResponse PlaceAskResponse
)

func PlaceBid(req PlaceBidRequest) (*PlaceBidResult, error) {
	req.Symbol = utils.UppercaseSymbol(req.Symbol)

	return placeOrder[PlaceBidResult]("/api/v3/market/place-bid", req)
}
