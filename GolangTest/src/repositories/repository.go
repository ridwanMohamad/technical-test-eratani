package repositories

import (
	"context"
	"golangtest/src/models"
	"golangtest/src/payload/request"

	"gorm.io/gorm"
)

type repositories struct {
	db *gorm.DB
}

func InitialRepositories(db *gorm.DB) Repositories {
	return &repositories{
		db: db,
	}
}

type Repositories interface {
	GetDataSpendByCountry(ctx context.Context) (dataSpendCountry []*models.TotalSpendCountry, err error)
	GetDataTotalCardType(ctx context.Context) (dataTotalCardType []*models.TotalCardType, err error)
	GetDataByCountry(ctx context.Context, country string) (data []*models.Users, err error)
	InsertDataUser(ctx context.Context, req request.ReqUsers) (data *models.Users, err error)
}
