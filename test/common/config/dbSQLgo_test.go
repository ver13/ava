package config_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	storedModelConfigAVA "github.com/ver13/ava/pkg/common/config/model/stored"
	storedFileConfigAVA "github.com/ver13/ava/pkg/common/config/source/file/stored"
)

type dbSQLSuite struct {
	suite.Suite
}

func TestDatabaseViperInit(t *testing.T) {
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

func (r *dbSQLSuite) TestDatabaseViper_Parser() {
	Convey("Given a database configurationServiceI ", r.T(), func() {
		Convey("When it's empty ", func() {
			databaseViper := &storedFileConfigAVA.DbSQL{}
			database, err := databaseViper.Parser()

			So(err, ShouldNotBeNil)
			So(database, ShouldBeNil)
		})
		Convey("When it has dialect field ok and DbName field ok ", func() {
			databaseViper := &storedFileConfigAVA.DbSQL{
				Dialect: "PostgreSQL",
				SSL:     true,
				DBName:  "AVATest",
			}
			database, err := databaseViper.Parser()

			So(err, ShouldNotBeNil)
			So(database, ShouldBeNil)
		})
		Convey("When it has DbName field empty ", func() {
			databaseViper := &storedFileConfigAVA.DbSQL{
				Dialect:     "PostgreSQL",
				SSL:         true,
				DBName:      "",
				URL:         "jdbc:postgresql://localhost/test",
				User:        "ava",
				Password:    "password",
				Port:        storedModelConfigAVA.PortPostgreSQLDefault,
				Name:        "AVA",
				AutoMigrate: false,
				Debug:       true,
			}
			database, err := databaseViper.Parser()

			So(err, ShouldNotBeNil)
			So(database, ShouldBeNil)
		})
		Convey("When it has url field empty ", func() {
			databaseViper := &storedFileConfigAVA.DbSQL{
				Dialect:     "PostgreSQL",
				SSL:         true,
				DBName:      "AVATest",
				URL:         "",
				User:        "ava",
				Password:    "password",
				Port:        storedModelConfigAVA.PortPostgreSQLDefault,
				Name:        "AVA",
				AutoMigrate: false,
				Debug:       true,
			}
			database, err := databaseViper.Parser()

			So(err, ShouldNotBeNil)
			So(database, ShouldBeNil)
		})
		Convey("When it has User field empty ", func() {
			databaseViper := &storedFileConfigAVA.DbSQL{
				Dialect:     "PostgreSQL",
				SSL:         true,
				DBName:      "AVATest",
				URL:         "jdbc:postgresql://localhost/test",
				User:        "",
				Password:    "password",
				Port:        storedModelConfigAVA.PortPostgreSQLDefault,
				Name:        "AVA",
				AutoMigrate: false,
				Debug:       true,
			}
			database, err := databaseViper.Parser()

			So(err, ShouldNotBeNil)
			So(database, ShouldBeNil)
		})
		Convey("When it has Password field empty ", func() {
			databaseViper := &storedFileConfigAVA.DbSQL{
				Dialect:     "PostgreSQL",
				SSL:         true,
				DBName:      "AVATest",
				URL:         "jdbc:postgresql://localhost/test",
				User:        "ava",
				Password:    "",
				Port:        storedModelConfigAVA.PortPostgreSQLDefault,
				Name:        "AVA",
				AutoMigrate: false,
				Debug:       true,
			}
			database, err := databaseViper.Parser()

			So(err, ShouldNotBeNil)
			So(database, ShouldBeNil)
		})
		Convey("When it has port field empty ", func() {
			databaseViper := &storedFileConfigAVA.DbSQL{
				Dialect:     "PostgreSQL",
				SSL:         true,
				DBName:      "AVATest",
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
			databaseViper := &storedFileConfigAVA.DbSQL{
				Dialect:     "PostgreSQL",
				SSL:         true,
				DBName:      "AVATest",
				URL:         "jdbc:postgresql://localhost/test",
				User:        "ava",
				Password:    "password",
				Port:        storedModelConfigAVA.PortPostgreSQLDefault,
				Name:        "",
				AutoMigrate: false,
				Debug:       true,
			}
			database, err := databaseViper.Parser()

			So(err, ShouldNotBeNil)
			So(database, ShouldBeNil)
		})
		Convey("When it has AutoMigrate field is true ", func() {
			databaseViper := &storedFileConfigAVA.DbSQL{
				Dialect:     "PostgreSQL",
				SSL:         true,
				DBName:      "AVATest",
				URL:         "jdbc:postgresql://localhost/test",
				User:        "ava",
				Password:    "password",
				Port:        storedModelConfigAVA.PortPostgreSQLDefault,
				Name:        "AVA",
				AutoMigrate: true,
				Debug:       true,
			}
			database, err := databaseViper.Parser()

			So(err, ShouldBeNil)
			So(database, ShouldNotBeNil)
		})
		Convey("When it has debug field is true ", func() {
			databaseViper := &storedFileConfigAVA.DbSQL{
				Dialect:     "PostgreSQL",
				SSL:         true,
				DBName:      "AVATest",
				URL:         "jdbc:postgresql://localhost/test",
				User:        "ava",
				Password:    "password",
				Port:        storedModelConfigAVA.PortPostgreSQLDefault,
				Name:        "AVA",
				AutoMigrate: false,
				Debug:       true,
			}
			database, err := databaseViper.Parser()

			So(err, ShouldBeNil)
			So(database, ShouldNotBeNil)
		})
		Convey("When it has all fields ok ", func() {
			databaseViper := &storedFileConfigAVA.DbSQL{
				Dialect:     "PostgreSQL",
				SSL:         true,
				DBName:      "AVATest",
				URL:         "jdbc:postgresql://localhost/test",
				User:        "ava",
				Password:    "password",
				Port:        storedModelConfigAVA.PortPostgreSQLDefault,
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
