package Model

import (
	"liamelior-api/Database"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Photo string `form:"photo" json:"photo" binding:"required"`
	Context string `form:"context" json:"context" binding:"required"`
}

func (p *Photo) Save() (*Photo, error) {
	var err error
	err = Database.Database.Create(&p).Error
	if err != nil {
		return &Photo{}, err
	}
	return p, nil
}

func FindPhotoByContext(context string) (Photo, error) {
	var photo Photo
	err := Database.Database.Where("context = ?", context).Find(&photo).Error
	if err != nil {
		return Photo{}, err
	}
	return photo, nil
}