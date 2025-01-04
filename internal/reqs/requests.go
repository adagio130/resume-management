package reqs

import "resume/internal/models"

// CreateUserRequest represents a request to create a user.
type CreateUserRequest struct {
	// Name is the name of the user.
	Name     string `json:"name"`
	Account  string `json:"account"`
	Gender   string `json:"gender"`
	Location string `json:"location"`
}

// GetUserRequest represents a request to get a user.
type GetUserRequest struct {
	// ID is the ID of the user.
	ID string `json:"id"`
}

// CreateResumeRequest represents a request to create a resume.
type CreateResumeRequest struct {
	// UserID is the ID of the user.
	UserID string `json:"user_id"`
	// Title is the name of the resume.
	Title string `json:"title"`
	// Email is the email of the resume.
	Email string `json:"email"`
	// Phone is the phone number of the resume.
	Phone string `json:"phone"`
	// Experience is the experience of the resume.
	Experience []*Experience `json:"experience"`
	// Skills is the skills of the resume.
	Skills []string `json:"skills"`
	// Education is the education of the resume.
	Education []*Education `json:"education"`
}

type Experience struct {
	// Company is the company of the experience.
	Company string `json:"company"`
	// Position is the position of the experience.
	Position string `json:"position"`
	// IsPresent is the present status of the experience.
	IsPresent bool `json:"is_present"`
	// StartDate is the start date of the experience.
	StartDate string `json:"start_date"`
	// EndDate is the end date of the experience.
	EndDate string `json:"end_date"`
	// Description is the description of the experience.
	Description string `json:"description"`
}

type Education struct {
	// School is the school of the education.
	School string `json:"school"`
	// Major is the major of the education.
	Major string `json:"major"`
	// Degree is the degree of the education.
	Degree string `json:"degree"`
	// StartDate is the start date of the education.
	StartDate string `json:"start_date"`
	// EndDate is the end date of the education.
	EndDate string `json:"end_date"`
}

// GetResumeRequest represents a request to get a resume.
type GetResumeRequest struct {
	// ID is the ID of the resume.
	ID string `json:"id"`
}

// UpdateResumeRequest represents a request to update a resume.
type UpdateResumeRequest struct {
	// ID is the ID of the resume.
	ID string `json:"id"`
	// Name is the name of the resume.
	UpdateRequest *CreateResumeRequest `json:"update_request"`
}

// DeleteResumeRequest represents a request to delete a resume.
type DeleteResumeRequest struct {
	// ID is the ID of the resume.
	ID string `json:"id"`
}

// ListResumesRequest represents a request to list resumes.
type ListResumesRequest struct {
	// UserID is the ID of the user.
	UserID string `json:"user_id"`
}

// ListResumesResponse represents a response to list resumes.
type ListResumesResponse struct {
	// Resumes is the list of resumes.
	Resumes []*models.Resume `json:"resumes"`
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

// GetResumeResponse represents a response to get a resume.
type GetResumeResponse struct {
	// Resume is the resume.
	Resume *models.Resume `json:"resume"`
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

// GetUserResponse represents a response to get a user.
type GetUserResponse struct {
	// User is the user.
	User *models.User `json:"user"`
}
