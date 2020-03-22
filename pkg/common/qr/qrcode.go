package qr

import (
	"github.com/skip2/go-qrcode"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	errorQRCodeAVA "github.com/ver13/ava/pkg/common/qr/error"
)

type QRCode struct {
}

func NewQRCode() QRCodeI {
	return &QRCode{}
}

// Deserializer a QR code and return a raw PNG image.
func (qr *QRCode) Generate(content string) ([]byte, *errorAVA.Error) {
	var png []byte

	png, errEncode := qrcode.Encode(content, qrcode.High, 256)
	if errEncode != nil {
		return nil, errorQRCodeAVA.EncodeQRError(errEncode, content)
	}

	return png, nil
}
