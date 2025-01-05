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
	ID string `json:"id,omitempty"`
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
	ID string `json:"id,omitempty"`
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

// UpdateResumeRequest represents a request to update a resume.
type UpdateResumeRequest struct {
	// ID is the ID of the resume.
	ID string `json:"id" validate:"required`
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
