package dto

type InstructorResponse struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Expertise string  `json:"expertise"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	Token     *string `json:"token,omitempty"`
}

type InstructorCreateRequest struct {
	Name      string `json:"name" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	Expertise string `json:"expertise" validate:"required"`
}

type InstructorUpdateRequest struct {
	Name      string `json:"name" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password,omitempty" validate:"omitempty,min=6"`
	Expertise string `json:"expertise" validate:"required"`
}

type InstructorLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}