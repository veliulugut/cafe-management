package dto

/*
field.String("url").Unique(),

	field.Bytes("image"),
*/
type QRCode struct {
	Url   string `json:"url"`
	Image string `json:"image"`
}
