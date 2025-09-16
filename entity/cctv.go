package entity

import (
	"gorm.io/gorm"
)

type CCTV struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	NamaCCTV string `gorm:"size:255;not null;" json:"nama_cctv"`
	Objek    uint   `gorm:"not null;" json:"objek"`
}
