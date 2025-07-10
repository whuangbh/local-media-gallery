package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Image struct {
	Id   uint   `json:"-" gorm:"primarykey;autoIncrement"`
	Name string `json:"-" gorm:"not null;size:256;index:idx_member,priority:2"`
	Path string `json:"-" gorm:"not null;unique;size:512"`
	Url  string `json:"url" gorm:"not null:unique;size:512"`

	Width  int `json:"width" gorm:"not null"`
	Height int `json:"height" gorm:"not null"`

	DirectoryId uint      `json:"-" gorm:"index:idx_member,priority:1"`
	Directory   Directory `json:"-" gorm:"foreignkey:DirectoryId"`
}

func GetImagesByParentDirId(db *gorm.DB, parentDirId *uint) ([]*Image, error) {
	if parentDirId == nil {
		return []*Image{}, fmt.Errorf("GetDirectoriesByParentDirId error: parentDirId is nil")
	}

	var images []*Image

	thumbnailImageIDsQuery := db.Table("thumbnails").Select("image_id")

	err := db.
		Where("directory_id = ?", parentDirId).
		Not("id IN (?)", thumbnailImageIDsQuery).
		Order("name").
		Find(&images).
		Error
	if err != nil {
		return images, err
	}

	return images, nil
}
