package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/dakdikduk/galactic-api/config"
	"gorm.io/datatypes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Spacecraft model
type Spacecraft struct {
	ID        uint           `gorm:"primary_key"`
	Name      string         `gorm:"not null"`
	Class     string         `gorm:"not null"`
	Armament  datatypes.JSON `gorm:"type:JSON;not null"`
	Crew      int            `gorm:"not null"`
	Image     string         `gorm:"default:null"`
	Value     float64        `gorm:"not null"`
	Status    string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"default:now()"`
	UpdatedAt time.Time      `gorm:"default:now() ON UPDATE now()"`
}

// Armament model
type Armament struct {
	Title string `json:"title"`
	Qty   string `json:"qty"`
}

func main() {
	cfg := config.Get()
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.MysqlDbUser, cfg.MysqlDbPass, cfg.MysqlDbHost, cfg.MysqlDbPort, cfg.MySqlDbName)
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal(err)
	}

	// Seed the database with 100 random entries
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 5; i++ {
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
	armamentsJSON, err := json.Marshal(armaments)
	if err != nil {
		log.Fatal("Error marshaling armaments:", err)
	}

	return Spacecraft{
		Name:     names[rand.Intn(len(names))],
		Class:    classes[rand.Intn(len(classes))],
		Armament: datatypes.JSON(armamentsJSON),
		Crew:     rand.Intn(50000) + 1000,
		Image:    "https://url.to.image",
		Value:    float64(rand.Intn(1000)+500) + rand.Float64(),
		Status:   statuses[rand.Intn(len(statuses))],
	}
}
