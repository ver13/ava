package config_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	. "github.com/ver13/ava/pkg/common/config/source/file/stored"
)

type databaseSuite struct {
	suite.Suite
}

func TestDatabaseViperInit(t *testing.T) {
	suite.Run(t, new(databaseSuite))
}

func (r *databaseSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *databaseSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *databaseSuite) SetupSuite() {
	r.T().Log("SetupSuite")
}

func (r *databaseSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *databaseSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *databaseSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

func (r *databaseSuite) TestDatabaseViper_Parser() {
	Convey("Given a database configurationServiceI ", r.T(), func() {
		Convey("When it's empty ", func() {
			databaseViper := &DatabaseSQL{}
			database, err := databaseViper.Parser()

			So(err, ShouldNotBeNil)
			So(database, ShouldBeNil)
		})
		Convey("When it has dialect field ok and DbName field ok ", func() {
			databaseViper := &DatabaseSQL{
				Dialect: "postgres",
				Ssl:     true,
				DbName:  "AVATest",
			}
			database, err := databaseViper.Parser()

			So(err, ShouldNotBeNil)
			So(database, ShouldBeNil)
		})
		Convey("When it has DbName field empty ", func() {
			databaseViper := &DatabaseSQL{
				Dialect:     "postgres",
				Ssl:         true,
				DbName:      "",
				URL:         "jdbc:postgresql://localhost/test",
				User:        "ava",
				Password:    "password",
				Port:        PortPostgreSQLDefault,
				Name:        "AVA",
				AutoMigrate: false,
				Debug:       true,
			}
			database, err := databaseViper.Parser()

			So(err, ShouldNotBeNil)
			So(database, ShouldBeNil)
		})
		Convey("When it has url field empty ", func() {
			databaseViper := &DatabaseSQL{
				Dialect:     "postgres",
				Ssl:         true,
				DbName:      "AVATest",
				URL:         "",
				User:        "ava",
				Password:    "password",
				Port:        PortPostgreSQLDefault,
				Name:        "AVA",
				AutoMigrate: false,
				Debug:       true,
			}
			database, err := databaseViper.Parser()

			So(err, ShouldNotBeNil)
			So(database, ShouldBeNil)
		})
		Convey("When it has User field empty ", func() {
			databaseViper := &DatabaseSQL{
				Dialect:     "postgres",
				Ssl:         true,
				DbName:      "AVATest",
				URL:         "jdbc:postgresql://localhost/test",
				User:        "",
				Password:    "password",
				Port:        PortPostgreSQLDefault,
				Name:        "AVA",
				AutoMigrate: false,
				Debug:       true,
			}
			database, err := databaseViper.Parser()

			So(err, ShouldNotBeNil)
			So(database, ShouldBeNil)
		})
		Convey("When it has Password field empty ", func() {
			databaseViper := &DatabaseSQL{
				Dialect:     "postgres",
				Ssl:         true,
				DbName:      "AVATest",
				URL:         "jdbc:postgresql://localhost/test",
				User:        "ava",
				Password:    "",
				Port:        PortPostgreSQLDefault,
				Name:        "AVA",
				AutoMigrate: false,
				Debug:       true,
			}
			database, err := databaseViper.Parser()

			So(err, ShouldNotBeNil)
			So(database, ShouldBeNil)
		})
		Convey("When it has port field empty ", func() {
			databaseViper := &DatabaseSQL{
				Dialect:     "postgres",
				Ssl:         true,
				DbName:      "AVATest",
				URL:         "jdbc:postgresql://localhost/test",
				User:        "ava",
				Password:    "password",
				Port:        0,
				Name:        "AVA",
				AutoMigrate: false,
				Debug:       true,
			}
			database, err := databaseViper.Parser()

			So(err, ShouldNotBeNil)
			So(database, ShouldBeNil)
		})
		Convey("When it has name field empty ", func() {
			databaseViper := &DatabaseSQL{
				Dialect:     "postgres",
				Ssl:         true,
				DbName:      "AVATest",
				URL:         "jdbc:postgresql://localhost/test",
				User:        "ava",
				Password:    "password",
				Port:        PortPostgreSQLDefault,
				Name:        "",
				AutoMigrate: false,
				Debug:       true,
			}
			database, err := databaseViper.Parser()

			So(err, ShouldNotBeNil)
			So(database, ShouldBeNil)
		})
		Convey("When it has AutoMigrate field is true ", func() {
			databaseViper := &DatabaseSQL{
				Dialect:     "postgres",
				Ssl:         true,
				DbName:      "AVATest",
				URL:         "jdbc:postgresql://localhost/test",
				User:        "ava",
				Password:    "password",
				Port:        PortPostgreSQLDefault,
				Name:        "AVA",
				AutoMigrate: true,
				Debug:       true,
			}
			database, err := databaseViper.Parser()

			So(err, ShouldBeNil)
			So(database, ShouldNotBeNil)
		})
		Convey("When it has debug field is true ", func() {
			databaseViper := &DatabaseSQL{
				Dialect:     "postgres",
				Ssl:         true,
				DbName:      "AVATest",
				URL:         "jdbc:postgresql://localhost/test",
				User:        "ava",
				Password:    "password",
				Port:        PortPostgreSQLDefault,
				Name:        "AVA",
				AutoMigrate: false,
				Debug:       true,
			}
			database, err := databaseViper.Parser()

			So(err, ShouldBeNil)
			So(database, ShouldNotBeNil)
		})
		Convey("When it has all fields ok ", func() {
			databaseViper := &DatabaseSQL{
				Dialect:     "postgres",
				Ssl:         true,
				DbName:      "AVATest",
				URL:         "jdbc:postgresql://localhost/test",
				User:        "ava",
				Password:    "password",
				Port:        PortPostgreSQLDefault,
				Name:        "AVA",
				AutoMigrate: false,
				Debug:       true,
			}
			database, err := databaseViper.Parser()

			So(err, ShouldBeNil)
			So(database, ShouldNotBeNil)
		})
	})
}
