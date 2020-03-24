package validator

import (
	"sync"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	regexpValidatorAVA "github.com/ver13/ava/pkg/common/validator/regexp"
)

type validator struct {
}

var (
	l *validator

	once sync.Once
)

func init() {
	once.Do(func() {
		l = &validator{}
	})
}

func GetInstance() ValidatorI {
	return l
}

func (v *validator) CheckURL(str string) *errorAVA.Error {
	// return match(str, regexpValidatorAVA.URLRegex)
	// TODO: Falta implementar
	return nil
}
func CheckURL(str string) *errorAVA.Error {
	return GetInstance().CheckURL(str)
}

func (v *validator) CheckNIF(nif string) ([]string, *errorAVA.Error) {
	return match(nif, regexpValidatorAVA.NIFRegex)
}

func (v *validator) CheckName(name string) ([]string, *errorAVA.Error) {
	return match(name, regexpValidatorAVA.NameRegex)
}

func (v *validator) CheckLocation(location string) ([]string, *errorAVA.Error) {
	return match(location, regexpValidatorAVA.LocationRegex)
}

func (v *validator) CheckCountry(country string) ([]string, *errorAVA.Error) {
	return match(country, regexpValidatorAVA.CountryRegex)
}

func (v *validator) CheckContactPerson(contactPerson string) ([]string, *errorAVA.Error) {
	return match(contactPerson, regexpValidatorAVA.ContactRegex)
}

func (v *validator) CheckUsername(username string) ([]string, *errorAVA.Error) {
	return match(username, regexpValidatorAVA.UsernameRegex)
}

func (v *validator) CheckPassword(password string) ([]string, *errorAVA.Error) {
	return match(password, regexpValidatorAVA.PasswordRegex)
}

func (v *validator) CheckInvoiceNumber(invoiceNumber string) ([]string, *errorAVA.Error) {
	return match(invoiceNumber, regexpValidatorAVA.InvoiceNumberRegex)
}

func (v *validator) CheckPaymentTerms(paymentTerms int64) ([]string, *errorAVA.Error) {
	return match(string(paymentTerms), regexpValidatorAVA.PaymentTermsRegex)
}

func (v *validator) CheckPaymentMethod(paymentMethod string) ([]string, *errorAVA.Error) {
	return match(string(paymentMethod), regexpValidatorAVA.PaymentMethod)
}

func (v *validator) CheckInvoiceAmount(invoiceAmount int64) ([]string, *errorAVA.Error) {
	return match(string(invoiceAmount), regexpValidatorAVA.InvoiceAmountRegex)
}

func (v *validator) CheckEmail(email string) ([]string, *errorAVA.Error) {
	return match(email, regexpValidatorAVA.EmailRegex)
}

// Date finds all date strings
func (v *validator) CheckDate(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.DateRegex)
}

// Times finds all time strings
func (v *validator) CheckTime(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.TimeRegex)
}

// Phones finds all phone numbers
func (v *validator) CheckPhones(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.PhoneRegex)
}

// PhonesWithExts finds all phone numbers with ext
func (v *validator) CheckPhonesWithExts(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.PhonesWithExtsRegex)
}

// Links finds all link strings
func (v *validator) CheckLinks(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.LinkRegex)
}

// Emails finds all email strings
func (v *validator) CheckEmails(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.EmailRegex)
}

// IPv4s finds all IPv4 addresses
func (v *validator) CheckIPv4s(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.IPv4Regex)
}

// IPv6s finds all IPv6 addresses
func (v *validator) CheckIPv6s(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.IPv6Regex)
}

// IPs finds all IP addresses (both IPv4 and IPv6)
func (v *validator) CheckIPs(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.IPRegex)
}

// NotKnownPorts finds all not-known port numbers
func (v *validator) CheckNotKnownPorts(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.NotKnownPortRegex)
}

// Prices finds all price strings
func (v *validator) CheckPrices(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.PriceRegex)
}

// HexColors finds all hex color values
func (v *validator) CheckHexColors(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.HexColorRegex)
}

// CreditCards finds all credit card numbers
func (v *validator) CheckCreditCards(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.CreditCardRegex)
}

// BtcAddresses finds all bitcoin addresses
func (v *validator) CheckBtcAddresses(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.BtcAddressRegex)
}

// StreetAddresses finds all street addresses
func (v *validator) CheckStreetAddresses(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.StreetAddressRegex)
}

// ZipCodes finds all zip codes
func (v *validator) CheckZipCodes(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.ZipCodeRegex)
}

// PoBoxes finds all po-box strings
func (v *validator) CheckPoBoxes(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.PoBoxRegex)
}

// SSNs finds all SSN strings
func (v *validator) CheckSSNs(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.SSNRegex)
}

// MD5Hexes finds all MD5 hex strings
func (v *validator) CheckMD5Hexes(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.MD5HexRegex)
}

// SHA1Hexes finds all SHA1 hex strings
func (v *validator) CheckSHA1Hexes(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.SHA1HexRegex)
}

// SHA256Hexes finds all SHA256 hex strings
func (v *validator) CheckSHA256Hexes(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.SHA256HexRegex)
}

// GUIDs finds all GUID strings
func (v *validator) CheckGUIDs(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.GUIDRegex)
}

// ISBN13s finds all ISBN13 strings
func (v *validator) CheckISBN13s(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.ISBN13Regex)
}

// ISBN10s finds all ISBN10 strings
func (v *validator) CheckISBN10s(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.ISBN10Regex)
}

// VISACreditCards finds all VISA credit card numbers
func (v *validator) CheckVISACreditCards(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.VISACreditCardRegex)
}

// CheckMasterCardCreditCards finds all MasterCard credit card numbers
func (v *validator) CheckMasterCardCreditCards(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.MCCreditCardRegex)
}

// MACAddresses finds all MAC addresses
func (v *validator) CheckMACAddresses(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.MACAddressRegex)
}

// IBANs finds all IBAN strings
func (v *validator) CheckIBANs(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorAVA.IBANRegex)
}

// Check if Address is Valid
func (v *validator) CheckAddressEth(address string) ([]string, *errorAVA.Error) {
	return match(address, regexpValidatorAVA.AddressEthereumRegex)
}
