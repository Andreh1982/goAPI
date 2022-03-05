package database

import (
	"goAPI/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConectaComBancoDeDados() {
	host := "localhost"
	port := "5432"
	dbname := "root"
	user := "root"
	password := "root"

	db, err := gorm.Open("postgres",
		"host="+host+
			" port="+port+
			" user="+user+
			" dbname="+dbname+
			" sslmode=disable password="+password)

	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)
	db.AutoMigrate(models.Personalidade{})
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}

func ClearTable() {
	DB.Exec("DELETE FROM root")
}

//
//	REFACTORING
//

func GetPersons(db *gorm.DB) ([]models.Personalidade, error) {
	personalidade := []models.Personalidade{}
	query := db.Select("personalidades.*")
	if err := query.Find(&personalidade).Error; err != nil {
		return personalidade, err
	}
	return personalidade, nil
}

func GetPersonByID(id string, db *gorm.DB) (models.Personalidade, bool, error) {
	b := models.Personalidade{}
	query := db.Select("personalidades.*")
	err := query.Where("personalidades.id = ?", id).First(&b).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return b, false, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return b, false, nil
	}
	return b, true, nil
}

func DeletePerson(id string, db *gorm.DB) error {
	var b models.Personalidade
	if err := db.Where("id = ? ", id).Delete(&b).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePerson(db *gorm.DB, b *models.Personalidade) error {
	if err := db.Save(&b).Error; err != nil {
		return err
	}
	return nil
}
