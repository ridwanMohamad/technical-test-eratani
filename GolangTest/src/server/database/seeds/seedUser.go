package seeds

import (
	"encoding/json"
	"golangtest/src/constanta"
	"golangtest/src/models"
	"io/ioutil"
	"log"
)

func (seed *seeds) SeedUsers() error {
	var model *models.Users
	db := seed.db
	if !isTableEmpty(db, &model) {
		log.Println("Skipping seeding belanja as the table is not empty.")
		return nil
	}
	data, err := ioutil.ReadFile(seed.json["users"])
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
	}

	// Unmarshal the JSON data into an array of Person struct
	var users []*models.Users
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Fatal("Error unmarshaling JSON data:", err)
	}

	err = db.Model(&model).CreateInBatches(users, constanta.BulkChunkSize).Error
	if err != nil {
		log.Fatal("Error insert data:", err)
	}
	return nil
}
