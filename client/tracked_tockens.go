package client

type TrackedTokenResponse struct {
	Data   []TrackedToken `json:"data"`
	Status string         `json:"status"`
}

type TrackedToken struct {
	CapUsd       float64 `json:"cap_usd"`
	CoinId       int     `json:"coinid"`
	Holders      int     `json:"holders"`
	LiquidityUsd float64 `json:"liquidity_usd"`
	PriceBip     float64 `json:"price_bip"`
	PriceUsd     float64 `json:"price_usd"`
	Ticker       string  `json:"ticker"`
	Tiks24H      int     `json:"tiks24h"`
	Volume24HUsd float64 `json:"volume24h_usd"`
}
