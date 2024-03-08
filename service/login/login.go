package login

import (
	"cafe-management/pkg/jwt"
	"cafe-management/pkg/passwd"
	"cafe-management/pkg/repository/dto"
	"cafe-management/pkg/repository/entadp"
	"context"
	"fmt"
)

var _ Service = (*Login)(nil)

func New(r entadp.RepositoryInterface, j jwt.Interface, bc passwd.Interface) *Login {
	return &Login{
		repo: r,
		jwt:  j,
		bc:   bc,
	}
}

type Login struct {
	repo entadp.RepositoryInterface
	jwt  jwt.Interface
	bc   passwd.Interface
}

func (l *Login) Login(ctx context.Context, email, password string) (string, error) {
	user, err := l.repo.User().GetByUserEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("login :%w", err)
	}

	err = l.bc.Compare(user.Password, password)
	if err != nil {
		return "", fmt.Errorf("login:%w", err)
	}

	token, err := l.jwt.Generate(int(user.UserID))
	if err != nil {
		return "", fmt.Errorf("login :%w", err)
	}

	return token, nil
}

func (l *Login) Register(ctx context.Context, firstName, lastName, email, password, confirmPassword string) (string, error) {
	user := &dto.User{
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		Password:        password,
		ConfirmPassword: confirmPassword,
	}

	_, err := l.repo.User().GetByUserName(ctx, user.UserName)
	if err == nil {
		return "", fmt.Errorf("username %s is already taken", user.UserName)
	}

	if password != confirmPassword {
		return "", fmt.Errorf("passwords don't match")
	}

	hashedPassword, err := l.bc.Generate(password)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	user.Password = hashedPassword

	err = l.repo.User().CreateUser(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	token, err := l.jwt.Generate(int(user.UserID))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}
