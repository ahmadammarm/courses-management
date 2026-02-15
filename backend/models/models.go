package models

import (
	"time"

	"gorm.io/gorm"
)

type Instructor struct {
	ID        string         `json:"id" gorm:"primaryKey;type:varchar(30)"`
	Name      string         `json:"name" gorm:"type:varchar(191)"`
	Username  string         `json:"username" gorm:"type:varchar(191);uniqueIndex"`
	Email     string         `json:"email" gorm:"type:varchar(191);uniqueIndex"`
	Password  string         `json:"-"`
	Expertise string         `json:"expertise" gorm:"type:varchar(191)"`
	Courses   []Course       `json:"courses" gorm:"foreignKey:InstructorID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Course struct {
	ID           string         `json:"id" gorm:"primaryKey;type:varchar(30)"`
	Title        string         `json:"title" gorm:"type:varchar(191);index"`
	Slug         string         `json:"slug" gorm:"type:varchar(191);uniqueIndex"`
	Description  string         `json:"description" gorm:"type:text"`
	Content      string         `json:"content" gorm:"type:longtext"`
	Status       string         `json:"status" gorm:"type:varchar(50);default:'draft'"`
	InstructorID string         `json:"instructor_id" gorm:"type:varchar(30);index"`
	Instructor   Instructor     `json:"instructor" gorm:"foreignKey:InstructorID;references:ID"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}