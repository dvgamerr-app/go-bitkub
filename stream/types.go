package stream

import "time"

type StreamType string

const (
	StreamTypeTrade     StreamType = "market.trade"
	StreamTypeTicker    StreamType = "market.ticker"
	StreamTypeOrderBook StreamType = "orderbook"
)

type TradeData struct {
	Stream string  `json:"stream"`
	Symbol string  `json:"sym"`
	TxnID  string  `json:"txn"`
	Rate   string  `json:"rat"`
	Amount float64 `json:"amt"`
	BuyID  string  `json:"bid"`
	SellID string  `json:"sid"`
	TS     int64   `json:"ts"`
}

type TickerData struct {
	Stream         string  `json:"stream"`
	ID             int     `json:"id"`
	Last           float64 `json:"last"`
	LowestAsk      float64 `json:"lowestAsk"`
	LowestAskSize  float64 `json:"lowestAskSize"`
	HighestBid     float64 `json:"highestBid"`
	HighestBidSize float64 `json:"highestBidSize"`
	Change         float64 `json:"change"`
	PercentChange  float64 `json:"percentChange"`
	BaseVolume     float64 `json:"baseVolume"`
	QuoteVolume    float64 `json:"quoteVolume"`
	IsFrozen       int     `json:"isFrozen"`
	High24hr       float64 `json:"high24hr"`
	Low24hr        float64 `json:"low24hr"`
	Open           float64 `json:"open"`
	Close          float64 `json:"close"`
}

type OrderBookOrder struct {
	Volume   float64 `json:"volume"`
	Rate     float64 `json:"rate"`
	Amount   float64 `json:"amount"`
	Reserved float64 `json:"reserved"`
	IsNew    bool    `json:"isNew"`
	IsOwner  bool    `json:"isOwner"`
}

type TradeOrder struct {
	Timestamp int64   `json:"timestamp"`
	Rate      float64 `json:"rate"`
	Amount    float64 `json:"amount"`
	Side      string  `json:"side"`
	Reserved1 int     `json:"reserved1"`
	Reserved2 int     `json:"reserved2"`
	IsNew     bool    `json:"isNew"`
	IsBuyer   bool    `json:"isBuyer"`
	IsSeller  bool    `json:"isSeller"`
}

type OrderBookData struct {
	Data      any    `json:"data"`
	Event     string `json:"event"`
	PairingID int    `json:"pairing_id,omitempty"`
}

type StreamConfig struct {
	ReconnectInterval time.Duration
	MaxReconnect      int
	PingInterval      time.Duration
	ReadTimeout       time.Duration
	MessageBuffer     int
}

type Message struct {
	Data      any
	Type      string
	Error     error
	Timestamp time.Time
}
