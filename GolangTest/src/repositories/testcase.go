package repositories

import (
	"context"
	"golangtest/src/models"
	"golangtest/src/payload/request"
	"strings"
)

func (rep *repositories) GetDataSpendByCountry(ctx context.Context) (dataSpendCountry []*models.TotalSpendCountry, err error) {
	db := rep.db.Model(&models.Users{}).WithContext(ctx)
	db.Select("country", "COALESCE(SUM(belanjas.total_buy),0) as total_spend").
		Joins("LEFT JOIN public.belanjas on users.id = belanjas.id_user").
		Group("country").
		Order("total_spend desc")
	var data []*models.TotalSpendCountry
	err = db.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
func (rep *repositories) GetDataTotalCardType(ctx context.Context) (dataTotalCardType []*models.TotalCardType, err error) {
	db := rep.db.Model(&models.Users{}).WithContext(ctx)
	db.Select("credit_card_type", "COUNT(credit_card_type) as total").
		Group("credit_card_type").
		Order("total desc")
	var data []*models.TotalCardType
	err = db.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
func (rep *repositories) GetDataByCountry(ctx context.Context, country string) (data []*models.Users, err error) {
	db := rep.db.Model(&models.Users{}).WithContext(ctx).Where("UPPER(country) = ?", strings.ToUpper(country))
	err = db.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
func (rep *repositories) InsertDataUser(ctx context.Context, req request.ReqUsers) (data *models.Users, err error) {
	db := rep.db.Model(&models.Users{}).WithContext(ctx)
	var lastdata *models.Users
	err = db.Last(&lastdata).Error
	if err != nil {
		return nil, err
	}

	currentId := lastdata.ID + 1
	insertData := models.Users{
		ID:             currentId,
		Country:        req.Country,
		CreditCardType: req.CreditCardType,
		CreditCard:     req.CreditCard,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
	}

	err = db.Create(&insertData).Error
	if err != nil {
		return nil, err
	}

	return &insertData, nil
}
