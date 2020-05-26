package themap

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

var (
	// ErrBadJSON error throws when JSON marshal/unmarshal problem occurs
	ErrBadJSON = errors.New("Bad reply payload")

	// ErrBadStatusReply is bad gateway status code
	ErrBadStatusReply = errors.New("Bad status reply")

	// ErrReplyWithError business-logic error
	ErrReplyWithError = errors.New("Error in reply")

	// ErrBadSignature wrong signature error
	ErrBadSignature = errors.New("Wrong signature")
)

// APILink sets payment gateway domain
var APILink string = "https://api-stage.mapcard.pro" // no trailing slash

// newRequest creates new http request.
// The params is path is a url part
// HTTP method, then map[string]string with additional headers
// and a body of request
func newRequest(method, path string, headers map[string]string, payload []byte) (*http.Request, error) {

	req, err := http.NewRequest(method, APILink+path, bytes.NewBuffer(payload))
	req.Header.Set("User-Agent", userAgent+"/"+Version)
	req.Header.Set("Content-Type", "application/json")

	// Additional HTTP headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Set(key, val)
		}
	}

	return req, err
}

// doRequest executing http.Request and return body
func doRequest(req *http.Request) (*http.Response, error) {
	client := http.Client{}
	return client.Do(req)
}

// proceedRequest deal with data prep and preceedRequest
// handle response and pack all data back to our structure
func proceedRequest(method, path string, p *Payment) error {
	var err error
	var payload []byte
	var result bytes.Buffer
	var req *http.Request
	var resp *http.Response

	payload, err = json.Marshal(p)
	if err != nil {
		log.Println("[THEMAP]", "Cant marshal payload")
		return ErrBadJSON
	}

	req, err = newRequest(method, path, nil, payload)
	if err != nil {
		return err
	}

	resp, err = doRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("[THEMAP]", "Resp code:", resp.StatusCode)
		return ErrBadStatusReply
	}

	_, err = io.Copy(&result, resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(result.Bytes(), &p)
	if err != nil {
		log.Println(result.String(), err)
		log.Println("[THEMAP]", "Cant unmarshal JSON payload")
		return ErrBadJSON
	}

	return err
}
