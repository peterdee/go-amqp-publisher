package fmp

type Service struct {
	API_KEY  string
	ENDPOINT string
}

var ServiceInstance Service

func (s Service) New(apiKey, endpoint string) Service {
	s.API_KEY = apiKey
	s.ENDPOINT = endpoint
	ServiceInstance = s
	return ServiceInstance
}
