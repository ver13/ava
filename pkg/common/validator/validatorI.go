//go:generate avaGenerateWrap gen -f=${GOFILE} -t implAVA.tmpl -o ${GOFILE}Impl.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t loggerAVA.tmpl -o ${GOFILE}Logger.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t prometheus.tmpl -o ${GOFILE}Metrics.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t circuitBreakerAVA.tmpl -o ${GOFILE}CircuitBreaker.go
//go:generate avaGenerateWrap gen -f=${GOFILE} -t opentracing.tmpl -o ${GOFILE}Tracing.go
//go:generate avaGenerateTest -f=${GOFILE}

package validator

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

type ValidatorI interface {
	CheckURL(str string) *errorAVA.Error
	CheckEmail(email string) ([]string, *errorAVA.Error)
	CheckNIF(nif string) ([]string, *errorAVA.Error)
	CheckName(name string) ([]string, *errorAVA.Error)
	CheckLocation(location string) ([]string, *errorAVA.Error)
	CheckCountry(country string) ([]string, *errorAVA.Error)
	CheckContactPerson(contactPerson string) ([]string, *errorAVA.Error)
	CheckUsername(username string) ([]string, *errorAVA.Error)
	CheckPassword(password string) ([]string, *errorAVA.Error)
	CheckInvoiceNumber(invoiceNumber string) ([]string, *errorAVA.Error)
	CheckPaymentTerms(paymentTerms int64) ([]string, *errorAVA.Error)
	CheckPaymentMethod(paymentMethod string) ([]string, *errorAVA.Error)
	CheckInvoiceAmount(invoiceAmount int64) ([]string, *errorAVA.Error)
	CheckDate(text string) ([]string, *errorAVA.Error)
	CheckTime(text string) ([]string, *errorAVA.Error)
	CheckPhones(text string) ([]string, *errorAVA.Error)
	CheckPhonesWithExts(text string) ([]string, *errorAVA.Error)
	CheckLinks(text string) ([]string, *errorAVA.Error)
	CheckEmails(text string) ([]string, *errorAVA.Error)
	CheckIPv4s(text string) ([]string, *errorAVA.Error)
	CheckIPv6s(text string) ([]string, *errorAVA.Error)
	CheckIPs(text string) ([]string, *errorAVA.Error)
	CheckNotKnownPorts(text string) ([]string, *errorAVA.Error)
	CheckPrices(text string) ([]string, *errorAVA.Error)
	CheckHexColors(text string) ([]string, *errorAVA.Error)
	CheckCreditCards(text string) ([]string, *errorAVA.Error)
	CheckBtcAddresses(text string) ([]string, *errorAVA.Error)
	CheckStreetAddresses(text string) ([]string, *errorAVA.Error)
	CheckZipCodes(text string) ([]string, *errorAVA.Error)
	CheckPoBoxes(text string) ([]string, *errorAVA.Error)
	CheckSSNs(text string) ([]string, *errorAVA.Error)
	CheckMD5Hexes(text string) ([]string, *errorAVA.Error)
	CheckSHA1Hexes(text string) ([]string, *errorAVA.Error)
	CheckSHA256Hexes(text string) ([]string, *errorAVA.Error)
	CheckGUIDs(text string) ([]string, *errorAVA.Error)
	CheckISBN13s(text string) ([]string, *errorAVA.Error)
	CheckISBN10s(text string) ([]string, *errorAVA.Error)
	CheckVISACreditCards(text string) ([]string, *errorAVA.Error)
	CheckMasterCardCreditCards(text string) ([]string, *errorAVA.Error)
	CheckMACAddresses(text string) ([]string, *errorAVA.Error)
	CheckIBANs(text string) ([]string, *errorAVA.Error)
	CheckAddressEth(address string) ([]string, *errorAVA.Error)
	/*
	   CheckNewPassword(password, passwordConfirmation string, minimumlength uint, flagComplexity PasswordComplexityRulesType) PasswordResultType
	   RuneHasSymbol(ru rune) bool
	*/
}
