package http

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type TLS struct {
	Enable                   bool
	PublicKey                string
	PrivateKey               string
	MinVersion               SupportedVersionsType
	MaxVersion               SupportedVersionsType
	CurvePreferences         []SupportedCurvesType
	PreferServerCipherSuites bool
	CipherSuites             []CipherSuitesType
}

func NewTLS(enable bool, publicKey string, privateKey string, minVersion SupportedVersionsType, maxVersion SupportedVersionsType, preferences []SupportedCurvesType, suites bool, cipherSuites []CipherSuitesType) (*TLS, *errorAVA.Error) {
	return &TLS{
		Enable:                   enable,
		PublicKey:                publicKey,
		PrivateKey:               privateKey,
		MinVersion:               minVersion,
		MaxVersion:               maxVersion,
		CurvePreferences:         preferences,
		PreferServerCipherSuites: suites,
		CipherSuites:             cipherSuites,
	}, nil
}

func NewTLSDefault() (*TLS, *errorAVA.Error) {
	return &TLS{
		Enable:                   true,
		PublicKey:                "",
		PrivateKey:               "",
		MinVersion:               SupportedVersionsTypeTLS13,
		MaxVersion:               SupportedVersionsTypeTLS13,
		CurvePreferences:         []SupportedCurvesType{SupportedCurvesTypeCurveP256},
		PreferServerCipherSuites: true,
		CipherSuites:             []CipherSuitesType{CipherSuitesTypeTLSAES128GCMSHA256},
	}, nil
}

func (tls *TLS) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(tls)
}
