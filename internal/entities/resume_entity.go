package entities

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Resume struct {
	ID          string        `gorm:"primaryKey;type:varchar(36)"`
	UserID      string        `gorm:"not null;type:varchar(36)"`
	Title       string        `gorm:"not null;type:varchar(255)"`
	Email       string        `gorm:"not null;type:varchar(255)"`
	Phone       string        `gorm:"not null;type:varchar(20)"`
	Skills      StringArray   `gorm:"type:json"`
	Experiences []*Experience `gorm:"foreignKey:ResumeID;constraint:OnDelete:CASCADE"`
	Educations  []*Education  `gorm:"foreignKey:ResumeID;constraint:OnDelete:CASCADE"`
}

func NewResumeEntity(id, userID, title, email, phone string, skills []string) *Resume {
	return &Resume{
		ID:     id,
		UserID: userID,
		Title:  title,
		Email:  email,
		Phone:  phone,
		Skills: skills,
	}
}

type Experience struct {
	ID          string     `gorm:"primaryKey;type:varchar(36)"`
	ResumeID    string     `gorm:"type:varchar(36);not null;index"`
	Company     string     `gorm:"not null;type:varchar(255)"`
	Position    string     `gorm:"not null;type:varchar(255)"`
	IsPresent   bool       `gorm:"not null;default:false"`
	StartDate   time.Time  `gorm:"not null;type:date"`
	EndDate     *time.Time `gorm:"type:date;default:null"`
	Description string     `gorm:"type:text"`
}

type Education struct {
	ID        string     `gorm:"primaryKey;type:varchar(36)"`
	ResumeID  string     `gorm:"type:varchar(36);not null;index"`
	School    string     `gorm:"type:varchar(255);not null"`
	Major     string     `gorm:"type:varchar(255);not null"`
	Degree    string     `gorm:"type:varchar(50);not null"`
	StartDate time.Time  `gorm:"not null;type:date"`
	EndDate   *time.Time `gorm:"type:date;default:null"`
}

type ResumeQueryParam struct {
}

func NewExperienceEntity(id string, resumeId string, company string, position string, present bool, startDate string, endDate string, description string) *Experience {
	start, _ := time.Parse("2006-01", startDate)
	end := parseNullableTime(endDate)
	entity := &Experience{
		ID:          id,
		ResumeID:    resumeId,
		Company:     company,
		Position:    position,
		IsPresent:   present,
		StartDate:   start,
		EndDate:     end,
		Description: description,
	}
	return entity
}

func NewEducationEntity(id string, resumeId string, school string, major string, degree string, startDate string, endDate string) *Education {
	start, _ := time.Parse("2006-01", startDate)
	end := parseNullableTime(endDate)
	entity := &Education{
		ID:        id,
		ResumeID:  resumeId,
		School:    school,
		Major:     major,
		Degree:    degree,
		StartDate: start,
		EndDate:   end,
	}
	return entity
}

type StringArray []string

func (s *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, s)
}

func (s StringArray) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func parseNullableTime(dateStr string) *time.Time {
	if dateStr == "" {
		return nil
	}
	parsedTime, err := time.Parse("2006-01", dateStr)
	if err != nil {
		return nil
	}
	return &parsedTime
}
