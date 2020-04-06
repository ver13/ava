package crypto_test

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	. "github.com/ver13/ava/pkg/common/crypto"
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

type passwordSuite struct {
	suite.Suite
}

func TestPasswordInit(t *testing.T) {
	suite.Run(t, new(passwordSuite))
}

func (r *passwordSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *passwordSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *passwordSuite) SetupSuite() {
	r.T().Log("SetupSuite")
}

func (r *passwordSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *passwordSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *passwordSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *passwordSuite) TestPassword_ComparePassword() {
	Convey("Compare password ", r.T(), func() {
		tests := []struct {
			name        string
			p           *Password
			newPassword string
			want        *errorAVA.Error
		}{
			{
				"Equals",
				NewPassword("Password"),
				"Password",
				nil,
			},
			{
				"Not equals",
				NewPassword("Password"),
				"Different",
				nil,
			},
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				if got := tt.p.ComparePassword(tt.newPassword); reflect.DeepEqual(got, tt.want) {
					r.T().Errorf("ComparePassword() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *passwordSuite) TestPassword_Hash() {
	Convey("Hash ", r.T(), func() {
		tests := []struct {
			name string
			p    Password
			want string
			err  errorAVA.Error
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			r.T().Run(tt.name, func(t *testing.T) {
				got, err := tt.p.Hash()
				if got != tt.want {
					t.Errorf("Hash() got = %v, want %v", got, tt.want)
				}
				if !reflect.DeepEqual(err, tt.err) {
					r.T().Errorf("Hash() err = %v, errWwant %v", err, tt.err)
				}
			})
		}
	})
}
