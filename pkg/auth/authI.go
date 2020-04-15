package auth

import (
	errorAVA "github.com/ver13/ava/pkg/common/error"
)

// Auth providers authentication and authorization
type AuthI interface {
	// Init the auth
	Init(opts ...Option)
	// Options set for auth
	Options() Options
	// Generate a new account
	Generate(id string, opts ...GenerateOption) (*Account, *errorAVA.Error)
	// Grant access to a resource
	Grant(role string, res *Resource) *errorAVA.Error
	// Revoke access to a resource
	Revoke(role string, res *Resource) *errorAVA.Error
	// Verify an account has access to a resource
	Verify(acc *Account, res *Resource) *errorAVA.Error
	// Inspect a token
	Inspect(token string) (*Account, *errorAVA.Error)
	// Token generated using an account ID and secret
	Token(id, secret string, opts ...TokenOption) (*Token, *errorAVA.Error)
	// String returns the name of the implementation
	String() string
}
