package entadp_test

import (
	"cafe-management/ent"
	"cafe-management/ent/enttest"
	"cafe-management/pkg/repository/dto"
	"cafe-management/pkg/repository/entadp"
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func Test_CreateQrCode(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := entadp.NewQRCodeRepository(client)

	qrCode := &dto.QRCode{
		Url:   "www.google.com",
		Image: string([]byte("test")),
	}

	t.Run("CreateQrCode", func(t *testing.T) {
		err := repo.CreateQRCode(context.Background(), qrCode.Url, []byte(qrCode.Image))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

func Test_ListQrCode(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := entadp.NewQRCodeRepository(client)

	qrCode := []dto.QRCode{
		{
			Url:   "www.google.com",
			Image: "test",
		},
		{
			Url:   "www.instagram.com",
			Image: "test1",
		},
	}

	t.Run("CreateQrCode", func(t *testing.T) {
		for _, v := range qrCode {
			err := repo.CreateQRCode(context.Background(), &v)
			if err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("ListQrCode", func(t *testing.T) {
		_, err := repo.ListQRCode(context.Background())
		if err != nil {
			t.Error(err)
		}
	})
}

func Test_UpdateQrCode(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	qrCode := dto.QRCode{
		Url:   "www.google.com",
		Image: "test",
	}

	t.Run("CreateQrCode", func(t *testing.T) {
		repo := entadp.NewQRCodeRepository(client)
		err := repo.CreateQRCode(context.Background(), &qrCode)
		if err != nil {
			t.Error(err)
		}
	})

	updateQrCode := dto.QRCode{
		Url:   "www.instagram.com",
		Image: "test1",
	}

	t.Run("UpdateQrCode", func(t *testing.T) {
		repo := entadp.NewQRCodeRepository(client)
		err := repo.UpdateQRCode(context.Background(), 1, &updateQrCode)
		if err != nil {
			t.Error(err)
		}
	})
}

func Test_DeleteQrCode(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	repo := entadp.NewQRCodeRepository(client)

	qrCode := dto.QRCode{
		Url:   "www.google.com",
		Image: "test",
	}

	t.Run("CreateQrCode", func(t *testing.T) {
		err := repo.CreateQRCode(context.Background(), &qrCode)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("DeleteQrCode", func(t *testing.T) {
		err := repo.DeleteQRCode(context.Background(), 1)
		if err != nil {
			t.Error(err)
		}
	})
}
