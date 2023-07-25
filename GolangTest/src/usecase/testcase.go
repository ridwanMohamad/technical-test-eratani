package usecase

import (
	"context"
	"golangtest/src/models"
	"golangtest/src/payload/request"
)

func (u *usecase) GetCountrySpend(ctx context.Context) (countrySpend []*models.TotalSpendCountry, err error) {
	getDataSpendCountry, err := u.repositories.GetDataSpendByCountry(ctx)
	if err != nil {
		return nil, err
	}
	return getDataSpendCountry, err
}
func (u *usecase) GetCardTypeTotal(ctx context.Context) (cardTypeTotal []*models.TotalCardType, err error) {
	getDataTotalCardType, err := u.repositories.GetDataTotalCardType(ctx)
	if err != nil {
		return nil, err
	}
	return getDataTotalCardType, err
}
func (u *usecase) GetDataByCountry(ctx context.Context, country string) (data []*models.Users, err error) {
	getData, err := u.repositories.GetDataByCountry(ctx, country)
	if err != nil {
		return nil, err
	}
	return getData, err
}
func (u *usecase) CreateDataUser(ctx context.Context, req request.ReqUsers) (data *models.Users, err error) {
	getData, err := u.repositories.InsertDataUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return getData, err
}
