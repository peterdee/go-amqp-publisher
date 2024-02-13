package utilities

import (
	"github.com/monaco-io/request"
)

func HttpGet(endpoint string) (interface{}, error) {
	client := request.Client{
		Method: "GET",
		URL:    endpoint,
	}
	var result interface{}
	response := client.Send().Scan(&result)
	if !(response.OK() && 100 < response.Code() && response.Code() < 300) {
		return nil, response.Error()
	}
	return result, nil
}
