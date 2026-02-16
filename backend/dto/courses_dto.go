package dto

type CourseResponse struct {
    ID          uint   `json:"id"`
    Title       string `json:"title"`
    Slug        string `json:"slug"`
    Description string `json:"description"`
    Content     string `json:"content"`
    Status      string `json:"status"`
    InstructorID string `json:"instructor_id"`
}

type CreateCourseRequest struct {
    Title       string `json:"title" binding:"required"`
    Slug        string `json:"slug" binding:"required"`
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