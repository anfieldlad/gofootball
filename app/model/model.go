package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // GORM MySQL
)

// Club model
type Club struct {
	gorm.Model
	Status bool   `json:"status"`
	Name   string `json:"name"`
	League string `json:"league"`
}

// Player model
type Player struct { // player model
	gorm.Model
	Status   bool   `json:"status"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Club     Club   `gorm:"foreignkey:ClubId"` // use ClubId as foreign key
	ClubID   uint   `json:"clubID"`
}

// Disable Club
func (c *Club) Disable() {
	c.Status = false
}

// Enable Club
func (c *Club) Enable() {
	c.Status = true
}

// Disable Player
func (p *Player) Disable() {
	p.Status = false
}

// Enable Player
func (p *Player) Enable() {
	p.Status = true
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Club{})
	db.AutoMigrate(&Player{})
	db.Model(&Player{}).AddForeignKey("club_id", "clubs(id)", "RESTRICT", "RESTRICT")
	return db
}
