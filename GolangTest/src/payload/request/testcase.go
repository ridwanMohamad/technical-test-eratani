package request

type ReqUsers struct {
	Country        string `json:"country" validate:"required"`
	CreditCardType string `json:"credit_card_type" validate:"required"`
	CreditCard     string `json:"credit_card" validate:"required"`
	FirstName      string `json:"first_name" validate:"required"`
	LastName       string `json:"last_name" validate:"required"`
}
