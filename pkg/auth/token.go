package auth

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/metadata"
)

// Token can be short or long lived
type Token struct {
	// The token itself
	Token string `json:"token"`
	// Type of token, e.g. JWT
	Type string `json:"type"`
	// Time of token creation
	Created time.Time `json:"created"`
	// Time of token expiry
	Expiry time.Time `json:"expiry"`
	// Subject of the token, e.g. the account ID
	Subject string `json:"subject"`
	// Roles granted to the token
	Roles []string `json:"roles"`
	// Metadata embedded in the token
	Metadata map[string]string `json:"metadata"`
	// Namespace the token belongs to
	Namespace string `json:"namespace"`
}

// BearerScheme used for Authorization header
const BearerScheme = "Bearer "

// ContextWithToken sets the auth token in the context
func ContextWithToken(ctx context.Context, token string) context.Context {
	return metadata.Set(ctx, "Authorization", fmt.Sprintf("%v%v", BearerScheme, token))
}
