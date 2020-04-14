package store

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
)

type dbSQLSuite struct {
	suite.Suite
}

func TestDBSQLSuiteInit(t *testing.T) {
	suite.Run(t, new(dbSQLSuite))
}

func (r *dbSQLSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *dbSQLSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *dbSQLSuite) SetupSuite() {
	r.T().Log("SetupSuite")
}

func (r *dbSQLSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *dbSQLSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *dbSQLSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

/*
TestSpec
	Subject: read SQL database configuration file
		Given an initial configuration file
			When the file does not exist
				Must return an error of type FileNotFount
			When the file exists
				When the file is empty
					Must return an error of type ReadFile
				When the file is empty
					Must return an error of type ReadFile
				When the file is filled
					It has all the correct fields
					When the wrong Dialect is OK
						Must return an null error
					When the wrong Dialect field
						Must return a DialectTypeUnknown dialect type
						Must return an error of type InvalidConfig
					When the DBName field is OK
						Must return an null error
					When the DBName field is empty
						Must return an error of type InvalidConfig
					When the Name field is OK
						Must return an null error
					When the Name field is empty
						Must return an error of type InvalidConfig
					When the Host field is OK
						Must return an null error
					When the wrong Host field
						When the Host field is empty
							Must return an error of type InvalidConfig
						When the Host field is a url wrong
							Must return an error of type URLWrong
					When the Password field is OK
						Must return null error
					When the wrong Password field
						When the Password field is empty
							Must return an error of type InvalidConfig
						When the Password field is wrong
							When the Password field has less than 8 characters
								Must return an error of type PasswordWrong
							When the Password field does not have at least 1 alphanumeric character, 1 number and a symbol
								Must return an error of type PasswordWrong
							Must return an error of type InvalidConfig
					When the User field is OK
						Must return null error
					When the User field is wrong
						When the User field is empty
							Must return an error of type InvalidConfig
						When the User field has a invalid value
							Must return an error of type UsernameWrong
					When the Port field is OK
						Must return null error
					When the Port field is wrong
						When the Port field is 0
							Must return an error of type InvalidConfig
						When the Port field has a invalid value
							Must return an error of type CheckNotKnownPorts
*/
func (r *dbSQLSuite) TestDbSQLConfig_Parser() {
	Convey("Subject: Parser SQL database configuration file", r.T(), func() {

		Convey("Given an initial configuration file", func() {

			Convey("When the file is filled", func() {

				Convey("It has all the correct fields", func() {

					Convey("Must return an null error", nil)

				})

				Convey("When the wrong Dialect is OK", func() {

					Convey("Must return an null error", nil)

				})

				Convey("When the wrong Dialect field", func() {

					Convey("Must return a DialectTypeUnknown dialect type", nil)

					Convey("Must return an error of type InvalidConfig", nil)

				})

				Convey("When the DBName field is OK", func() {

					Convey("Must return an null error", nil)

				})

				Convey("When the DBName field is empty", func() {

					Convey("Must return an error of type InvalidConfig", nil)

				})

				Convey("When the Name field is OK", func() {

					Convey("Must return an null error", nil)

				})

				Convey("When the Name field is empty", func() {

					Convey("Must return an error of type InvalidConfig", nil)

				})

				Convey("When the Host field is OK", func() {

					Convey("Must return an null error", nil)

				})

				Convey("When the wrong Host field", func() {

					Convey("When the Host field is empty", func() {

						Convey("Must return an error of type InvalidConfig", nil)

					})

					Convey("When the Host field is a url wrong", func() {

						Convey("Must return an error of type URLWrong", nil)

					})

				})

				Convey("When the Password field is OK", func() {

					Convey("Must return null error", nil)

				})

				Convey("When the wrong Password field", func() {

					Convey("When the Password field is empty", func() {

						Convey("Must return an error of type InvalidConfig", nil)

					})

					Convey("When the Password field is wrong", func() {

						Convey("When the Password field has less than 8 characters", func() {

							Convey("Must return an error of type PasswordWrong", nil)

						})

						Convey("When the Password field does not have at least 1 alphanumeric character, 1 number and a symbol", func() {

							Convey("Must return an error of type PasswordWrong", nil)

						})

						Convey("Must return an error of type InvalidConfig", nil)

					})

				})

				Convey("When the User field is OK", func() {

					Convey("Must return null error", nil)

				})

				Convey("When the User field is wrong", func() {

					Convey("When the User field is empty", func() {

						Convey("Must return an error of type InvalidConfig", nil)

					})

					Convey("When the User field has a invalid value", func() {

						Convey("Must return an error of type UsernameWrong", nil)

					})

				})

				Convey("When the Port field is OK", func() {

					Convey("Must return null error", nil)

				})

				Convey("When the Port field is wrong", func() {

					Convey("When the Port field is 0", func() {

						Convey("Must return an error of type InvalidConfig", nil)

					})

					Convey("When the Port field has a invalid value", func() {

						Convey("Must return an error of type CheckNotKnownPorts", nil)

					})

				})

			})

		})

	})
}

/*
TestSpec
	Subject: read SQL database configuration file
		Given an initial configuration file
			When the file does not exist
				Must return an error of type FileNotFount
			When the file exists
				When the file is empty
					Must return an error of type ReadFile
				When the file is empty
					Must return an error of type ReadFile
				When the file is filled
					It has all the correct fields
					When the wrong Dialect is OK
						Must return an null error
					When the wrong Dialect field
						Must return a DialectTypeUnknown dialect type
						Must return an error of type InvalidConfig
					When the DBName field is OK
						Must return an null error
					When the DBName field is empty
						Must return an error of type InvalidConfig
					When the Name field is OK
						Must return an null error
					When the Name field is empty
						Must return an error of type InvalidConfig
					When the Host field is OK
						Must return an null error
					When the wrong Host field
						When the Host field is empty
							Must return an error of type InvalidConfig
						When the Host field is a url wrong
							Must return an error of type URLWrong
					When the Password field is OK
						Must return null error
					When the wrong Password field
						When the Password field is empty
							Must return an error of type InvalidConfig
						When the Password field is wrong
							When the Password field has less than 8 characters
								Must return an error of type PasswordWrong
							When the Password field does not have at least 1 alphanumeric character, 1 number and a symbol
								Must return an error of type PasswordWrong
							Must return an error of type InvalidConfig
					When the User field is OK
						Must return null error
					When the User field is wrong
						When the User field is empty
							Must return an error of type InvalidConfig
						When the User field has a invalid value
							Must return an error of type UsernameWrong
					When the Port field is OK
						Must return null error
					When the Port field is wrong
						When the Port field is 0
							Must return an error of type InvalidConfig
						When the Port field has a invalid value
							Must return an error of type CheckNotKnownPorts
*/
func (r *dbSQLSuite) TestDbSQLConfig_ReadLocal() {
	Convey("Subject: read SQL database configuration file", r.T(), func() {

		Convey("Given an initial configuration file", func() {

			Convey("When the file does not exist", func() {

				Convey("Must return an error of type FileNotFount", nil)

			})

			Convey("When the file exists", func() {

				Convey("When the file is empty", func() {

					Convey("Must return an error of type ReadFile", nil)

				})

				Convey("When the file is empty", func() {

					Convey("Must return an error of type ReadFile", nil)

				})

				Convey("When the file is filled", func() {

					Convey("It has all the correct fields", nil)

					Convey("When the wrong Dialect is OK", func() {

						Convey("Must return an null error", nil)

					})

					Convey("When the wrong Dialect field", func() {

						Convey("Must return a DialectTypeUnknown dialect type", nil)

						Convey("Must return an error of type InvalidConfig", nil)

					})

					Convey("When the DBName field is OK", func() {

						Convey("Must return an null error", nil)

					})

					Convey("When the DBName field is empty", func() {

						Convey("Must return an error of type InvalidConfig", nil)

					})

					Convey("When the Name field is OK", func() {

						Convey("Must return an null error", nil)

					})

					Convey("When the Name field is empty", func() {

						Convey("Must return an error of type InvalidConfig", nil)

					})

					Convey("When the Host field is OK", func() {

						Convey("Must return an null error", nil)

					})

					Convey("When the wrong Host field", func() {

						Convey("When the Host field is empty", func() {

							Convey("Must return an error of type InvalidConfig", nil)

						})

						Convey("When the Host field is a url wrong", func() {

							Convey("Must return an error of type URLWrong", nil)

						})

					})

					Convey("When the Password field is OK", func() {

						Convey("Must return null error", nil)

					})

					Convey("When the wrong Password field", func() {

						Convey("When the Password field is empty", func() {

							Convey("Must return an error of type InvalidConfig", nil)

						})

						Convey("When the Password field is wrong", func() {

							Convey("When the Password field has less than 8 characters", func() {

								Convey("Must return an error of type PasswordWrong", nil)

							})

							Convey("When the Password field does not have at least 1 alphanumeric character, 1 number and a symbol", func() {

								Convey("Must return an error of type PasswordWrong", nil)

							})

							Convey("Must return an error of type InvalidConfig", nil)

						})

					})

					Convey("When the User field is OK", func() {

						Convey("Must return null error", nil)

					})

					Convey("When the User field is wrong", func() {

						Convey("When the User field is empty", func() {

							Convey("Must return an error of type InvalidConfig", nil)

						})

						Convey("When the User field has a invalid value", func() {

							Convey("Must return an error of type UsernameWrong", nil)

						})

					})

					Convey("When the Port field is OK", func() {

						Convey("Must return null error", nil)

					})

					Convey("When the Port field is wrong", func() {

						Convey("When the Port field is 0", func() {

							Convey("Must return an error of type InvalidConfig", nil)

						})

						Convey("When the Port field has a invalid value", func() {

							Convey("Must return an error of type CheckNotKnownPorts", nil)

						})

					})

				})

			})

		})

	})
}

func (r *dbSQLSuite) TestDbSQLConfig_Serializer() {

}

func (r *dbSQLSuite) TestNewDbSQL() {

}

func (r *dbSQLSuite) TestNewDbSQLDefault() {

}
