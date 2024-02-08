package fmp

import (
	"fmt"

	"go-amqp-publisher/utilities"
)

type FMPService struct {
	apiKey   string
	endpoint string
}

var FMP FMPService

// Get a list of all stocks
func (instance FMPService) GetStocks() ([]byte, error) {
	endpoint := fmt.Sprintf("%s/stock/list?apikey=%s", instance.endpoint, instance.apiKey)
	fmt.Println(endpoint)
	return utilities.HttpGet(fmt.Sprintf("%s/stocks/list?apikey=%s", instance.endpoint, instance.apiKey))
}

// Create a new instace of FMP service
func (instance FMPService) New(apiKey, endpoint string) FMPService {
	instance.apiKey = apiKey
	instance.endpoint = endpoint
	FMP = instance
	return FMP
}
