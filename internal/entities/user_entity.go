package entities

type User struct {
	ID       string `gorm:"primaryKey;type:varchar(36)"`
	Account  string `gorm:"not null;unique;type:varchar(255)"`
	Name     string `gorm:"not null;type:varchar(255)"`
	Gender   string `gorm:"not null;type:varchar(10)"`
	Location string `gorm:"not null;type:varchar(255)"`
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
