package models

import (
	"github.com/google/uuid"
	"resume/internal/reqs"
)

type Resume struct {
	ID         string                 `json:"id"`
	UserID     string                 `json:"user_id"`
	Title      string                 `json:"title"`
	Email      string                 `json:"email"`
	Phone      string                 `json:"phone"`
	Experience map[string]*Experience `json:"experience"`
	Skills     []string               `json:"skills"`
	Education  map[string]*Education  `json:"education"`
}

type Experience struct {
	ID          string `json:"id"`
	Company     string `json:"company"`
	Position    string `json:"position"`
	IsPresent   bool   `json:"is_present"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Description string `json:"description"`
}

type Education struct {
	ID        string `json:"id"`
	School    string `json:"school"`
	Major     string `json:"major"`
	Degree    string `json:"degree"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

// NewResume creates a new resume.
func NewResumeFromReqs(userID, title, email, phone string, experience []*reqs.Experience, skills []string, education []*reqs.Education) *Resume {
	resumeId := uuid.New().String()
	resume := &Resume{
		ID:         resumeId,
		UserID:     userID,
		Title:      title,
		Email:      email,
		Phone:      phone,
		Experience: make(map[string]*Experience),
		Skills:     skills,
		Education:  make(map[string]*Education),
	}
	for _, exp := range experience {
		resume.AddExperience(exp)
	}
	for _, edu := range education {
		resume.AddEducation(edu)
	}
	return resume
}

func (r *Resume) AddExperience(experience *reqs.Experience) {
	expId := uuid.New().String()
	exp := &Experience{
		ID:          expId,
		Company:     experience.Company,
		Position:    experience.Position,
		IsPresent:   experience.IsPresent,
		StartDate:   experience.StartDate,
		EndDate:     experience.EndDate,
		Description: experience.Description,
	}
	r.Experience[expId] = exp
}

func (r *Resume) AddEducation(education *reqs.Education) {
	eduId := uuid.New().String()
	edu := &Education{
		ID:        eduId,
		School:    education.School,
		Major:     education.Major,
		Degree:    education.Degree,
		StartDate: education.StartDate,
		EndDate:   education.EndDate,
	}
	r.Education[eduId] = edu
}

func (r *Resume) UpdateExperience(id string, experience *reqs.Experience) {
	if _, ok := r.Experience[id]; ok {
		exp := &Experience{
			ID:          id,
			Company:     experience.Company,
			Position:    experience.Position,
			IsPresent:   experience.IsPresent,
			StartDate:   experience.StartDate,
			EndDate:     experience.EndDate,
			Description: experience.Description,
		}
		r.Experience[id] = exp
	} else {
		r.AddExperience(experience)
	}
}

func (r *Resume) UpdateEducation(id string, education *reqs.Education) {
	if _, ok := r.Education[id]; ok {
		edu := &Education{
			ID:        id,
			School:    education.School,
			Major:     education.Major,
			Degree:    education.Degree,
			StartDate: education.StartDate,
			EndDate:   education.EndDate,
		}
		r.Education[id] = edu
	} else {
		r.AddEducation(education)
	}
}
