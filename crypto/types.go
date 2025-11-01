package crypto

type PaginatedResponse struct {
	Page      int `json:"page"`
	TotalPage int `json:"total_page"`
	TotalItem int `json:"total_item"`
}

type Address struct {
	Symbol    string `json:"symbol"`
	Network   string `json:"network"`
	Address   string `json:"address"`
	Memo      string `json:"memo,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type AddressesResponse struct {
	PaginatedResponse
	Items []Address `json:"items"`
}

// CreateAddressRequest is the request body for POST /api/v4/crypto/addresses
type CreateAddressRequest struct {
	Symbol  string `json:"symbol"`
	Network string `json:"network"`
}

// Deposit represents a crypto deposit transaction
type Deposit struct {
	Hash          string `json:"hash"`
	Symbol        string `json:"symbol"`
	Network       string `json:"network"`
	Amount        string `json:"amount"`
	FromAddress   string `json:"from_address"`
	ToAddress     string `json:"to_address"`
	Confirmations int    `json:"confirmations"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
	CompletedAt   string `json:"completed_at,omitempty"`
}

// DepositsResponse is the response for GET /api/v4/crypto/deposits
type DepositsResponse struct {
	PaginatedResponse
	Items []Deposit `json:"items"`
}

// Withdraw represents a crypto withdrawal transaction
type Withdraw struct {
	TxnID       string `json:"txn_id"`
	ExternalRef string `json:"external_ref,omitempty"`
	Hash        string `json:"hash,omitempty"`
	Symbol      string `json:"symbol"`
	Network     string `json:"network"`
	Amount      string `json:"amount"`
	Fee         string `json:"fee"`
	Address     string `json:"address"`
	Memo        string `json:"memo,omitempty"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	CompletedAt string `json:"completed_at,omitempty"`
}

// WithdrawsResponse is the response for GET /api/v4/crypto/withdraws
type WithdrawsResponse struct {
	PaginatedResponse
	Items []Withdraw `json:"items"`
}

// CreateWithdrawRequest is the request body for POST /api/v4/crypto/withdraws
type CreateWithdrawRequest struct {
	Symbol  string `json:"symbol"`
	Amount  string `json:"amount"`
	Address string `json:"address"`
	Memo    string `json:"memo,omitempty"`
	Network string `json:"network"`
}

// CreateWithdrawResponse is the response for POST /api/v4/crypto/withdraws
type CreateWithdrawResponse struct {
	TxnID     string `json:"txn_id"`
	Symbol    string `json:"symbol"`
	Network   string `json:"network"`
	Amount    string `json:"amount"`
	Fee       string `json:"fee"`
	Address   string `json:"address"`
	Memo      string `json:"memo,omitempty"`
	CreatedAt string `json:"created_at"`
}

// Network represents a cryptocurrency network
type Network struct {
	Name                  string `json:"name"`
	Network               string `json:"network"`
	AddressRegex          string `json:"address_regex"`
	MemoRegex             string `json:"memo_regex"`
	Explorer              string `json:"explorer"`
	ContractAddress       string `json:"contract_address"`
	WithdrawMin           string `json:"withdraw_min"`
	WithdrawFee           string `json:"withdraw_fee"`
	WithdrawInternalMin   string `json:"withdraw_internal_min"`
	WithdrawInternalFee   string `json:"withdraw_internal_fee"`
	WithdrawDecimalPlaces int    `json:"withdraw_decimal_places"`
	MinConfirm            int    `json:"min_confirm"`
	Decimal               int    `json:"decimal"`
	DepositEnable         bool   `json:"deposit_enable"`
	WithdrawEnable        bool   `json:"withdraw_enable"`
	IsMemo                bool   `json:"is_memo"`
}

// Coin represents a cryptocurrency with its networks
type Coin struct {
	Name           string    `json:"name"`
	Symbol         string    `json:"symbol"`
	Networks       []Network `json:"networks"`
	DepositEnable  bool      `json:"deposit_enable"`
	WithdrawEnable bool      `json:"withdraw_enable"`
}

// CoinsResponse is the response for GET /api/v4/crypto/coins
type CoinsResponse struct {
	Items []Coin `json:"items"`
}

// Compensation represents a crypto compensation transaction
type Compensation struct {
	TxnID       string `json:"txn_id"`
	Symbol      string `json:"symbol"`
	Type        string `json:"type"` // COMPENSATE or DECOMPENSATE
	Amount      string `json:"amount"`
	Status      string `json:"status"` // PENDING or COMPLETED
	CreatedAt   string `json:"created_at"`
	CompletedAt string `json:"completed_at,omitempty"`
}

// CompensationsResponse is the response for GET /api/v4/crypto/compensations
type CompensationsResponse struct {
	PaginatedResponse
	Items []Compensation `json:"items"`
}
