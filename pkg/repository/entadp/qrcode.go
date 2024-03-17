package entadp

import (
	"cafe-management/ent"
	"cafe-management/pkg/repository/dto"
	"context"
	"fmt"
)

var _ QRCodeRepository = (*QRCode)(nil)

func NewQRCodeRepository(dbClient *ent.Client) *QRCode {
	return &QRCode{
		dbClient: dbClient,
	}
}

type QRCode struct {
	dbClient *ent.Client
}

func (r *QRCode) CreateQRCode(ctx context.Context, url string, image []byte) error {
	_, err := r.dbClient.QrCode.Create().
		SetURL(url).
		SetImage(string(image)).
		Save(ctx)
	return err
}

func (q *QRCode) DeleteQRCode(ctx context.Context, id int) error {
	return q.dbClient.QrCode.DeleteOneID(id).Exec(ctx)
}

func (q *QRCode) ListQRCode(ctx context.Context) ([]*ent.QrCode, error) {
	return q.dbClient.QrCode.Query().All(ctx)
}

func (q *QRCode) UpdateQRCode(ctx context.Context, id int, c *dto.QRCode) error {
	_, err := q.dbClient.QrCode.UpdateOneID(id).SetURL(c.Url).SetImage(c.Image).Save(ctx)
	if err != nil {
		return fmt.Errorf("qrcode / update qrcode :%w", err)
	}

	return nil
}
