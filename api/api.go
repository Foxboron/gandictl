package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const URL = "https://api.gandi.net/v5"

type Client struct {
	ContentTextPlain bool
	AcceptTextPlain  bool
	AcceptJSON       bool
	ApiKey           string
	http             *http.Client
	Method           string
	Buffer           []byte
}

func NewClient(ApiKey string) *Client {
	return &Client{Method: http.MethodGet, ApiKey: ApiKey, http: &http.Client{}, Buffer: []byte{}}
}

func (c *Client) ResetOptions() {
	c.Method = http.MethodGet
	c.http = &http.Client{}
	c.Buffer = []byte{}
	c.AcceptTextPlain = false
	c.ContentTextPlain = false
}

func (c *Client) DoRequest(uri string) ([]byte, *ErrorReponse) {
	defer c.ResetOptions() // So we don't do silly mistakes
	req, err := http.NewRequest(c.Method, URL+uri, bytes.NewBuffer(c.Buffer))
	if err != nil {
		log.Fatal(err)
	}
	if c.AcceptTextPlain {
		req.Header.Set("Accept", "text/plain")
	}
	if c.ContentTextPlain {
		req.Header.Set("Content-Type", http.DetectContentType(c.Buffer))
	}
	req.Header.Set("Authorization", fmt.Sprintf("Apikey %s", c.ApiKey))
	res, err := c.http.Do(req)
	if err != nil {
		return nil, &ErrorReponse{}
	}
	if res.StatusCode != 200 && res.StatusCode != 201 {
		body, _ := ioutil.ReadAll(res.Body)
		return []byte{}, NewErrorResponse(body)
	}
	body, _ := ioutil.ReadAll(res.Body)
	return body, nil
}

func GetRecords(c *Client, d string) []*Record {
	resp, err := c.DoRequest(fmt.Sprintf("/livedns/domains/%s/records", d))
	if err != nil {
		log.Fatal(err)
	}
	var records []*Record
	json.Unmarshal(resp, &records)
	return records
}

func GetDomains(c *Client) []*Domain {
	resp, err := c.DoRequest("/livedns/domains")
	if err != nil {
		log.Fatal(err)
	}
	var domains []*Domain
	json.Unmarshal(resp, &domains)
	return domains
}

func GetZonefile(c *Client, d string) []byte {
	c.AcceptTextPlain = true
	resp, err := c.DoRequest(fmt.Sprintf("/livedns/domains/%s/records", d))
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func WriteZonefile(c *Client, domain string, buf []byte) []byte {
	c.AcceptTextPlain = false
	c.ContentTextPlain = true
	c.Method = http.MethodPut
	c.Buffer = buf
	resp, err := c.DoRequest(fmt.Sprintf("/livedns/domains/%s/records", domain))
	if err != nil {
		log.Fatal(err)
	}
	return resp
}
