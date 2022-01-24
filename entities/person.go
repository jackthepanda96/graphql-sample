package entities

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Nama     string
	HP       string
	Umur     int
	Password string
	Books    []Book `gorm:"foreignKey:Author"`
}
