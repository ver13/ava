package qr_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	. "github.com/ver13/ava/pkg/common/qr"
)

type qrCodeSuite struct {
	suite.Suite

	qr QRCodeI
}

func TestQRInit(t *testing.T) {
	suite.Run(t, new(qrCodeSuite))
}

func (r *qrCodeSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *qrCodeSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *qrCodeSuite) SetupSuite() {
	r.T().Log("SetupSuite")
	r.qr = NewQRCode()
}

func (r *qrCodeSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *qrCodeSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *qrCodeSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *qrCodeSuite) TestQRCode_Generate() {
	Convey("Generate QR ", r.T(), func() {
		result, err := r.qr.Generate("lksjhdaoohd")

		So(err, ShouldBeNil)
		So(result, ShouldResemble, []uint8{137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 1, 0, 0, 0, 1, 0, 1, 3, 0, 0, 0, 102, 188, 58, 37, 0, 0, 0, 6, 80, 76, 84, 69, 255, 255, 255, 0, 0, 0, 85, 194, 211, 126, 0, 0, 1, 26, 73, 68, 65, 84, 120, 218, 236, 152, 191, 109, 232, 32, 24, 196, 63, 139, 130, 146, 17, 60, 10, 163, 217, 163, 49, 10, 35, 184, 164, 64, 220, 19, 127, 204, 11, 36, 182, 210, 37, 10, 119, 149, 141, 126, 213, 39, 238, 56, 16, 138, 162, 168, 223, 46, 131, 42, 103, 113, 137, 74, 91, 251, 13, 4, 102, 160, 126, 30, 78, 118, 0, 105, 195, 199, 213, 165, 0, 13, 64, 14, 103, 189, 232, 2, 28, 192, 69, 224, 13, 48, 32, 240, 13, 32, 111, 185, 40, 4, 158, 129, 230, 77, 91, 130, 233, 197, 188, 171, 3, 61, 204, 189, 9, 239, 105, 191, 54, 208, 149, 129, 215, 158, 240, 199, 1, 29, 84, 146, 123, 58, 121, 78, 167, 29, 172, 71, 160, 19, 40, 171, 110, 255, 186, 38, 17, 232, 82, 213, 123, 237, 15, 167, 245, 98, 8, 124, 2, 52, 226, 6, 156, 205, 155, 57, 204, 115, 81, 32, 48, 3, 38, 232, 152, 79, 61, 192, 27, 196, 114, 0, 194, 155, 75, 8, 76, 192, 255, 180, 7, 130, 170, 54, 157, 247, 228, 10, 192, 112, 165, 45, 97, 14, 39, 195, 36, 9, 76, 125, 210, 139, 168, 148, 51, 234, 169, 112, 174, 14, 220, 55, 148, 114, 254, 165, 2, 76, 222, 36, 48, 61, 14, 136, 108, 185, 82, 185, 125, 124, 77, 34, 48, 77, 242, 190, 161, 184, 167, 45, 183, 56, 208, 189, 105, 16, 235, 36, 199, 135, 220, 69, 128, 59, 204, 75, 179, 76, 173, 38, 93, 4, 102, 128, 162, 40, 234, 103, 244, 47, 0, 0, 255, 255, 105, 71, 175, 101, 95, 4, 142, 253, 0, 0, 0, 0, 73, 69, 78, 68, 174, 66, 96, 130})
	})
}
