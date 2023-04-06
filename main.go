package main

import (
	"io/ioutil"
	"net/http"
)

func callAPI() ([]byte, error) {
	url := "https://desbravador.movidesk.com/public/v1"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// func decodeResponse(body []byte) (*Response, error) {
//     var resp Response
//     err := json.Unmarshal(body, &resp)
//     if err != nil {
//         return nil, err
//     }
//     return &resp, nil
// }
