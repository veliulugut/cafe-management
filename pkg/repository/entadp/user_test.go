package entadp

import (
	"cafe-management/ent"
	"cafe-management/ent/enttest"
	"cafe-management/pkg/repository/dto"
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func TestUser_CreateUser(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	userRepo := NewUserRepository(client)

	user := []dto.User{
		{
			UserID:          1,
			FirstName:       "elif",
			LastName:        "test",
			Password:        "ufo361",
			ConfirmPassword: "ufo361",
			UserName:        "tiffaniyelif",
			Email:           "a2asfa@gmail.com",
			Avatar:          "link",
			Phone:           "222 222 222",
			UpdatedAt:       time.Now(),
			CreatedAt:       time.Now(),
		},
		{
			UserID:          2,
			FirstName:       "asdf",
			LastName:        "abc oglu",
			Password:        "asdfghj",
			ConfirmPassword: "asdfghj",
			UserName:        "johndoe",
			Email:           "a@gmail.com",
			Avatar:          "link",
			Phone:           "3131 31 3131",
			UpdatedAt:       time.Now(),
			CreatedAt:       time.Now(),
		},
	}

	for _, u := range user {
		err := userRepo.CreateUser(context.Background(), &u)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	}
}

func TestUser_DeleteUser(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	userRepo := NewUserRepository(client)

	user := dto.User{
		UserID:          1,
		FirstName:       "john",
		LastName:        "doe",
		Password:        "asdfghjk",
		ConfirmPassword: "asdfghjk",
		UserName:        "johndoe31",
		Email:           "a2asfa@gmail.com",
		Avatar:          "link",
		Phone:           "3131 3131",
		UpdatedAt:       time.Now(),
		CreatedAt:       time.Now(),
	}

	if err := userRepo.CreateUser(context.Background(), &user); err != nil {
		t.Error(err)
	}
	test := []struct {
		description string
		userId      int
		expectedErr error
	}{
		{
			description: "pass",
			userId:      1,
			expectedErr: nil,
		},
		{
			description: "did not pass",
			userId:      2,
			expectedErr: ErrorNotFound,
		},
	}

	for _, tt := range test {
		t.Run(fmt.Sprintf("DeleteUser_%s", tt.description), func(t *testing.T) {
			err := userRepo.DeleteUser(context.Background(), tt.userId)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error :%v but got error :%v", tt.expectedErr, err)
			}
		})
	}

}

func TestUser_GetByIdUser(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewUserRepository(client)

	user := dto.User{
		UserID:          1,
		FirstName:       "john",
		LastName:        "doe",
		Password:        "asdfghjk",
		ConfirmPassword: "asdfghjk",
		UserName:        "johndoe31",
		Email:           "a2asfa@gmail.com",
		Avatar:          "link",
		Phone:           "3131 3131",
		UpdatedAt:       time.Now(),
		CreatedAt:       time.Now(),
	}

	t.Run("CreateUser", func(t *testing.T) {
		if err := repo.CreateUser(context.Background(), &user); err != nil {
			t.Error(err)
		}
	})

	test := []struct {
		description string
		userID      int
		expectedErr error
	}{
		{
			description: "pass",
			userID:      1,
			expectedErr: nil,
		},
		{
			description: "did not passed",
			userID:      31,
			expectedErr: ErrorNotFound,
		},
	}

	for _, tt := range test {
		t.Run(fmt.Sprintf("GetByUserId_%s", tt.description), func(t *testing.T) {
			_, err := repo.GetById(context.Background(), tt.userID)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error :%v but got error :%v", tt.expectedErr, err)
			}
		})
	}

}

func TestUser_UpdateUser(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewUserRepository(client)

	_, err := repo.DbClient.User.Create().
		SetUserID(1).
		SetUserName("testadmin").
		SetFirstName("test").
		SetLastName("testd").
		SetPhone("332 111").
		SetPassword("32123556").
		SetConfirmPassword("32123556").
		SetEmail("test@mail").
		SetAvatar("safdsgsf").
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		t.Error(err)
	}

	userUpdate := dto.User{
		Password:        "ufo361@",
		ConfirmPassword: "ufo361@",
		UserName:        "asdg",
		Phone:           "333 333 333",
		Email:           "test2@mail",
		UpdatedAt:       time.Now(),
	}

	t.Run("UpdateUser", func(t *testing.T) {
		if err := repo.UpdateUser(context.Background(), 1, &userUpdate); err != nil {
			t.Error(err)
		}
	})
}

func TestUser_GetByUserName(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewUserRepository(client)

	_, err := repo.DbClient.User.Create().
		SetUserID(1).
		SetUserName("testadmin").
		SetFirstName("test").
		SetLastName("testd").
		SetPhone("332 111").
		SetEmail("test@mail").
		SetPassword("32123556").
		SetConfirmPassword("32123556").
		SetAvatar("safdsgsf").
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		t.Error(err)
	}

	t.Run("GetByUserName", func(t *testing.T) {
		_, err := repo.GetByUserName(context.Background(), "testadmin")
		if err != nil {
			t.Error(err)
		}
	})

}

func TestUser_GetByUserEmail(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := NewUserRepository(client)

	_, err := repo.DbClient.User.Create().
		SetUserID(1).
		SetUserName("testadmin").
		SetFirstName("test").
		SetLastName("testd").
		SetPhone("332 111").
		SetEmail("test@mail").
		SetPassword("32123556").
		SetConfirmPassword("32123556").
		SetAvatar("safdsgsf").
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		t.Error(err)
	}

	t.Run("GetByUserEmail", func(t *testing.T) {
		_, err := repo.GetByUserEmail(context.Background(), "test@mail")
		if err != nil {
			t.Error(err)
		}
	})
}
