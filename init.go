package themap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Init obtain session token from TheMAP payment gateway
func (p *Payment) Init(amount int) error {
	p.Amount = amount
	return p.makeInit()
}

func (p *Payment) makeInit() error {
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

	req, err = makeRequest("/Init", "POST", nil, payload)

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

	if p.Reply.ErrCode != "" {
		err = fmt.Errorf("[THEMAP] %w: %s\n", ErrReplyWithError, p.Reply.ErrCode)
	}

	return err
}
