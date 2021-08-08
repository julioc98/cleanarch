// Package domain entities
package domain

import (
	"testing"
	"time"
)

func TestUser_Age(t *testing.T) {
	type fields struct {
		Name      string
		Email     string
		Password  string
		BirthDate time.Time
	}

	type args struct {
		t time.Time
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Age 25 OK",
			fields: fields{
				BirthDate: time.Date(1996, 4, 15, 0, 0, 0, 0, time.UTC),
			},
			args: args{
				t: time.Date(2021, 7, 25, 0, 0, 0, 0, time.UTC),
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Name:      tt.fields.Name,
				Email:     tt.fields.Email,
				Password:  tt.fields.Password,
				BirthDate: tt.fields.BirthDate,
			}
			if got := u.Age(tt.args.t); got != tt.want {
				t.Errorf("User.Age() = %v, want %v", got, tt.want)
			}
		})
	}
}
