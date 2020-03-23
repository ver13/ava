package error

const (
	statusBcryptHash                      = 1
	statusPasswordIsEmpty                 = 2
	statusMismatchedHashAndPassword       = 3
	statusGenerateKeyWrong                = 4
	statusCreateCertificate               = 5
	statusCreatePrivateKey                = 6
	statusX509KeyPair                     = 7
	statusGenerateRandomSerialNumberWrong = 8
)

var statusText = map[int]string{
	statusBcryptHash:                      "Bcrypt hash error.",
	statusPasswordIsEmpty:                 "Password is empty.",
	statusMismatchedHashAndPassword:       "Mismatched hash and password.",
	statusGenerateKeyWrong:                "Generate key wrong.",
	statusCreateCertificate:               "Create certificate wrong.",
	statusCreatePrivateKey:                "Create private key wrong.",
	statusX509KeyPair:                     "X509 key pair wrong.",
	statusGenerateRandomSerialNumberWrong: "Generate random serial number wrong.",
}

// statusTextFunc returns a text for the logger status code.
// It returns the empty string if the code is unknown.
func statusTextFunc(code int) string {
	return statusText[code]
}
