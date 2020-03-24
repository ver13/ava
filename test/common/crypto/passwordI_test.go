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

	password PasswordI
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
	r.password = NewPassword("password")
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

func (r *passwordSuite) TestPassword_ComparePassword(t *testing.T) {
	Convey("Compare password ", r.T(), func() {
		type args struct {
			newPassword string
		}
		tests := []struct {
			name string
			p    Password
			args args
			want errorAVA.Error
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := tt.p.ComparePassword(tt.args.newPassword); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ComparePassword() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func (r *passwordSuite) TestPassword_Hash(t *testing.T) {
	Convey("Hash ", r.T(), func() {
		tests := []struct {
			name  string
			p     Password
			want  string
			want1 errorAVA.Error
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, got1 := tt.p.Hash()
				if got != tt.want {
					t.Errorf("Hash() got = %v, want %v", got, tt.want)
				}
				if !reflect.DeepEqual(got1, tt.want1) {
					t.Errorf("Hash() got1 = %v, want %v", got1, tt.want1)
				}
			})
		}
	})
}
