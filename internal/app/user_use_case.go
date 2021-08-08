// Package app use cases
package app

import (
	"log"

	"github.com/julioc98/cleanarch/internal/domain"
	"github.com/julioc98/cleanarch/internal/infra/repository"
)

type storager interface {
	Store(user *domain.User) (*domain.User, error)
}

type encrypter interface {
	Encrypt(s string) string
}

type authentifier interface {
	GenerateToken(user *domain.User) (string, error)
}

type checker interface {
	Struct(s interface{}) error
}

type messenger interface {
	Send(recipient, msg string) error
}

// UserUseCase user auth uses case.
type UserUseCase struct {
	repository storager
	encrypter  encrypter
	auth       authentifier
	validate   checker
	message    messenger
}

// NewUserUseCase factory.
func NewUserUseCase(s storager, e encrypter, a authentifier, v checker, m messenger) *UserUseCase {
	return &UserUseCase{
		repository: s,
		encrypter:  e,
		auth:       a,
		validate:   v,
		message:    m,
	}
}

// SignUp create a new user.
func (u *UserUseCase) SignUp(user *domain.User) (*domain.User, error) {
	log.Println(repository.UserGorm{})

	err := u.validate.Struct(user)
	if err != nil {
		return nil, ErrInvalid
	}

	user.Password = u.encrypter.Encrypt(user.Password)

	newUser, err := u.repository.Store(user)
	if err != nil {
		return nil, ErrOnSave
	}

	token, err := u.auth.GenerateToken(newUser)
	if err != nil {
		return nil, ErrOnGenerateToken
	}

	err = u.message.Send(newUser.Email, token)
	if err != nil {
		return nil, ErrOnSendMessage
	}

	return newUser, nil
}
