package gokirimwa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c Client) do(method string, url string, queries map[string]string, body any, dest any) error {
	var (
		b       bytes.Buffer
		reqBody io.Reader = nil
	)

	if err := json.NewEncoder(&b).Encode(body); err != nil && body != nil {
		return err
	}

	if body != nil {
		reqBody = &b
	}

	url = fmt.Sprintf("%s/%s", c.baseURl, url)
	log.Printf("sending request [%s] %s", method, url)
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.apiKey)
	req.URL.RawQuery = c.appendQueries(req, queries)

	response, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if response.StatusCode >= http.StatusBadRequest {
		var errResp ErrorResponse
		json.NewDecoder(response.Body).Decode(&errResp)
		return errResp
	}

	if err = json.NewDecoder(response.Body).Decode(&dest); err != nil && dest != nil {
		return err
	}

	return nil
}

func (c Client) appendQueries(req *http.Request, queries map[string]string) string {
	q := req.URL.Query()

	for key, value := range queries {
		q.Add(key, value)
	}

	return q.Encode()
}
