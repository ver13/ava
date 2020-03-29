package http

import (
	errorConfigAVA "github.com/ver13/ava/pkg/common/config/error"
	"github.com/ver13/ava/pkg/common/config/model/http"
	errorAVA "github.com/ver13/ava/pkg/common/error"
	fileAVA "github.com/ver13/ava/pkg/common/file"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
	stringAVA "github.com/ver13/ava/pkg/common/string"
)

// TLSConfig defines the configuration service params for enabling tls (HTTPS & HTTP/2) at the router layer
type TLSConfig struct {
	Enable                   bool     `mapstructure:"enable,omitempty"`
	PublicKey                string   `mapstructure:"public_key"`
	PrivateKey               string   `mapstructure:"private_key"`
	MinVersion               string   `mapstructure:"min_version,omitempty"`
	MaxVersion               string   `mapstructure:"max_version,omitempty"`
	CurvePreferences         []string `mapstructure:"curve_preferences"`
	PreferServerCipherSuites bool     `mapstructure:"prefer_server_cipher_suites"`
	CipherSuites             []string `mapstructure:"cipher_suites"`
}

func (tls *TLSConfig) Parser() (*http.TLS, *errorAVA.Error) {

	if !tls.Enable {
		tls := new(http.TLS)
		tls.Enable = false
		return tls, nil
	}

	if tls.PublicKey == "" {
		return nil, errorConfigAVA.PublicKeyNotExist(nil, "It's empty.")
	} else {
		if err := fileAVA.NewFile().FileExists(tls.PublicKey); err != nil {
			return nil, err
		}
	}

	if tls.PrivateKey == "" {
		return nil, errorConfigAVA.PrivateKeyNotExist(nil, "It's empty.")
	} else {
		if err := fileAVA.NewFile().FileExists(tls.PrivateKey); err != nil {
			return nil, err
		}
	}

	var cipherSuites = make([]http.CipherSuitesType, len(tls.CipherSuites))
	for i := 0; i < len(tls.CipherSuites); i++ {
		cipherSuite, err := http.ParseCipherSuitesType(tls.CipherSuites[i])
		if err != nil {
			return nil, errorConfigAVA.CipherSuiteWrong(err, tls.CipherSuites[i])
		}
		cipherSuites[i] = cipherSuite
	}

	var curvePreferences = make([]http.SupportedCurvesType, len(tls.CurvePreferences))
	for i := 0; i < len(tls.CurvePreferences); i++ {
		curvePreference, err := http.ParseSupportedCurvesType(tls.CurvePreferences[i])
		if err != nil {
			return nil, errorConfigAVA.CurvePreferencesWrong(err, tls.CurvePreferences[i])
		}
		curvePreferences[i] = curvePreference
	}

	maxVersion, err := http.ParseSupportedVersionsType(stringAVA.StringToUpper(tls.MaxVersion))
	if err != nil {
		return nil, errorConfigAVA.TLSVersionWrong(err, tls.MaxVersion)
	}

	minVersion, err := http.ParseSupportedVersionsType(stringAVA.StringToUpper(tls.MinVersion))
	if err != nil {
		return nil, errorConfigAVA.TLSVersionWrong(err, tls.MinVersion)
	}

	return http.NewTLS(tls.Enable, tls.PublicKey, tls.PrivateKey, minVersion, maxVersion, curvePreferences, tls.PreferServerCipherSuites, cipherSuites)
}

func (tls *TLSConfig) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(tls)
}
