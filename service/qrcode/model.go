package qrcode

type CreateQrCodeModel struct {
	Url   string `json:"url"`
	Image string `json:"image"`
}

type UpdateQrCodeModel struct {
	Url   string `json:"url"`
	Image string `json:"image"`
}

type ListQRCodeModel struct {
	Url   string `json:"url"`
	Image string `json:"image"`
}
