package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Person struct {
	Name string
	Age  int
}

func getInitialization() []Person {
	my_data := []Person{
		{Name: "fares", Age: 25},
		{Name: "Islem", Age: 24},
		{Name: "Youssef", Age: 16},
		{Name: "Imen", Age: 21},
	}
	return my_data
}
func StartDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("person.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect !!")
	}
	return db
}

// get all persons
func GetAllPersons() []Person {
	db := StartDb()
	var persons []Person
	db.Find(&persons)

	return persons
}

// create a person record
func MakeRecord(p *Person) {
	db := StartDb()
	db.Create(p)
}

// initialize the database
func Seed(db *gorm.DB) {
	records := getInitialization()
	for _, element := range records {
		db.Create(element)
	}
}

// load the database
func LoadDatabase() {
	db := StartDb()
	db.AutoMigrate(&Person{})
	Seed(db)
}
