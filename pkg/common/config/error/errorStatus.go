package error

const (
	statusInvalidConfig           = 1
	statusEnvironmentWrong        = 2
	statusAPISEmpty               = 3
	statusEnvironmentsIsEmpty     = 4
	statusConfigVersionWrong      = 5
	statusOutputEncodingWrong     = 6
	statusAllowedMethodWrong      = 7
	statusHTTPVerbWrong           = 8
	statusURLWrong                = 10
	statusCurvePreferencesWrong   = 11
	statusTLSVersionWrong         = 12
	statusPublicKeyNotExist       = 13
	statusPrivateKeyNotExist      = 14
	statusCipherSuiteWrong        = 15
	statusAllowedHeaderWrong      = 16
	statusExposedHeaderWrong      = 17
	statusEnvironmentUnknown      = 18
	statusAllowMethodWrong        = 19
	statusAPIIsEmpty              = 20
	statusBrokerServiceUnknown    = 21
	statusConfigFileTypeUnknown   = 22
	statusDiscoveryServiceUnknown = 23
	statusURLIsEmpty              = 24
	statusDialectIsWrong          = 25
	statusDialectTypeUnknown      = 26
	statusHTTPHeaderUnknown       = 27
	statusSSLTypeUnknown          = 28
)

var statusText = map[int]string{
	statusInvalidConfig:           "Invalid configuration.",
	statusEnvironmentWrong:        "Environment wrong.",
	statusAPISEmpty:               "API is empty.",
	statusEnvironmentsIsEmpty:     "Environments is empty.",
	statusConfigVersionWrong:      "Config version wrong.",
	statusOutputEncodingWrong:     "Output encoding wrong.",
	statusAllowedMethodWrong:      "Allowed method wrong.",
	statusHTTPVerbWrong:           "HTTP verb wrong.",
	statusURLIsEmpty:              "URL is empty",
	statusURLWrong:                "URL wrong.",
	statusCurvePreferencesWrong:   "Curve preferences wrong.",
	statusTLSVersionWrong:         "TLS version wrong.",
	statusPublicKeyNotExist:       "Public key not exist.",
	statusPrivateKeyNotExist:      "Private key not exist.",
	statusCipherSuiteWrong:        "TLS cipher suite wrong.",
	statusAllowedHeaderWrong:      "Allowed header wrong.",
	statusExposedHeaderWrong:      "Exposed header wrong.",
	statusEnvironmentUnknown:      "Environment unknown.",
	statusAllowMethodWrong:        "Allow method wrong",
	statusAPIIsEmpty:              "API is empty",
	statusBrokerServiceUnknown:    "Broker service unknown",
	statusConfigFileTypeUnknown:   "Config file type unknown",
	statusDiscoveryServiceUnknown: "Discovery service unknown",
	statusDialectIsWrong:          "Database dialect is wrong.",
	statusDialectTypeUnknown:      "Database dialect type unknown.",
	statusHTTPHeaderUnknown:       "HTTP header unknown.",
	statusSSLTypeUnknown:          "Database SSL type unknown.",
}

// statusTextFunc returns a text for the General status code.
// It returns the empty string if the code is unknown.
func statusTextFunc(code int) string {
	return statusText[code]
}
