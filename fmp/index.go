package fmp

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/opensaucerer/goaxios"
)

type FMPService struct {
	apiKey   string
	endpoint string
}

var FMP FMPService

// Get quote data for a stock
func (instance FMPService) GetQuote(symbol string) (QuoteData, error) {
	request := goaxios.GoAxios{
		Method:         "GET",
		ResponseStruct: &QuoteDataResponse{},
		Url:            fmt.Sprintf("%s/quote/%s?apikey=%s", instance.endpoint, symbol, instance.apiKey),
	}
	response := request.RunRest()

	quoteData := *response.Body.(*QuoteDataResponse)
	if len(quoteData) == 0 {
		return QuoteData{}, errors.New("could not load quote data")
	}
	return quoteData[0], nil
}

// Get a list of all stocks
func (instance FMPService) GetStocks() (ListQuoteData, error) {
	request := goaxios.GoAxios{
		Method:         "GET",
		ResponseStruct: &ListQuoteData{},
		Url:            fmt.Sprintf("%s/stock/list?apikey=%s", instance.endpoint, instance.apiKey),
	}
	response := request.RunRest()

	listQuoteData := *response.Body.(*ListQuoteData)
	if len(listQuoteData) == 0 {
		return ListQuoteData{}, errors.New("could not load quote list")
	}
	return listQuoteData, nil
}

// Create a new instace of FMP service
func (instance FMPService) New() FMPService {
	apiKey := os.Getenv("FMP_API_KEY")
	endpoint := os.Getenv("FMP_ENDPOINT")
	if apiKey == "" || endpoint == "" {
		log.Fatal("Could not load FMP configuration")
	}

	instance.apiKey = apiKey
	instance.endpoint = endpoint
	FMP = instance
	return FMP
}
