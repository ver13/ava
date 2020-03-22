package validator_test

import (
	"testing"
	
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

)

type validatorSuite struct {
	suite.Suite
}

func TestValidatorInit(t *testing.T) {
	suite.Run(t, new(validatorSuite))
}

func (r *validatorSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *validatorSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *validatorSuite) SetupSuite() {
	r.T().Log("SetupSuite")
}

func (r *validatorSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *validatorSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *validatorSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *validatorSuite) Test_IsURL() {
	Convey("Given a url ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_IsEmail() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckNIF() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckName() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckLocation() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckCountry() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckContactPerson() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckUsername() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckPassword() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckInvoiceNumber() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckPaymentTerms() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckPaymentMethod() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckInvoiceAmount() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckDate() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckTime() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckPhones() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckPhonesWithExts() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckLinks() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckEmails() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckIPv4s() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckIPv6s() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckIPs() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckNotKnownPorts() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckPrices() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckHexColors() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckCreditCards() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckBtcAddresses() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckStreetAddresses() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckZipCodes() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckPoBoxes() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckSSNs() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckMD5Hexes() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckSHA1Hexes() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckSHA256Hexes() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckGUIDs() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckISBN13s() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckISBN10s() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckVISACreditCards() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckMasterCardCreditCards() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckMACAddresses() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckIBANs() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}

func (r *validatorSuite) Test_CheckAddressEth() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}
