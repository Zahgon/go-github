package scrape

func (c *Client) OrgPaymentInformation(org string) (PaymentInformation, error) {
	_ = "STUB: not implemented"
	return *new(PaymentInformation), nil
}

type PaymentInformation struct {
	PaymentMethod    string
	LastPayment      string
	Coupon           string
	ExtraInformation string
}
