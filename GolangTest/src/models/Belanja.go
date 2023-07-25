package models

type Belanja struct {
	ID       int `gorm:"type:int;column:id" json:"id"`
	IdUser   int `gorm:"type:int;column:id_user" json:"id_user"`
	TotalBuy int `gorm:"tyoe:int;column:total_buy" json:"total_buy"`
}

type TotalSpendCountry struct {
	Country    string `gorm:"column:country" json:"country"`
	TotalSpend int    `gorm:"column:total_spend" json:"total_spend"`
}
