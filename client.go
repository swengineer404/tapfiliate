package tapfiliate

import "fmt"

type Client struct {
	Click *Click
	rest  *restClient
}

func New(apiKey string) *Client {
	rc := newRestClient(apiKey)

	c := &Client{
		rest: rc,
	}

	c.Click = NewClickService(c)

	return c
}

func (c *Client) Do(method, path string, dto, result any) error {
	if err := c.rest.do(method, path, dto, result); err != nil {
		return fmt.Errorf("tapfiliate_api_error: %w", err)
	}

	return nil
}
