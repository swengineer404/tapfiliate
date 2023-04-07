package tapfiliate

type ConversionCreateParams struct {
	ReferralCode string  `json:"referral_code"`
	ExternalID   string  `json:"external_id"`
	Amount       float64 `json:"amount"`
}

type ConversionCreateResult struct {
	ID int `json:"id"`
}

type Conversion struct {
	client *Client
}

func NewConversionService(client *Client) *Conversion {
	return &Conversion{
		client: client,
	}
}

func (c *Conversion) Create(params *ConversionCreateParams) (*ConversionCreateResult, error) {
	var result ConversionCreateResult
	return &result, c.client.Do("POST", "/conversions/", params, &result)
}
