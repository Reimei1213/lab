package entity

type CardStatus int32

const (
	CardStatusTodo CardStatus = iota
	CardStatusDoing
	CardStatusDone
)

type Card struct {
	ID        string
	Title     string
	Status    CardStatus
	UserID    *string
	User      *User `gorm:"foreignKey:UserID"`
	IsActive  bool  `gorm:"default:1"`
	CreatedAt int64
	UpdatedAt int64
}

type Cards []*Card
