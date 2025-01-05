package models

import (
	"github.com/google/uuid"
	"resume/internal/reqs"
)

type Resume struct {
	ID         string                 `json:"id"`
	UserID     string                 `json:"user_id,omitempty"`
	Title      string                 `json:"title"`
	Email      string                 `json:"email,omitempty"`
	Phone      string                 `json:"phone,omitempty"`
	Experience map[string]*Experience `json:"experience,omitempty"`
	Skills     []string               `json:"skills,omitempty"`
	Education  map[string]*Education  `json:"education,omitempty"`
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

func NewResumeFromReqs(resumeId, userID, title, email, phone string, experience []*reqs.Experience, skills []string, education []*reqs.Education) *Resume {
	if resumeId == "" {
		resumeId = uuid.New().String()
	}
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
	expId := experience.ID
	if expId == "" {
		expId = uuid.New().String()
	}
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
	eduId := education.ID
	if eduId == "" {
		eduId = uuid.New().String()
	}
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

func (r *Resume) UpdateExperience(experience *reqs.Experience) {
	exp := r.Experience[experience.ID]
	if exp == nil {
		return
	}
	exp.Company = experience.Company
	exp.Position = experience.Position
	exp.IsPresent = experience.IsPresent
	exp.StartDate = experience.StartDate
	exp.EndDate = experience.EndDate
	exp.Description = experience.Description
}

func (r *Resume) UpdateEducation(education *reqs.Education) {
	edu := r.Education[education.ID]
	if edu == nil {
		return
	}
	edu.School = education.School
	edu.Major = education.Major
	edu.Degree = education.Degree
	edu.StartDate = education.StartDate
	edu.EndDate = education.EndDate
}
