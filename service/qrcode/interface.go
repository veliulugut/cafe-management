package qrcode

import "context"

type ServiceQrCode interface {
	GenerateQRCode(ctx context.Context, c *CreateQrCodeModel) (string, []byte, error)
	DeleteQRCode(ctx context.Context, id int) error
	UpdateQRCode(ctx context.Context, id int, c *UpdateQrCodeModel) error
	ListQRCode(ctx context.Context) ([]*ListQRCodeModel, error)
}
