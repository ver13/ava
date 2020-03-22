package validator

import (
	"sync"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	regexpValidatorGmf "github.com/ver13/ava/pkg/common/validator/regexp"
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
	// return match(str, regexpValidatorGmf.URLRegex)
	// TODO: Falta implementar
	return nil
}
func CheckURL(str string) *errorAVA.Error {
	return GetInstance().CheckURL(str)
}

func (v *validator) CheckNIF(nif string) ([]string, *errorAVA.Error) {
	return match(nif, regexpValidatorGmf.NIFRegex)
}

func (v *validator) CheckName(name string) ([]string, *errorAVA.Error) {
	return match(name, regexpValidatorGmf.NameRegex)
}

func (v *validator) CheckLocation(location string) ([]string, *errorAVA.Error) {
	return match(location, regexpValidatorGmf.LocationRegex)
}

func (v *validator) CheckCountry(country string) ([]string, *errorAVA.Error) {
	return match(country, regexpValidatorGmf.CountryRegex)
}

func (v *validator) CheckContactPerson(contactPerson string) ([]string, *errorAVA.Error) {
	return match(contactPerson, regexpValidatorGmf.ContactRegex)
}

func (v *validator) CheckUsername(username string) ([]string, *errorAVA.Error) {
	return match(username, regexpValidatorGmf.UsernameRegex)
}

func (v *validator) CheckPassword(password string) ([]string, *errorAVA.Error) {
	return match(password, regexpValidatorGmf.PasswordRegex)
}

func (v *validator) CheckInvoiceNumber(invoiceNumber string) ([]string, *errorAVA.Error) {
	return match(invoiceNumber, regexpValidatorGmf.InvoiceNumberRegex)
}

func (v *validator) CheckPaymentTerms(paymentTerms int64) ([]string, *errorAVA.Error) {
	return match(string(paymentTerms), regexpValidatorGmf.PaymentTermsRegex)
}

func (v *validator) CheckPaymentMethod(paymentMethod string) ([]string, *errorAVA.Error) {
	return match(string(paymentMethod), regexpValidatorGmf.PaymentMethod)
}

func (v *validator) CheckInvoiceAmount(invoiceAmount int64) ([]string, *errorAVA.Error) {
	return match(string(invoiceAmount), regexpValidatorGmf.InvoiceAmountRegex)
}

func (v *validator) CheckEmail(email string) ([]string, *errorAVA.Error) {
	return match(email, regexpValidatorGmf.EmailRegex)
}

// Date finds all date strings
func (v *validator) CheckDate(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.DateRegex)
}

// Times finds all time strings
func (v *validator) CheckTime(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.TimeRegex)
}

// Phones finds all phone numbers
func (v *validator) CheckPhones(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.PhoneRegex)
}

// PhonesWithExts finds all phone numbers with ext
func (v *validator) CheckPhonesWithExts(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.PhonesWithExtsRegex)
}

// Links finds all link strings
func (v *validator) CheckLinks(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.LinkRegex)
}

// Emails finds all email strings
func (v *validator) CheckEmails(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.EmailRegex)
}

// IPv4s finds all IPv4 addresses
func (v *validator) CheckIPv4s(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.IPv4Regex)
}

// IPv6s finds all IPv6 addresses
func (v *validator) CheckIPv6s(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.IPv6Regex)
}

// IPs finds all IP addresses (both IPv4 and IPv6)
func (v *validator) CheckIPs(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.IPRegex)
}

// NotKnownPorts finds all not-known port numbers
func (v *validator) CheckNotKnownPorts(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.NotKnownPortRegex)
}

// Prices finds all price strings
func (v *validator) CheckPrices(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.PriceRegex)
}

// HexColors finds all hex color values
func (v *validator) CheckHexColors(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.HexColorRegex)
}

// CreditCards finds all credit card numbers
func (v *validator) CheckCreditCards(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.CreditCardRegex)
}

// BtcAddresses finds all bitcoin addresses
func (v *validator) CheckBtcAddresses(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.BtcAddressRegex)
}

// StreetAddresses finds all street addresses
func (v *validator) CheckStreetAddresses(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.StreetAddressRegex)
}

// ZipCodes finds all zip codes
func (v *validator) CheckZipCodes(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.ZipCodeRegex)
}

// PoBoxes finds all po-box strings
func (v *validator) CheckPoBoxes(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.PoBoxRegex)
}

// SSNs finds all SSN strings
func (v *validator) CheckSSNs(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.SSNRegex)
}

// MD5Hexes finds all MD5 hex strings
func (v *validator) CheckMD5Hexes(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.MD5HexRegex)
}

// SHA1Hexes finds all SHA1 hex strings
func (v *validator) CheckSHA1Hexes(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.SHA1HexRegex)
}

// SHA256Hexes finds all SHA256 hex strings
func (v *validator) CheckSHA256Hexes(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.SHA256HexRegex)
}

// GUIDs finds all GUID strings
func (v *validator) CheckGUIDs(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.GUIDRegex)
}

// ISBN13s finds all ISBN13 strings
func (v *validator) CheckISBN13s(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.ISBN13Regex)
}

// ISBN10s finds all ISBN10 strings
func (v *validator) CheckISBN10s(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.ISBN10Regex)
}

// VISACreditCards finds all VISA credit card numbers
func (v *validator) CheckVISACreditCards(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.VISACreditCardRegex)
}

// CheckMasterCardCreditCards finds all MasterCard credit card numbers
func (v *validator) CheckMasterCardCreditCards(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.MCCreditCardRegex)
}

// MACAddresses finds all MAC addresses
func (v *validator) CheckMACAddresses(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.MACAddressRegex)
}

// IBANs finds all IBAN strings
func (v *validator) CheckIBANs(text string) ([]string, *errorAVA.Error) {
	return match(text, regexpValidatorGmf.IBANRegex)
}

// Check if Address is Valid
func (v *validator) CheckAddressEth(address string) ([]string, *errorAVA.Error) {
	return match(address, regexpValidatorGmf.AddressEthereumRegex)
}
