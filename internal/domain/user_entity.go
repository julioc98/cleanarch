// Package domain entities
package domain

import (
	"time"
)

// User entity.
type User struct {
	ID          string     `gorm:"primaryKey;unique;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name        string     `json:"name" validate:"required"`
	Email       string     `gorm:"index;not null" validate:"required,email" json:"email"`
	Password    string     `gorm:"index" validate:"required" json:"password"`
	Role        string     `gorm:"index" json:"role"`
	Token       string     `gorm:"-" json:"token"`
	BirthDate   time.Time  `gorm:"not null" validate:"required" json:"birthDate"`
	ActivatedAt *time.Time `json:"activatedAt,omitempty"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt,omitempty"`
	DeletedAt   *time.Time `gorm:"index" json:"deletedAt,omitempty"`
}

// Age from user.
func (u *User) Age(t time.Time) int {
	return t.AddDate(-u.BirthDate.Year(), -int(u.BirthDate.Month()), -u.BirthDate.Day()).Year()
}
