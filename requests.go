package themap

import (
	"bytes"
	"errors"
	"net/http"
)

var (
	ErrBadJSON        = errors.New("Bad reply payload")
	ErrBadStatusReply = errors.New("Bad status reply")
)

const _APIlink string = "https://api-stage.mapisacard.com" // no trailing slash

func makeRequest(link string, method string, headers map[string]string, payload []byte) (*http.Request, error) {

	req, err := http.NewRequest(method, _APIlink+link, bytes.NewBuffer(payload))
	req.Header.Set("User-Agent", "Sendtips/"+version)
	req.Header.Set("Content-Type", "application/json")
	//req.SetBasicAuth(clientID, secret)

	// Additional HTTP headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Set(key, val)
		}
	}

	return req, err
}

// proceedRequest executing http.Request and return body
func proceedRequest(req *http.Request) (*http.Response, error) {

	client := http.Client{}

	return client.Do(req)
}
