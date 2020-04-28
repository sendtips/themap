package themap

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

// Requester makes request to theMAP
type Sender interface {
	Send() error
}

var (
	// ErrBadJSON error throws when JSON marshal/unmarshal problem occurs
	ErrBadJSON = errors.New("Bad reply payload")
	// ErrBadStatusReply is bad gateway stratus code
	ErrBadStatusReply = errors.New("Bad status reply")
	// ErrReplyWithError business-logic error
	ErrReplyWithError = errors.New("Error in reply")
)

const apilink string = "https://api-stage.mapisacard.com" // no trailing slash

// newRequest creates new http request
// link is a url
// method HTTP method
// headers can contain additional headers
// payload is a body of HTTP request
func newRequest(method, link string, headers map[string]string, payload []byte) (*http.Request, error) {

	req, err := http.NewRequest(method, apilink+link, bytes.NewBuffer(payload))
	req.Header.Set("User-Agent", "Sendtips/"+clientversion)
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
func proceedRequest(method, link string, p Sender) error {
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

	req, err = newRequest(method, link, nil, payload)
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

	err = json.Unmarshal(result.Bytes(), &p)
	if err != nil {
		log.Println(result.String(), err)
		log.Println("[THEMAP]", "Cant unmarshal JSON payload")
		return ErrBadJSON
	}

	return err
}
