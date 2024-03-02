package user

import (
	"cafe-management/ent"
	"cafe-management/pkg/passwd"
	"cafe-management/pkg/repository/dto"
	"cafe-management/pkg/repository/entadp"
	"cafe-management/pkg/validator"
	"context"
	"fmt"
	"time"
)

var _ ServiceUser = (*User)(nil)

func New(repo entadp.RepositoryInterface, pw passwd.Interface) *User {
	return &User{
		repo: repo,
		pw:   pw,
	}
}

type User struct {
	repo entadp.RepositoryInterface
	pw   passwd.Interface
}

func (u *User) CreateUser(ctx context.Context, c *CreateUserModel) error {
	var err error

	emailVld := validator.IsEmailValid{Email: c.Email}
	passVld := validator.IsInRange{
		Text: c.Password,
		Max:  32,
		Min:  6,
	}

	if err = validator.ValidateAll(&emailVld, &passVld); err != nil {
		return fmt.Errorf("create user :%w", err)
	}

	var hashedPW string

	if hashedPW, err = u.pw.Generate(c.Password); err != nil {
		return fmt.Errorf("create user / generate password :%w", err)
	}

	userDTO := dto.User{
		Email:    c.Email,
		Password: hashedPW,
		UserName: c.UserName,
	}

	if err = u.repo.User().CreateUser(ctx, &userDTO); err != nil {
		return fmt.Errorf("create user/service:%w", err)
	}

	return nil
}

func (u *User) DeleteUser(ctx context.Context, id int) error {
	if err := u.repo.User().DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("delete user/service:%w", err)
	}

	return nil
}

func (u *User) GetById(ctx context.Context, id int) (*UserModel, error) {
	var (
		user *dto.User
		err  error
	)

	if user, err = u.repo.User().GetById(ctx, id); err != nil {
		return nil, fmt.Errorf("get user/service:%w", err)
	}

	return &UserModel{
		UserName:  user.UserName,
		Email:     user.Email,
		Phone:     user.Phone,
		Avatar:    user.Avatar,
		CreatedAt: time.Now(),
	}, nil
}

func (u *User) GetByUserName(ctx context.Context, name string) (*UserModel, error) {
	var (
		user *dto.User
		err  error
	)

	if user, err = u.repo.User().GetByUserName(ctx, name); err != nil {
		return nil, fmt.Errorf("get user/service:%w", err)
	}

	return &UserModel{
		UserName:  user.UserName,
		Email:     user.Email,
		Phone:     user.Phone,
		Avatar:    user.Avatar,
		CreatedAt: time.Now(),
	}, nil
}

func (u *User) ListUser(ctx context.Context) ([]*UserModel, error) {
	var (
		users []*ent.User
		err   error
	)

	if users, err = u.repo.User().ListUser(ctx); err != nil {
		return nil, fmt.Errorf("list user/service:%w", err)
	}

	var userModels []*UserModel
	for _, user := range users {
		userModels = append(userModels, &UserModel{
			UserName:  user.UserName,
			Email:     user.Email,
			Phone:     user.Phone,
			Avatar:    user.Avatar,
			CreatedAt: time.Now(),
		})
	}

	return userModels, nil

}

func (u *User) UpdateUser(ctx context.Context, id int, c *UpdateUserModel) error {
	updateDTO := dto.User{
		Email:     c.Email,
		UserName:  c.UserName,
		Phone:     c.Phone,
		Avatar:    c.Avatar,
		UpdatedAt: time.Now(),
	}

	if err := u.repo.User().UpdateUser(ctx, id, &updateDTO); err != nil {
		return fmt.Errorf("update user/service:%w", err)
	}

	return nil
}
