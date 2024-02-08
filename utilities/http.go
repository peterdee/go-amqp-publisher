package utilities

import "github.com/monaco-io/request"

func HttpGet(endpoint string) ([]byte, error) {
	client := request.Client{
		Method: "GET",
		URL:    endpoint,
	}
	var result interface{}
	response := client.Send().Scan(&result)
	if !response.OK() {
		return nil, response.Error()
	}
	return response.Bytes(), nil
}
