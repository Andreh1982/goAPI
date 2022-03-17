package database

import (
	"goAPI/models"
	"goAPI/shared"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	host := "localhost"
	port := "5432"
	dbname := "root"
	user := "root"
	password := "root"
	db, err := gorm.Open("postgres", "host="+host+" port="+port+" user="+user+" dbname="+dbname+" sslmode=disable password="+password)
	if err != nil {
		shared.LogCustom([]string{"Database Connection Failed!"}, "error")
	}
	db.LogMode(true)
	// db.AutoMigrate(models.Person{})
	DB = db
	shared.LogCustom([]string{"Database Connected!"}, "info")
}

func GetDB() *gorm.DB {
	return DB
}

func ClearTable() {
	DB.Exec("DELETE FROM root")
}

func GetPersons(db *gorm.DB) ([]models.Person, error) {
	person := []models.Person{}
	query := db.Select("people.*")
	if err := query.Find(&person).Error; err != nil {
		return person, err
	}
	return person, nil
}

func GetPersonByID(id string, db *gorm.DB) (models.Person, bool, error) {
	b := models.Person{}
	query := db.Select("people.*")
	err := query.Where("people.id = ?", id).First(&b).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return b, false, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return b, false, nil
	}
	return b, true, nil
}

func DeletePerson(id string, db *gorm.DB) error {
	var b models.Person
	if err := db.Where("id = ? ", id).Delete(&b).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePerson(db *gorm.DB, b *models.Person) error {
	if err := db.Save(&b).Error; err != nil {
		return err
	}
	return nil
}
