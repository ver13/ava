//go:generate avaEnum -f=$GOFILE --marshal --lower --noprefix

package validator

// PasswordComplexityRulesType x ENUM(
// PasswordComplexityLowest = 1 // PasswordComplexityLowest There's no rules besides the minimum length
// PasswordComplexityRequireLetter = 2 // PasswordComplexityRequireLetter At least one letter is required in order to aprove password
// PasswordComplexityRequireUpperCase = 4 // PasswordComplexityRequireUpperCase At least one uppercase letter is required in order to aprove password. Only works if PasswordComplexityRequireLetter is included/activated
// PasswordComplexityRequireNumber = 8 // PasswordComplexityRequireNumber At least one number is required in order to aprove password
// PasswordComplexityRequireSpace = 16 // PasswordComplexityRequireSpace The password must contain at least one space
// PasswordComplexityRequireSymbol = 32 // PasswordComplexityRequireSymbol User have to include at least one special character, like # or -
// )
type PasswordComplexityRulesType uint8
