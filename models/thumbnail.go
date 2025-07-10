package models

type Thumbnail struct {
	Id uint `json:"-" gorm:"primarykey;autoIncrement"`

	ImageID uint  `json:"-" gorm:"not null"`
	Image   Image `json:"image" gorm:"foreignkey:ImageID"`
}
