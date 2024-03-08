package helper

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
)

func DbUserTo(d *ent.User) *dto.User {
	return &dto.User{
		UserID:          d.ID,
		Avatar:          d.Avatar,
		FirstName:       d.FirstName,
		LastName:        d.LastName,
		UserName:        d.UserName,
		Password:        d.Password,
		Phone:           d.Phone,
		Email:           d.Email,
		ConfirmPassword: d.ConfirmPassword,
	}
}
