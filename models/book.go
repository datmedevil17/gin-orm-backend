package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title" binding:"required"`
	Author      string `gorm:"not null" json:"author" binding:"required"`
	Price       int    `gorm:"not null" json:"price" binding:"required"`
	PublishedAt string `gorm:"not null" json:"published_at" binding:"required"`
}
