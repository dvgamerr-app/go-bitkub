package crypto

type Pagination struct {
	Page  int
	Limit int
}

type DateRange struct {
	CreatedStart string
	CreatedEnd   string
}

type SymbolNetwork struct {
	Symbol  string
	Network string
}

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

type CreateAddressRequest struct {
	Symbol  string `json:"symbol"`
	Network string `json:"network"`
}

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

type DepositsResponse struct {
	PaginatedResponse
	Items []Deposit `json:"items"`
}

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

type WithdrawsResponse struct {
	PaginatedResponse
	Items []Withdraw `json:"items"`
}

type CreateWithdrawRequest struct {
	Symbol  string `json:"symbol"`
	Amount  string `json:"amount"`
	Address string `json:"address"`
	Memo    string `json:"memo,omitempty"`
	Network string `json:"network"`
}

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

type Coin struct {
	Name           string    `json:"name"`
	Symbol         string    `json:"symbol"`
	Networks       []Network `json:"networks"`
	DepositEnable  bool      `json:"deposit_enable"`
	WithdrawEnable bool      `json:"withdraw_enable"`
}

type CoinsResponse struct {
	Items []Coin `json:"items"`
}

type Compensation struct {
	TxnID       string `json:"txn_id"`
	Symbol      string `json:"symbol"`
	Type        string `json:"type"`
	Amount      string `json:"amount"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	CompletedAt string `json:"completed_at,omitempty"`
}

type CompensationsResponse struct {
	PaginatedResponse
	Items []Compensation `json:"items"`
}
