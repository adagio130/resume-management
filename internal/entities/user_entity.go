package entities

type User struct {
	ID       string `gorm:"primaryKey;"`
	Account  string `gorm:"not null;unique"`
	Name     string `gorm:"not null"`
	Gender   string `gorm:"not null"`
	Location string `gorm:"not null"`
}

func NewUserEntity(userId, account, name, gender, location string) *User {
	return &User{
		ID:       userId,
		Account:  account,
		Name:     name,
		Gender:   gender,
		Location: location,
	}
}
