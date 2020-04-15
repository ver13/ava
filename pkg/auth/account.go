package auth

import (
	"context"
	"encoding/json"
	"errors"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	"github.com/ver13/ava/pkg/metadata"
)

var (
	// ErrNotFound is returned when a resouce cannot be found
	ErrNotFound = errors.New("not found")
	// ErrEncodingToken is returned when the service encounters an error during encoding
	ErrEncodingToken = errors.New("error encoding the token")
	// ErrInvalidToken is returned when the token provided is not valid
	ErrInvalidToken = errors.New("invalid token provided")
	// ErrInvalidRole is returned when the role provided was invalid
	ErrInvalidRole = errors.New("invalid role")
	// ErrForbidden is returned when a user does not have the necessary roles to access a resource
	ErrForbidden = errors.New("resource forbidden")
)

const (
	// MetadataKey is the key used when storing the account in metadata
	MetadataKey = "auth-account"
	// TokenCookieName is the name of the cookie which stores the auth token
	TokenCookieName = "micro-token"
	// SecretCookieName is the name of the cookie which stores the auth secret
	SecretCookieName = "micro-secret"
)

// Account provided by an auth provider
type Account struct {
	// ID of the account (UUIDV4, email or username)
	ID string `json:"id"`
	// Secret used to renew the account
	Secret string `json:"secret"`
	// Roles associated with the Account
	Roles []string `json:"roles"`
	// Any other associated metadata
	Metadata map[string]string `json:"metadata"`
	// Namespace the account belongs to, default blank
	Namespace string `json:"namespace"`
}

// AccountFromContext gets the account from the context, which
// is set by the auth wrapper at the start of a call. If the account
// is not set, a nil account will be returned. The error is only returned
// when there was a problem retrieving an account
func AccountFromContext(ctx context.Context) (*Account, *errorAVA.Error) {
	str, ok := metadata.Get(ctx, MetadataKey)
	// there was no account set
	if !ok {
		return nil, nil
	}

	var acc *Account
	// metadata is stored as a string, so unmarshal to an account
	if err := json.Unmarshal([]byte(str), &acc); err != nil {
		return nil, err
	}

	return acc, nil
}

// ContextWithAccount sets the account in the context
func ContextWithAccount(ctx context.Context, account *Account) (context.Context, *errorAVA.Error) {
	// metadata is stored as a string, so marshal to bytes
	bytes, err := json.Marshal(account)
	if err != nil {
		return ctx, err
	}

	// generate a new context with the MetadataKey set
	return metadata.Set(ctx, MetadataKey, string(bytes)), nil
}
