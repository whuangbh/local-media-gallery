package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Video struct {
	Id   uint   `json:"-" gorm:"primarykey;autoIncrement"`
	Name string `json:"name" gorm:"not null"`
	Path string `json:"-" gorm:"not null;unique;size:512"`
	Url  string `json:"url" gorm:"not null:unique;size:512"`

	DirectoryId uint      `json:"-"`
	Directory   Directory `json:"-" gorm:"foreignkey:DirectoryId"`

	ThumbnailId *uint      `json:"-"`
	Thumbnail   *Thumbnail `json:"thumbnail" gorm:"foreignkey:ThumbnailId"`
}

func GetVideosByParentDirId(db *gorm.DB, parentDirId *uint) ([]*Video, error) {
	if parentDirId == nil {
		return []*Video{}, fmt.Errorf("GetDirectoriesByParentDirId error: parentDirId is nil")
	}

	var videos []*Video

	err := db.Where("directory_id = ?", parentDirId).Preload("Thumbnail.Image").Find(&videos).Error
	if err != nil {
		return videos, err
	}

	return videos, nil
}
