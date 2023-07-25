package usecase

import (
	"context"
	"golangtest/src/models"
	"golangtest/src/payload/request"
	"golangtest/src/repositories"
)

type usecase struct {
	repositories repositories.Repositories
}

type UseCase interface {
	GetCountrySpend(ctx context.Context) (countrySpend []*models.TotalSpendCountry, err error)
	GetCardTypeTotal(ctx context.Context) (cardTypeTotal []*models.TotalCardType, err error)
	GetDataByCountry(ctx context.Context, country string) (data []*models.Users, err error)
	CreateDataUser(ctx context.Context, req request.ReqUsers) (data *models.Users, err error)
}

func InitializeUseCase(repo repositories.Repositories) UseCase {
	return &usecase{
		repositories: repo,
	}
}
