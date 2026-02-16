package dto

type CourseResponse struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Slug         string `json:"slug"`
	Description  string `json:"description"`
	Content      string `json:"content"`
	Status       string `json:"status"`
	InstructorID string `json:"instructor_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type CourseWithInstructorResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	Content     string `json:"content"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`

	InstructorID       string `json:"instructor_id"`
	InstructorName     string `json:"instructor_name"`
	InstructorUsername string `json:"instructor_username"`
}

type CreateCourseRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Status      string `json:"status" binding:"required,oneof=draft published"`
}

type UpdateCourseRequest struct {
	Title       string `json:"title" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Description string `json:"description" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Status      string `json:"status" binding:"required,oneof=draft published"`
}
