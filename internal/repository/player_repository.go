package repository

import (
	"github.com/jinzhu/gorm"
	"gofootball/internal/model"
)

// PlayerRepository defines operations for working with players.
type PlayerRepository interface {
	FindAll() ([]model.Player, error)
	FindByName(name string) (*model.Player, error)
	Create(player *model.Player) error
	Update(player *model.Player) error
	Delete(player *model.Player) error
}

// GormPlayerRepository is a gorm implementation of PlayerRepository.
type GormPlayerRepository struct {
	DB *gorm.DB
}

func (r *GormPlayerRepository) FindAll() ([]model.Player, error) {
	var players []model.Player
	err := r.DB.Find(&players).Error
	return players, err
}

func (r *GormPlayerRepository) FindByName(name string) (*model.Player, error) {
	player := &model.Player{}
	if err := r.DB.First(player, model.Player{Name: name}).Error; err != nil {
		return nil, err
	}
	club := &model.Club{}
	if err := r.DB.Where("id = ?", player.ClubID).First(club).Error; err != nil {
		return nil, err
	}
	player.Club = *club
	return player, nil
}

func (r *GormPlayerRepository) Create(player *model.Player) error {
	return r.DB.Save(player).Error
}

func (r *GormPlayerRepository) Update(player *model.Player) error {
	return r.DB.Save(player).Error
}

func (r *GormPlayerRepository) Delete(player *model.Player) error {
	return r.DB.Delete(player).Error
}
