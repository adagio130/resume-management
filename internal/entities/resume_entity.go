package entities

import (
	"time"
)

type Resume struct {
	ID         string        `gorm:"primaryKey;"`
	UserID     string        `gorm:"not null"`
	Title      string        `gorm:"not null"`
	Email      string        `gorm:"not null;size:255"`
	Phone      string        `gorm:"size:20"`
	Location   string        `gorm:"size:255"`
	Experience []*Experience `gorm:"foreignKey:ResumeID;constraint:OnDelete:CASCADE"`
	Skills     []string      `gorm:"type:json"`
	Education  []*Education  `gorm:"foreignKey:ResumeID;constraint:OnDelete:CASCADE"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Experience struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement"`
	ResumeID    uint64    `gorm:"not null"`
	Company     string    `gorm:"not null;size:255"`
	Position    string    `gorm:"not null;size:255"`
	IsPresent   bool      `gorm:"not null"`
	StartDate   time.Time `gorm:"not null"`
	EndDate     *time.Time
	Description string `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Education struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	ResumeID  uint64    `gorm:"not null"`
	School    string    `gorm:"not null;size:255"`
	Major     string    `gorm:"size:255"`
	Degree    string    `gorm:"size:50"`
	StartDate time.Time `gorm:"not null"`
	EndDate   *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ResumeQueryParam struct {
}
