package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Directory struct {
	Id   uint   `json:"-" gorm:"primarykey;autoIncrement"`
	Name string `json:"name" gorm:"not null;size:256;index:idx_member,priority:2"`
	Path string `json:"-" gorm:"not null;unique;size:512"`
	Url  string `json:"url" gorm:"not null;unique;size:512"`

	ParentDirectoryId *uint      `json:"-" gorm:"index:idx_member,priority:1"`
	ParentDirectory   *Directory `json:"-" gorm:"foreignkey:ParentDirectoryId"`

	ThumbnailId *uint      `json:"-"`
	Thumbnail   *Thumbnail `json:"thumbnail" gorm:"foreignkey:ThumbnailId"`
}

func GetDirectoryByUrl(db *gorm.DB, url string) *Directory {
	var parentDirectory Directory

	db.
		Where("url = ?", url).
		First(&parentDirectory)

	return &parentDirectory
}

func GetDirectoriesByParentDirId(db *gorm.DB, parentDirId *uint) ([]*Directory, error) {
	if parentDirId == nil {
		return []*Directory{}, fmt.Errorf("GetDirectoriesByParentDirId error: parentDirId is nil")
	}

	var directories []*Directory

	err := db.
		Where("parent_directory_id = ?", parentDirId).
		Order("name").
		Preload("Thumbnail.Image").
		Find(&directories).
		Error
	if err != nil {
		return directories, err
	}

	return directories, nil
}
