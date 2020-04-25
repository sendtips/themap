package themap

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (p *Payment) Init(amount int) error {
	p.Amount = amount
	return p.makeInit()
}

func (p *Payment) makeInit() error {
	var err error
	var result bytes.Buffer
	var req *http.Request
	var resp *http.Response

	req, err = makeRequest("/Init", "POST", nil, nil)

	resp, err = proceedRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("[THEMAP]", "Resp code:", resp.StatusCode)
		return ErrBadStatusReply
	}

	_, err = io.Copy(&result, resp.Body)

	err = json.Unmarshal(result.Bytes(), &p.Reply)
	if err != nil {
		log.Println(result.String(), err)
		log.Println("[THEMAP]", "Cant unmarshal JSON payload")
		return ErrBadJSON
	}

	return nil
}
