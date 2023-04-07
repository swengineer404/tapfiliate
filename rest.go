package tapfiliate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type restClient struct {
	apiKey  string
	baseURL string
	c       *http.Client
}

func newRestClient(apiKey string) *restClient {
	return &restClient{
		apiKey:  apiKey,
		baseURL: "https://api.tapfiliate.com/1.6",
		c: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (c *restClient) do(method, path string, dto, result any) error {
	var b bytes.Buffer
	if dto != nil {
		if err := json.NewEncoder(&b).Encode(dto); err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, c.baseURL+path, &b)
	if err != nil {
		return err
	}

	req.Header.Set("X-Api-Key", c.apiKey)
	req.Header.Set("User-Agent", "go-tapfiliate-client")

	if dto != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := c.c.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	resb, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return fmt.Errorf("status_code: %d | message: %s", res.StatusCode, resb)
	}

	if err := json.Unmarshal(resb, result); err != nil {
		return fmt.Errorf("failed to decode[%s]: %s", err, resb)
	}

	return nil
}
