package fmp

type FMPService struct {
	apiKey   string
	endpoint string
}

var FMP FMPService

func (instance FMPService) New(apiKey, endpoint string) FMPService {
	instance.apiKey = apiKey
	instance.endpoint = endpoint
	FMP = instance
	return FMP
}
