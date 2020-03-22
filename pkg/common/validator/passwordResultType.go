//go:generate avaEnum -f=$GOFILE --marshal --lower --noprefix

package validator

// PasswordResultType x ENUM(
// PasswordResultOK // PasswordResultOK Means the checking ran alright
// PasswordResultDivergent // PasswordResultDivergent Password is different from confirmation
// PasswordResultTooShort // PasswordResultTooShort Password is too short
// PasswordResultTooSimple // PasswordResultTooSimple Given string doesn't satisfy complexity rules
// )
type PasswordResultType uint8
