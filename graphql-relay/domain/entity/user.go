package entity

type User struct {
	ID        string
	Name      string
	IsActive  bool `gorm:"default:1"`
	CreatedAt int64
	UpdatedAt int64
}

type Users []*User
