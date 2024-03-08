package entadp

import (
	"cafe-management/ent"
	"cafe-management/ent/user"
	"cafe-management/pkg/repository/dto"
	"cafe-management/pkg/repository/helper"
	"context"
	"fmt"
	"time"
)

var _ UserRepository = (*User)(nil)

func NewUserRepository(DbClient *ent.Client) *User {
	return &User{
		DbClient: DbClient,
	}
}

type User struct {
	DbClient *ent.Client
}

func (u *User) CreateUser(ctx context.Context, c *dto.User) error {
	_, err := u.DbClient.User.Create().
		SetUserID(c.UserID).SetAvatar(c.Avatar).
		SetFirstName(c.FirstName).
		SetLastName(c.LastName).
		SetUserName(c.UserName).
		SetPassword(c.Password).
		SetPhone(c.Phone).
		SetConfirmPassword(c.ConfirmPassword).
		SetEmail(c.Email).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("create user/repository:%w", err)
	}

	return nil
}

func (u *User) DeleteUser(ctx context.Context, id int) error {
	if err := u.DbClient.User.DeleteOneID(id).Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return fmt.Errorf("delete user/repository:%w", ErrorNotFound)
		}
	}

	return nil
}

func (u *User) GetById(ctx context.Context, id int) (*dto.User, error) {
	var (
		user *ent.User
		err  error
	)

	user, err = u.DbClient.User.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("get user/repository:%w", ErrorNotFound)
		}
	}

	return helper.DbUserTo(user), nil
}

func (u *User) GetByUserName(ctx context.Context, name string) (*dto.User, error) {
	var (
		dbUser *ent.User
		err    error
	)

	if dbUser, err = u.DbClient.User.Query().Where(user.UserNameEQ(name)).First(ctx); err != nil {
		return nil, fmt.Errorf("get by user name /repository:%w", err)
	}

	return helper.DbUserTo(dbUser), nil

}

func (u *User) GetByUserEmail(ctx context.Context, email string) (*dto.User, error) {
	var (
		dbUser *ent.User
		err    error
	)

	if dbUser, err = u.DbClient.User.Query().Where(user.EmailEQ(email)).First(ctx); err != nil {
		return nil, fmt.Errorf("get user by user email/repository:%w", err)
	}

	return helper.DbUserTo(dbUser), nil
}

func (u *User) UpdateUser(ctx context.Context, id int, c *dto.User) error {
	_, err := u.DbClient.User.UpdateOneID(id).
		SetAvatar(c.Avatar).
		SetFirstName(c.FirstName).
		SetLastName(c.LastName).
		SetUserName(c.UserName).
		SetPassword(c.Password).
		SetEmail(c.Email).
		SetPhone(c.Phone).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("update user/repository:%w", err)
	}

	return nil
}

func (u *User) ListUser(ctx context.Context) ([]*ent.User, error) {
	var (
		users []*ent.User
		err   error
	)

	users, err = u.DbClient.User.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("list user/repository:%w", err)
	}

	return users, nil
}
