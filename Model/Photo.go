package Model

import (
	"liamelior-api/Database"

	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Photo   string `form:"photo" json:"photo" binding:"required"`
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

func FindPhotosByContext(context string) ([]Photo, error) {
	var photos []Photo
	err := Database.Database.Where("context = ?", context).Find(&photos).Error
	if err != nil {
		return nil, err
	}
	return photos, nil
}

func FindPhotosByContextWithParam(context string, limit int) ([]Photo, error) {
	var photos []Photo
	err := Database.Database.Where("context = ?", context).Limit(limit).Find(&photos).Error
	if err != nil {
		return nil, err
	}
	return photos, nil
}
