package models

import "time"

type Instructor struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Expertise string    `json:"expertise"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Course struct {
	ID           string     `json:"id" gorm:"primaryKey"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Content      string     `json:"content"`
	InstructorID string     `json:"instructor_id"`
	Instructor   Instructor `json:"instructor" gorm:"foreignKey:InstructorID;references:ID"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
