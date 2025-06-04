package repository

import (
	"github.com/jinzhu/gorm"
	"gofootball/internal/model"
)

// ClubRepository defines operations for working with clubs.
type ClubRepository interface {
	FindAll() ([]model.Club, error)
	FindByName(name string) (*model.Club, error)
	Create(club *model.Club) error
	Update(club *model.Club) error
	Delete(club *model.Club) error
}

// GormClubRepository is a gorm implementation of ClubRepository.
type GormClubRepository struct {
	DB *gorm.DB
}

func (r *GormClubRepository) FindAll() ([]model.Club, error) {
	var clubs []model.Club
	err := r.DB.Find(&clubs).Error
	return clubs, err
}

func (r *GormClubRepository) FindByName(name string) (*model.Club, error) {
	club := &model.Club{}
	if err := r.DB.First(club, model.Club{Name: name}).Error; err != nil {
		return nil, err
	}
	return club, nil
}

func (r *GormClubRepository) Create(club *model.Club) error {
	return r.DB.Save(club).Error
}

func (r *GormClubRepository) Update(club *model.Club) error {
	return r.DB.Save(club).Error
}

func (r *GormClubRepository) Delete(club *model.Club) error {
	return r.DB.Delete(club).Error
}
