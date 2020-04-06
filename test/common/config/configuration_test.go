package config_test

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"

	. "github.com/ver13/ava/pkg/common/config"
	modelConfigAVA "github.com/ver13/ava/pkg/common/config/model"
	fileSourceConfigAVA "github.com/ver13/ava/pkg/common/config/source/file"
	loggerAVA "github.com/ver13/ava/pkg/common/logger"
)

type ConfigurationSuite struct {
	suite.Suite
}

func TestConfigurationInit(t *testing.T) {
	suite.Run(t, new(ConfigurationSuite))
}

func (r *ConfigurationSuite) BeforeTest() {
	r.T().Log("BeforeTest")
}

func (r *ConfigurationSuite) AfterTest() {
	r.T().Log("AfterTest")
}

func (r *ConfigurationSuite) SetupSuite() {
	r.T().Log("SetupSuite")
}

func (r *ConfigurationSuite) SetupTest() {
	r.T().Log("SetupTest")
}

func (r *ConfigurationSuite) TearDownSuite() {
	r.T().Log("TearDownSuite")
}

func (r *ConfigurationSuite) TearDownTest() {
	r.T().Log("TearDownTest")
}

/*
func (r *ConfigurationSuite) TestConfiguration_Parser() {
	Convey("Given a Configuration ", r.T(), func() {
		Convey("When it's empty ", func() {
			configurationViper := &file.Configuration{}
			configuration, err := configurationViper.Parser(types.EnvironmentTypeDevelopment, serializerGmf.SerializerTypeJson)

			So(err, ShouldNotBeNil)
			So(configuration, ShouldBeNil)
		})
		Convey("When it has the basic fields only filled ", func() {
			configurationViper := &file.Configuration{
				ProjectName:  "Prueba",
				Author:       "Accenture",
				Copyright:    "Texto",
				Environments: nil,
			}
			configuration, err := configurationViper.Parser(types.EnvironmentTypeDevelopment, serializerGmf.SerializerTypeJson)

			So(err, ShouldNotBeNil)
			So(configuration, ShouldBeNil)
		})
		Convey("When it has the basic fields and version information filled ", func() {
			configurationViper := &file.Configuration{
				ProjectName:  "Prueba",
				Author:       "Accenture",
				Copyright:    "Texto",
				Environments: nil,
			}
			configuration, err := configurationViper.Parser(types.EnvironmentTypeDevelopment, serializerGmf.SerializerTypeJson)

			So(err, ShouldNotBeNil)
			So(configuration, ShouldBeNil)
		})
		Convey("When it has all fields filled ", func() {
			configurationViper := &file.Configuration{
				ProjectName: "Prueba",
				Author:      "Accenture",
				Copyright:   "Texto",
				Environments: []*http.EnvironmentViper{
					{
						Name:   "Test",
						Type:   "development",
						Logger: nil,
						Tls:    nil,
						Api: &http.ApiViper{
							Endpoints:           nil,
							CacheTTL:            60,
							Host:                nil,
							Version:             1,
							OutputEncoding:      "JSON",
							Timeout:             20,
							APITime:             nil,
							CORS:                nil,
							DisableKeepAlives:   false,
							DisableCompression:  false,
							MaxIdleConns:        0,
							MaxIdleConnsPerHost: 0,
							DisableStrictREST:   false,
							Debug:               true,
						},
					},
				},
			}
			configuration, err := configurationViper.Parser(types.EnvironmentTypeDevelopment, serializerGmf.SerializerTypeJson)

			So(err, ShouldBeNil)
			So(configuration, ShouldNotBeNil)
		})
		Convey("When it has environment wrong ", func() {
			configurationViper := &file.Configuration{
				ProjectName: "Prueba",
				Author:      "Accenture",
				Copyright:   "Texto",
				Environments: []*http.EnvironmentViper{
					{
						Name:   "Test",
						Type:   "development",
						Logger: nil,
						Tls:    nil,
						Api: &http.ApiViper{
							Endpoints:           nil,
							CacheTTL:            60,
							Host:                nil,
							Version:             1,
							OutputEncoding:      "JSON",
							Timeout:             20,
							APITime:             nil,
							CORS:                nil,
							DisableKeepAlives:   false,
							DisableCompression:  false,
							MaxIdleConns:        0,
							MaxIdleConnsPerHost: 0,
							DisableStrictREST:   false,
							Debug:               true,
						},
					},
				},
			}
			configuration, err := configurationViper.Parser(types.EnvironmentTypeProduction, serializerGmf.SerializerTypeJson)

			So(err, ShouldNotBeNil)
			So(configuration, ShouldBeNil)
		})
		Convey("When it has environment wrong by api field is empty ", func() {
			configurationViper := &file.Configuration{
				ProjectName: "Prueba",
				Author:      "Accenture",
				Copyright:   "TEXTO",
				Environments: []*http.EnvironmentViper{{
					Name:   "Test",
					Type:   "production",
					Logger: nil,
					Tls:    nil,
					Api:    nil,
				},
				},
			}
			configuration, err := configurationViper.Parser(types.EnvironmentTypeProduction, serializerGmf.SerializerTypeJson)

			So(err, ShouldNotBeNil)
			So(configuration, ShouldBeNil)
		})
		Convey("When it has environment empty ", func() {
			configurationViper := &file.Configuration{
				ProjectName:  "Prueba",
				Author:       "Accenture",
				Copyright:    "Texto",
				Environments: []*http.EnvironmentViper{},
			}
			configuration, err := configurationViper.Parser(types.EnvironmentTypeProduction, serializerGmf.SerializerTypeJson)

			So(err, ShouldNotBeNil)
			So(configuration, ShouldBeNil)
		})
	})
}
*/

func (r *ConfigurationSuite) TestConfiguration_File() {
	Convey("Given a configurationServiceI file ", r.T(), func() {
		Convey("When configurationServiceI file exist ", func() {
			Convey("When there is a unique environment ", func() {
				Convey("When profile type is development ", func() {
					configPath := "/tmp"
					fileName := "error.yaml"
					content := `
ProjectName:  Prueba
Author:       Accenture
Copyright:    Texto
environments:
- name: Desarrollo
  type: development # [development, integration, production]

  logger:
        format: Text # [Text, JSON, ApacheCommonLog, ApacheCombinedLog, ApacheErrorLog, RFC3164Log, CommonLogfileFormat]
        level: debug # [panic, fatal, error, warn, warning, info, debug, trace]
        enable: true
        output:
            console: true
            file: true
        file:
            # Filename is the file to write logs to.
            Filename: projectName.log
            # MaxSize is the maximum size in megabytes of the logger file before it gets rotated.
            MaxSize: 100
            # maxAge is the maximum number of days to retain old logger files based on the timestamp encoded in their filename.
            maxAge: 10
            # MaxBackups is the maximum number of old logger files to retain.
            MaxBackups: 5
            # LocalTime determines if the time used for formatting the timestamps in backup files is the computer's local time.
            LocalTime: true
            # Compress determines if the rotated logger files should be compressed using gzip.
            Compress: true

  crypto:
        enable: true
        #
        public_key: /Users/v.encinas.rojas/workspace/src/github.com/ValentinEncinasRojas/ava/assets/userDB.cert
        # privateKey contains the private key corresponding to the public key in Leaf.
        private_key: /Users/v.encinas.rojas/workspace/src/github.com/ValentinEncinasRojas/ava/assets/userDB.key
        # minVersion contains the minimum SSL/tls version that is acceptable.
        min_version: TLS13 # [TLS10, TLS11, TLS12, TLS13]
        # maxVersion contains the maximum SSL/tls version that is acceptable.
        max_version: TLS13 # [TLS10, TLS11, TLS12, TLS13]
        # curvePreferences contains the elliptic curves that will be used in an ECDHE handshake, in preference order.
        curve_preferences: # [CurveP256, CurveP384, CurveP521, X25519]
        - CurveP256
        - CurveP384
        prefer_server_cipher_suites: false
        # cipherSuites is a list of supported cipher suites for tls versions up to tls 1.2.
        cipher_suites: # [TLS_RSA_WITH_RC4_128_SHA, TLS_RSA_WITH_3DES_EDE_CBC_SHA, TLS_RSA_WITH_AES_128_CBC_SHA, TLS_RSA_WITH_AES_256_CBC_SHA, TLS_RSA_WITH_AES_128_CBC_SHA256, TLS_RSA_WITH_AES_128_GCM_SHA256, TLS_RSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_ECDSA_WITH_RC4_128_SHA, TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_RSA_WITH_RC4_128_SHA, TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256, TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305, TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, TLS_AES_128_GCM_SHA256, TLS_AES_256_GCM_SHA384, TLS_CHACHA20_POLY1305_SHA256, TLS_FALLBACK_SCSV]
        - TLS_RSA_WITH_RC4_128_SHA

  api:
        endpoints:
        - # url pattern to be registered and exposed to the world.
          url_pattern: ping
          # HTTP method of the endpoint (GET, POST, PUT, etc)
          method: GET # [GET, HEAD, POST, PUT, DELETE, CONNECT, OPTIONS, TRACE, PATCH]
          # Set of definitions of the backends to be linked to this endpoint.
          backend:
          -
          # Number of concurrent calls this endpoint must send to the backends.
          concurrent_calls: 1
          # timeout of this endpoint.
          timeout: 10
          # Duration of the cache header.
          cache_ttl: 10
          # List of query string params to be extracted from the URI.
          querystring_params:
          -
          # headersToPass defines the list of headers to pass to the backends.
          headers_to_pass:
          -
          # OutputEncodingType defines the error strategy to use for the endpoint responses.
          output_encoding: JSON # [JSON, XML]

        # Default TTL for GET.
        cacheTTL: 10
        # Default set of hosts.
        host:
            - 127.0.0.1
        # version code of the configurationServiceI
        version: 1
        # OutputEncodingType defines the default error strategy to use for the endpoint responses.
        output_encoding: JSON # [JSON, XML]
        # Defafult timeout.
        timeout: 60

        Times:
            # ReadTimeout is the maximum duration for reading the entire request, including the body.
            read_timeout: 30
            # WriteTimeout is the maximum duration before timing out writes of the response. It is reset whenever a new request's header is read. Like ReadTimeout, it does not let Handlers make decisions on a per-request basis.
            write_timeout: 30
            # IdleTimeout is the maximum amount of time to wait for the next request when keep-alives are enabled. If IdleTimeout is zero, the value of ReadTimeout is used. If both are zero, ReadHeaderTimeout is used.
            idle_timeout: 30
            # ReadHeaderTimeout is the amount of time allowed to read request headers. The connection's read deadline is reset after reading the headers and the Handler can decide what is considered too slow for the body.
            read_header_timeout: 30
            # IdleConnTimeout is the maximum amount of time an idle (keep-alive) connection will remain idle before closing itself. Zero means no limit.
            idle_connection_timeout: 30
            # ResponseHeaderTimeout, if non-zero, specifies the amount of time to wait for a server's response headers after fully writing the request (including its body, if any). This time does not include the time to read the response body.
            response_header_timeout: 30
            # ExpectContinueTimeout, if non-zero, specifies the amount of time to wait for a server's first response headers after fully writing the request headers if the request has an "Expect: 100-continue" header. Zero means no timeout and causes the body to be sent immediately, without waiting for the server to approve.
            expect_continue_timeout: 30

            #
            dialer:
                # DialerTimeout is the maximum amount of time a dial will wait for a connect to complete. If Deadline is also set, it may fail earlier. The default is no timeout.
                dialer_timeout: 30
                # FallbackDelay specifies the length of time to wait before spawning a fallback connection, when DualStack is enabled. If zero, a default delay of 300ms is use.
                dialer_fallback_delay: 30
                # KeepAlive specifies the keep-alive period for an active network connection. If zero, keep-alives are not enabled.
                dialer_keep_alive: 30

        cors:
            # AllowedOrigins is a list of origins a cross-domain request can be executed from. If the special "*" value is present in the list, all origins will be allowed.
            allow_origins:
            - "*"
            # ExposedHeaders indicates which headers are safe to expose to the api of a cors api specification
            expose_headers:
            - Link
            # maxAge indicates how long (in seconds) the results of a preflight request can be cached
            max_age: 300
            # AllowedMethods is a list of methods the client is allowed to use with cross-domain requests. Default value is simple methods (GET and POST)
            allow_methods: # [GET, HEAD, POST, PUT, DELETE, CONNECT, OPTIONS, TRACE, PATCH]
            - GET
            - HEAD
            - POST
            - PUT
            # AllowedHeaders is list of non simple headers the client is allowed to use with cross-domain requests. If the special "*" value is present in the list, all headers will be allowed.
            allow_headers:
            - Accept
            - Accept-encoding
            - Authorization
            # allowCredentials indicates whether the request can include user credentials like cookies, HTTP authService or client side SSL certificates.
            allow_credentials: false
            debug: true

        # disableKeepAlives, if true, prevents re-use of TCP connections between different HTTP requests.
        disable_keep_alives: true
        # disableCompression, if true, prevents the Transport from requesting compression with an "Accept-encoding: gzip" request header when the Request contains no existing Accept-encoding value.
        disable_compression: true
        # maxIdleConns controls the maximum number of idle (keep-alive) connections across all hosts. Zero means no limit.
        max_idle_connections: 0
        # maxIdleConnsPerHost, if non-zero, controls the maximum idle (keep-alive) connections to keep per-host. If zero, DefaultMaxIdleConnsPerHost is used.
        max_idle_connections_per_host: 0
        # disableStrictREST flags if the REST enforcement is disabled.
        disable_rest: false
        # Run Gmf in debug mode.
        debug: true
`
					tmpfile, err := ioutil.TempFile(configPath, fileName)
					So(err, ShouldBeNil)

					if _, err := tmpfile.Write([]byte(content)); err != nil {
						tmpfile.Close()
						panic(err)
					}

					defer os.Remove(tmpfile.Name()) // clean up

					configuration, status := ReadLocal(tmpfile.Name(), modelConfigAVA.EnvironmentTypeDevelopment)

					So(status, ShouldBeNil)
					So(configuration, ShouldNotBeNil)
					So(configuration.Environment.Type, ShouldNotBeNil)
					So(configuration.Environment.Logger, ShouldNotBeNil)
					So(configuration.Environment.TLS, ShouldNotBeNil)
					So(configuration.Environment.API, ShouldNotBeNil)

					So(configuration.Environment.Logger.GetFormat(), ShouldEqual, loggerAVA.LogFormatterTypeText)
					So(configuration.Environment.Logger.IsEnable(), ShouldBeTrue)
					So(configuration.Environment.Logger.GetLevel(), ShouldEqual, loggerAVA.LogLevelTypeDebug)

					if err := tmpfile.Close(); err != nil {
						panic(err)
					}
				})
			})
		})
	})
}

func (r *ConfigurationSuite) TestConfiguration_Parser() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {
			Convey("When it's empty ", func() {
				configurationViper := &fileSourceConfigAVA.ConfigurationConfig{}
				configuration, err := configurationViper.Parser(modelConfigAVA.EnvironmentTypeDevelopment)

				So(err, ShouldNotBeNil)
				So(configuration, ShouldBeNil)
			})
		})
		Convey("Went it's wrong ", func() {
			Convey("When it has the basic fields only filled ", func() {
				configurationViper := &fileSourceConfigAVA.ConfigurationConfig{
					ProjectName:  "Prueba",
					Author:       "Accenture",
					Copyright:    "Copyright",
					Environments: nil,
				}
				configuration, err := configurationViper.Parser(modelConfigAVA.EnvironmentTypeDevelopment)

				So(err, ShouldNotBeNil)
				So(configuration, ShouldBeNil)
			})
		})
	})
}

func (r *ConfigurationSuite) TestConfiguration_Serializer() {
	Convey("Given a ", r.T(), func() {
		Convey("Went it's OK ", func() {

		})
		Convey("Went it's wrong ", func() {

		})
	})
}
