// Package app use cases
package app

import (
	"errors"
	"testing"

	"github.com/julioc98/cleanarch/internal/app/mock"
	"github.com/julioc98/cleanarch/internal/domain"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	id            = "user-id"
	name          = "user"
	email         = "user@test.com"
	password      = "user_test_pass"
	encryptedPass = "encrypt"
	token         = "user-token"
)

var err = errors.New("err")

func TestUserUseCase_SignUp(t *testing.T) {
	user := domain.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	type args struct {
		user *domain.User
	}

	tests := []struct {
		name          string
		args          args
		want          *domain.User
		wantErr       bool
		repositoryErr error
		authErr       error
		validateErr   error
		msgErr        error
	}{
		{
			name: "New User Ok",

			args: args{
				user: &user,
			},
			want: &domain.User{
				ID:       id,
				Name:     name,
				Email:    email,
				Password: encryptedPass,
			},
			wantErr: false,
		},
		{
			name: "Validate ERROR",
			args: args{
				user: &user,
			},
			validateErr: err,
			wantErr:     true,
		},
		{
			name: "Repository ERROR",
			args: args{
				user: &user,
			},
			repositoryErr: err,
			wantErr:       true,
		},
		{
			name: "Auth ERROR",
			args: args{
				user: &user,
			},
			authErr: err,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			repository := mock.NewMockstorager(controller)
			encrypt := mock.NewMockencrypter(controller)
			auth := mock.NewMockauthentifier(controller)
			validator := mock.NewMockchecker(controller)
			message := mock.NewMockmessenger(controller)

			u := NewUserUseCase(repository, encrypt, auth, validator, message)

			validatorCall := validator.EXPECT().Struct(gomock.Eq(tt.args.user)).Return(tt.validateErr).AnyTimes()

			encryptCall := encrypt.EXPECT().Encrypt(gomock.Eq(tt.args.user.Password)).Return(encryptedPass).AnyTimes().After(validatorCall)

			repositoryCall := repository.EXPECT().Store(gomock.Eq(tt.args.user)).Return(tt.want, tt.repositoryErr).AnyTimes().After(encryptCall)

			authCall := auth.EXPECT().GenerateToken(gomock.Eq(tt.want)).Return(token, tt.authErr).AnyTimes().After(repositoryCall)

			message.EXPECT().Send(gomock.Eq(email), gomock.Eq(token)).Return(tt.msgErr).AnyTimes().After(authCall)

			got, err := u.SignUp(tt.args.user)
			if !assert.Equal(t, tt.wantErr, err != nil) {
				return
			}
			assert.Equal(t, tt.want, got, "they should be equal")
		})
	}
}
