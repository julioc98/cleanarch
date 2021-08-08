// Package repository save data
package repository

import (
	"github.com/julioc98/cleanarch/internal/domain"

	"gorm.io/gorm"
)

// UserGorm repository.
type UserGorm struct {
	db *gorm.DB
}

// NeWUserGorm repository factory.
func NeWUserGorm(db *gorm.DB) *UserGorm {
	return &UserGorm{
		db: db,
	}
}

// Store an user.
func (g *UserGorm) Store(user *domain.User) (*domain.User, error) {
	if dbc := g.db.Create(user); dbc.Error != nil {
		return nil, dbc.Error
	}

	return user, nil
}
