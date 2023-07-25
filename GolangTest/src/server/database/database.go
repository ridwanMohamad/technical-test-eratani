package database

import (
	"fmt"
	"golangtest/src/models"
	"golangtest/src/server/config"
	"golangtest/src/server/database/seeds"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Database struct {
	*gorm.DB
}

func InitializeDatabase(ds *config.DefaultConfig) *Database {

	//mysql conn
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	ds.Username,
	//	ds.Password,
	//	ds.Url,
	//	ds.Port,
	//	ds.DatabaseName)

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=%s",
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable search_path=%s",
		ds.Database.Url,
		ds.Database.Username,
		ds.Database.DatabaseName,
		ds.Database.Port,
		ds.Database.Schema)

	fmt.Println(dsn)
	cfg := &gorm.Config{}
	if ds.Database.DebugMode {
		cfg = &gorm.Config{Logger: logger.Default.LogMode(logger.Info)}
	}

	cfg.NamingStrategy = schema.NamingStrategy{TablePrefix: ds.Database.Schema, SingularTable: false}

	db, err := gorm.Open(postgres.Open(dsn), cfg)
	if err != nil {
		fmt.Println("Failed to open connection")
		panic(err)
	}

	// Auto Migration Models
	db.AutoMigrate(
		models.Users{},
		models.Belanja{},
	)

	sqlDb, err := db.DB()
	if err != nil {
		fmt.Println("No database found")
		panic(err)
	}

	//Seeding data
	seed := seeds.InitialSeeds(db, ds.JsonSource)
	seed.SeedUsers()
	seed.SeedBelanja()

	sqlDb.SetConnMaxIdleTime(ds.Database.ConnectionTimeout)
	sqlDb.SetMaxIdleConns(ds.Database.MaxIdleConnection)
	sqlDb.SetMaxOpenConns(ds.Database.MaxOpenConnection)

	return &Database{db}
}
