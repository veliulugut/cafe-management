package qrcode

import (
	"cafe-management/pkg/repository/dto"
	"cafe-management/pkg/repository/entadp"
	"context"
	"fmt"

	"github.com/skip2/go-qrcode"
)

var _ ServiceQrCode = (*QrCode)(nil)

func New(r entadp.RepositoryInterface) *QrCode {
	return &QrCode{
		repo: r,
	}
}

type QrCode struct {
	repo entadp.RepositoryInterface
}

func (q *QrCode) GenerateQRCode(ctx context.Context, c *CreateQrCodeModel) (string, []byte, error) {
	qrCode, err := qrcode.Encode(c.Url, qrcode.Medium, 256)
	if err != nil {
		return "", nil, err
	}

	if err := q.repo.QRCode().CreateQRCode(ctx, c.Url, qrCode); err != nil {
		return "", nil, err
	}

	return c.Url, qrCode, nil
}

func (q *QrCode) DeleteQRCode(ctx context.Context, id int) error {
	return q.repo.QRCode().DeleteQRCode(ctx, id)
}

func (q *QrCode) ListQRCode(ctx context.Context) ([]*ListQRCodeModel, error) {
	qrCodes, err := q.repo.QRCode().ListQRCode(ctx)
	if err != nil {
		return nil, err
	}

	var qrCodeModels []*ListQRCodeModel
	for _, qrCode := range qrCodes {
		qrCodeModels = append(qrCodeModels, &ListQRCodeModel{
			Url:   qrCode.URL,
			Image: qrCode.Image,
		})
	}

	return qrCodeModels, nil
}

func (q *QrCode) UpdateQRCode(ctx context.Context, id int, c *UpdateQrCodeModel) error {
	qrUpdate := dto.QRCode{
		Url:   c.Url,
		Image: c.Image,
	}

	if err := q.repo.QRCode().UpdateQRCode(ctx, id, &qrUpdate); err != nil {
		return fmt.Errorf("qrcode / update qrcode :%w", err)
	}

	return nil
}
