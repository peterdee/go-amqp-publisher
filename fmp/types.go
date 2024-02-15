package fmp

type ListQuoteData []struct {
	Symbol            string  `json:"symbol"`
	Name              string  `json:"name"`
	Price             float64 `json:"price"`
	Exchange          string  `json:"exchange"`
	ExchangeShortName string  `json:"exchangeShortName"`
	Type              string  `json:"type"`
}

type QuoteData struct {
	Symbol               string  `json:"symbol"`
	Name                 string  `json:"name"`
	Price                float64 `json:"price"`
	ChangesPercentage    float64 `json:"changesPercentage"`
	Change               float64 `json:"change"`
	DayLow               float64 `json:"dayLow"`
	DayHigh              float64 `json:"dayHigh"`
	YearHigh             float64 `json:"yearHigh"`
	YearLow              float64 `json:"yearLow"`
	MarketCap            int     `json:"marketCap"`
	PriceAvg50           float64 `json:"priceAvg50"`
	PriceAvg200          float64 `json:"priceAvg200"`
	Exchange             string  `json:"exchange"`
	Volume               int     `json:"volume"`
	AvgVolume            int     `json:"avgVolume"`
	Open                 float64 `json:"open"`
	PreviousClose        float64 `json:"previousClose"`
	Eps                  float64 `json:"eps"`
	Pe                   float64 `json:"pe"`
	EarningsAnnouncement string  `json:"earningsAnnouncement"`
	SharesOutstanding    int     `json:"sharesOutstanding"`
	Timestamp            int     `json:"timestamp"`
}

type QuoteDataResponse []QuoteData
