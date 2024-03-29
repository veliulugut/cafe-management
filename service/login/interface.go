package login

import "context"

type Service interface {
	Login(ctx context.Context, email, password string) (string, error)
	Register(ctx context.Context, firstName, lastName, email, password, confirmPassword string) (string, error)
}
