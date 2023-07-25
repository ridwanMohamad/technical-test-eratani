package models

type Users struct {
	ID             int    `gorm:"type:int;column:id" json:"id"`
	Country        string `gorm:"type:varchar(255);column:country" json:"country"`
	CreditCardType string `gorm:"type:varchar(255);column:credit_card_type" json:"credit_card_type"`
	CreditCard     string `gorm:"type:varchar(255);column:credit_card" json:"credit_card"`
	FirstName      string `gorm:"type:varchar(255);column:first_name" json:"first_name"`
	LastName       string `gorm:"type:varchar(255);column:last_name" json:"last_name"`
}

type TotalCardType struct {
	CardType string `gorm:"column:credit_card_type" json:"credit_card_type"`
	Total    int    `gorm:"column:total" json:"Total"`
}
