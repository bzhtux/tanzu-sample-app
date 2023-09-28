package utils

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB(sqlite_file string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(sqlite_file), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to DB: %s", err.Error())
		return nil
	} else {
		log.Printf("Connected to DB: %v", db)
		return db
	}
}

// func AutoMigrate(db interface{}) {

// 	db.AutoMigrate(db)
// }

// func (h Handler) GetCodes() ([]models.HttpStatusCode, error) {
// 	var codes []models.HttpStatusCode
// 	res := h.DB.Find(&codes)
// 	if res.Error != nil {
// 		log.Printf("Error fetching all HTTP Status Codes from DB: %s", res.Error)
// 		return nil, res.Error
// 	} else {
// 		log.Printf("Nb of results: %d", res.RowsAffected)
// 		return codes, nil
// 	}
// }
