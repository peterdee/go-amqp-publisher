package fmp

import (
	"fmt"

	"go-amqp-publisher/utilities"

	"github.com/opensaucerer/goaxios"
)

type FMPService struct {
	apiKey   string
	endpoint string
}

var FMP FMPService

// Get quote data for a stock
func (instance FMPService) GetQuote(symbol string) (interface{}, error) {
	request := goaxios.GoAxios{
		Method:         "GET",
		ResponseStruct: &QuoteData{},
		Url:            fmt.Sprintf("%s/quote/%s?apikey=%s", instance.endpoint, symbol, instance.apiKey),
	}
	response := request.RunRest()

	quoteData := *response.Body.(*QuoteData)
	fmt.Println(
		"here",
		quoteData[0].Change,
	)
	return utilities.HttpGet(
		fmt.Sprintf("%s/quote/%s?apikey=%s", instance.endpoint, symbol, instance.apiKey),
	)
}

// Get a list of all stocks
func (instance FMPService) GetStocks() (interface{}, error) {
	return utilities.HttpGet(fmt.Sprintf("%s/stocks/list?apikey=%s", instance.endpoint, instance.apiKey))
}

// Create a new instace of FMP service
func (instance FMPService) New(apiKey, endpoint string) FMPService {
	instance.apiKey = apiKey
	instance.endpoint = endpoint
	FMP = instance
	return FMP
}
