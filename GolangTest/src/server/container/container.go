package container

import (
	"golangtest/src/repositories"
	"golangtest/src/server/config"
	"golangtest/src/server/database"
	"golangtest/src/usecase"
)

type DefaultContainer struct {
	Config  *config.DefaultConfig
	Usecase usecase.UseCase
}

func IntializeContainer() *DefaultContainer {

	config := config.ConfigApps("./resources/")
	db := database.InitializeDatabase(config)
	repo := repositories.InitialRepositories(db.DB)
	uc := usecase.InitializeUseCase(repo)

	_ = db.Raw("SELECT 1").Row()

	return &DefaultContainer{
		Config:  config,
		Usecase: uc,
	}

}
