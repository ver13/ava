package store

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	storedModelConfigAVA "github.com/ver13/ava/pkg/common/config/model/stored"
	. "github.com/ver13/ava/pkg/common/config/source/file/stored"
	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
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

func (r *dbSQLSuite) TestDbSQLConfig_Parser() {
	Convey("Given a Database configuration file ", r.T(), func() {
		Convey("Parser when it's empty ", func() {
			database := &DbSQLConfig{}
			db, err := database.Parser()

			So(err, ShouldNotBeNil)
			So(db, ShouldBeNil)
		})
		Convey("Parser when it's empty ", func() {
			type fields struct {
				Dialect     string
				Host        string
				Name        string
				User        string
				Password    string
				Port        uint64
				SSL         bool
				DBName      string
				Debug       bool
				AutoMigrate bool
			}
			tests := []struct {
				name   string
				fields fields
				want   *storedModelConfigAVA.DbSQL
				want1  *errorAVA.Error
			}{
				// TODO: Add test cases.
			}
			for _, tt := range tests {
				r.T().Run(tt.name, func(t *testing.T) {
					database := &DbSQLConfig{
						Dialect:     tt.fields.Dialect,
						Host:        tt.fields.Host,
						Name:        tt.fields.Name,
						User:        tt.fields.User,
						Password:    tt.fields.Password,
						Port:        tt.fields.Port,
						SSL:         tt.fields.SSL,
						DBName:      tt.fields.DBName,
						Debug:       tt.fields.Debug,
						AutoMigrate: tt.fields.AutoMigrate,
					}
					got, got1 := database.Parser()
					if !reflect.DeepEqual(got, tt.want) {
						t.Errorf("Parser() got = %v, want %v", got, tt.want)
					}
					if !reflect.DeepEqual(got1, tt.want1) {
						t.Errorf("Parser() got1 = %v, want %v", got1, tt.want1)
					}
				})
			}
		})
	})
}

func TestDbSQLConfig_ReadLocal(t *testing.T) {
	type fields struct {
		Dialect     string
		Host        string
		Name        string
		User        string
		Password    string
		Port        uint64
		SSL         bool
		DBName      string
		Debug       bool
		AutoMigrate bool
	}
	type args struct {
		fileName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *storedModelConfigAVA.DbSQL
		want1  *errorAVA.Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			database := &DbSQLConfig{
				Dialect:     tt.fields.Dialect,
				Host:        tt.fields.Host,
				Name:        tt.fields.Name,
				User:        tt.fields.User,
				Password:    tt.fields.Password,
				Port:        tt.fields.Port,
				SSL:         tt.fields.SSL,
				DBName:      tt.fields.DBName,
				Debug:       tt.fields.Debug,
				AutoMigrate: tt.fields.AutoMigrate,
			}
			got, got1 := database.ReadLocal(tt.args.fileName)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadLocal() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReadLocal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDbSQLConfig_Serializer(t *testing.T) {
	type fields struct {
		Dialect     string
		Host        string
		Name        string
		User        string
		Password    string
		Port        uint64
		SSL         bool
		DBName      string
		Debug       bool
		AutoMigrate bool
	}
	type args struct {
		t serializerAVA.SerializerType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
		want1  *errorAVA.Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			database := &DbSQLConfig{
				Dialect:     tt.fields.Dialect,
				Host:        tt.fields.Host,
				Name:        tt.fields.Name,
				User:        tt.fields.User,
				Password:    tt.fields.Password,
				Port:        tt.fields.Port,
				SSL:         tt.fields.SSL,
				DBName:      tt.fields.DBName,
				Debug:       tt.fields.Debug,
				AutoMigrate: tt.fields.AutoMigrate,
			}
			got, got1 := database.Serializer(tt.args.t)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Serializer() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Serializer() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewDbSQL(t *testing.T) {
	type args struct {
		dialect     string
		url         string
		name        string
		user        string
		password    string
		port        uint64
		ssl         bool
		dbName      string
		debug       bool
		autoMigrate bool
	}
	tests := []struct {
		name  string
		args  args
		want  *storedModelConfigAVA.DbSQL
		want1 *errorAVA.Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := NewDbSQL(tt.args.dialect, tt.args.url, tt.args.name, tt.args.user, tt.args.password, tt.args.port, tt.args.ssl, tt.args.dbName, tt.args.debug, tt.args.autoMigrate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDbSQL() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewDbSQL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewDbSQLDefault(t *testing.T) {
	tests := []struct {
		name  string
		want  *storedModelConfigAVA.DbSQL
		want1 *errorAVA.Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := NewDbSQLDefault()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDbSQLDefault() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewDbSQLDefault() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
