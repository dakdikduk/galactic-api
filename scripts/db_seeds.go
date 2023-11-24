package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Spacecraft model
type Spacecraft struct {
	ID        uint       `gorm:"primary_key"`
	Name      string     `gorm:"not null"`
	Class     string     `gorm:"not null"`
	Armament  []Armament `gorm:"type:JSON;not null"`
	Crew      int        `gorm:"not null"`
	Image     string     `gorm:"default:null"`
	Value     float64    `gorm:"not null"`
	Status    string     `gorm:"not null"`
	CreatedAt time.Time  `gorm:"default:current_timestamp"`
	UpdatedAt time.Time  `gorm:"default:current_timestamp"`
}

// Armament model
type Armament struct {
	Title string `json:"title"`
	Qty   string `json:"qty"`
}

func main() {
	//TODO: use tidyer / config for credentials.
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/galactic?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	// AutoMigrate will create the table based on the model
	db.AutoMigrate(&Spacecraft{})

	// Seed the database with 100 random entries
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 100; i++ {
		spacecraft := generateRandomSpacecraft()
		db.Create(&spacecraft)
	}

	fmt.Println("Seeding completed.")
}

func generateRandomSpacecraft() Spacecraft {
	names := []string{"Devastator", "Voyager", "Falcon", "Enterprise", "Interceptor", "Phoenix"}
	classes := []string{"Star Destroyer", "Transport Ship", "Frigate", "Cruiser"}
	statuses := []string{"Operational", "Under Maintenance", "Out of Service"}
	armaments := []Armament{
		{Title: "Turbo Laser", Qty: "60"},
		{Title: "Ion Cannons", Qty: "60"},
		{Title: "Tractor Beam", Qty: "10"},
	}

	return Spacecraft{
		Name:     names[rand.Intn(len(names))],
		Class:    classes[rand.Intn(len(classes))],
		Armament: armaments,
		Crew:     rand.Intn(50000) + 1000,
		Image:    "https://url.to.image",
		Value:    float64(rand.Intn(1000)+500) + rand.Float64(),
		Status:   statuses[rand.Intn(len(statuses))],
	}
}
