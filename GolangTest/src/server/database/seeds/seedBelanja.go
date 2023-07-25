package seeds

import (
	"encoding/json"
	"golangtest/src/constanta"
	"golangtest/src/models"
	"io/ioutil"
	"log"
)

func (seed *seeds) SeedBelanja() error {
	var model *models.Belanja
	db := seed.db
	if !isTableEmpty(db, &model) {
		log.Println("Skipping seeding belanja as the table is not empty.")
		return nil
	}

	data, err := ioutil.ReadFile(seed.json["belanja"])
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
	}

	// Unmarshal the JSON data into an array of Person struct
	var belanja []*models.Belanja
	err = json.Unmarshal(data, &belanja)
	if err != nil {
		log.Fatal("Error unmarshaling JSON data:", err)
	}

	err = db.Model(&model).CreateInBatches(belanja, constanta.BulkChunkSize).Error
	if err != nil {
		log.Fatal("Error insert data:", err)
	}
	return nil
}
