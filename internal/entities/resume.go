package entities

import "time"

type Resume struct {
	ID         string
	UserId     string
	Name       string
	Email      string
	Phone      string
	Location   string
	Experience []*Experience
	Skills     []string
	Education  []Education
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Experience struct {
	Company     string
	Position    string
	StartDate   string
	EndDate     string
	Description string
}

type Education struct {
	School    string
	Major     string
	Degree    string
	StartDate string
	EndDate   string
}

type ResumeQueryParam struct {
	UserId    string
	Name      string
	Location  string
	Company   string
	School    string
	Major     string
	Position  string
	StartDate string
	EndDate   string
}
