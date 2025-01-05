package reqs

// CreateUserRequest represents a request to create a user.
type CreateUserRequest struct {
	// Name is the name of the user.
	Name     string `json:"name" validate:"required"`
	Account  string `json:"account" validate:"required"`
	Gender   string `json:"gender" validate:"required"`
	Location string `json:"location"`
}

// GetUserRequest represents a request to get a user.
type GetUserRequest struct {
	// ID is the ID of the user.
	ID string `json:"id" validate:"required"`
}

// CreateResumeRequest represents a request to create a resume.
type CreateResumeRequest struct {
	// UserID is the ID of the user.
	UserID string `json:"user_id" validate:"required"`
	// Title is the name of the resume.
	Title string `json:"title" validate:"required"`
	// Email is the email of the resume.
	Email string `json:"email" validate:"required,email"`
	// Phone is the phone number of the resume.
	Phone string `json:"phone" validate:"required"`
	// Experience is the experience of the resume.
	Experience []*Experience `json:"experience"`
	// Skills is the skills of the resume.
	Skills []string `json:"skills"`
	// Education is the education of the resume.
	Education []*Education `json:"education"`
}

type Experience struct {
	// Company is the company of the experience.
	Company string `json:"company" validate:"required"`
	// Position is the position of the experience.
	Position string `json:"position" validate:"required"`
	// IsPresent is the present status of the experience.
	IsPresent bool `json:"is_present"`
	// StartDate is the start date of the experience.
	StartDate string `json:"start_date" validate:"required"`
	// EndDate is the end date of the experience.
	EndDate string `json:"end_date"`
	// Description is the description of the experience.
	Description string `json:"description"`
}

type Education struct {
	// School is the school of the education.
	School string `json:"school" validate:"required"`
	// Major is the major of the education.
	Major string `json:"major" validate:"required"`
	// Degree is the degree of the education.
	Degree string `json:"degree" validate:"required"`
	// StartDate is the start date of the education.
	StartDate string `json:"start_date" validate:"required"`
	// EndDate is the end date of the education.
	EndDate string `json:"end_date"`
}

// GetResumeRequest represents a request to get a resume.
type GetResumeRequest struct {
	// ID is the ID of the resume.
	ID string `json:"id" validate:"required"`
}

// UpdateResumeRequest represents a request to update a resume.
type UpdateResumeRequest struct {
	// ID is the ID of the resume.
	ID string `json:"id" validate:"required`
	// Name is the name of the resume.
	UpdateRequest *CreateResumeRequest `json:"update_request" validate:"required"`
}

// DeleteResumeRequest represents a request to delete a resume.
type DeleteResumeRequest struct {
	// ID is the ID of the resume.
	ID string `json:"id" validate:"required"`
}

// ListResumesRequest represents a request to list resumes.
type ListResumesRequest struct {
	// UserID is the ID of the user.
	UserID string `json:"user_id" validate:"required"`
}

// ErrorResponse represents an error response.
type ErrorResponse struct {
	// Message is the error message.
	Message string `json:"message"`
}

// CreateResumeResponse represents a response to create a resume.
type CreateResumeResponse struct {
	// ID is the ID of the resume.
	ID string `json:"id"`
}

// UpdateResumeResponse represents a response to update a resume.
type UpdateResumeResponse struct {
	// ID is the ID of the resume.
	ID string `json:"id"`
}

// DeleteResumeResponse represents a response to delete a resume.
type DeleteResumeResponse struct {
	// ID is the ID of the resume.
	ID string `json:"id"`
}

// CreateUserResponse represents a response to create a user.
type CreateUserResponse struct {
	// ID is the ID of the user.
	ID string `json:"id"`
}
