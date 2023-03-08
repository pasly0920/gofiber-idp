package model

type User struct {
	ID           uint64 `gorm:"primaryKey"`
	Username     string `gorm:"not null"`
	Name         string `gorm:"not null"`
	Nickname     string `gorm:"not null"`
	AvatarID     uint64
	PhoneNumber  string
	StudentYear  int
	StudentGroup string
	MajorCode    int
}

// user table name user to users

func (u *User) TableName() string {
	return "users"
}
