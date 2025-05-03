package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type Category struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
	News []News `gorm:"foreignKey:CategoryID"`
}

type News struct {
	gorm.Model
	Title      string    `gorm:"not null"`
	Content    string    `gorm:"type:text;not null"`
	CategoryID uint      `gorm:"not null"`
	Category   Category  `gorm:"foreignKey:CategoryID"`
	UserID     uint      `gorm:"not null"`
	User       User      `gorm:"foreignKey:UserID"`
	Comments   []Comment `gorm:"foreignKey:NewsID"`
}

type CustomPage struct {
	gorm.Model
	URL     string `gorm:"unique;not null"`
	Content string `gorm:"type:text;not null"`
	UserID  uint   `gorm:"not null"`
	User    User   `gorm:"foreignKey:UserID"`
}

type Comment struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Comment string `gorm:"type:text;not null"`
	NewsID  uint   `gorm:"not null"`
	News    News   `gorm:"foreignKey:NewsID"`
}
