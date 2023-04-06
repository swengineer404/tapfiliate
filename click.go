package tapfiliate

type ClickCreateParams struct {
	ReferralCode string `json:"referral_code"`
}

type ClickCreateResult struct {
	ID string `json:"id"`
}

type Click struct {
	client *Client
}

func NewClickService(client *Client) *Click {
	return &Click{
		client: client,
	}
}

func (c *Click) Create(params *ClickCreateParams) (*ClickCreateResult, error) {
	var result ClickCreateResult
	return &result, c.client.Do("POST", "/clicks/", params, &result)
}
