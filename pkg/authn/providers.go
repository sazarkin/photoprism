package authn

import (
	"github.com/photoprism/photoprism/pkg/clean"
	"github.com/photoprism/photoprism/pkg/list"
	"github.com/photoprism/photoprism/pkg/txt"
)

// ProviderType represents an authentication provider type.
type ProviderType string

// Standard authentication provider types.
const (
	ProviderUndefined   ProviderType = ""
	ProviderDefault     ProviderType = "default"
	ProviderClient      ProviderType = "client"
	ProviderApplication ProviderType = "application"
	ProviderAccessToken ProviderType = "access_token"
	ProviderLocal       ProviderType = "local"
	ProviderLDAP        ProviderType = "ldap"
	ProviderLink        ProviderType = "link"
	ProviderNone        ProviderType = "none"
)

// RemoteProviders contains remote auth providers.
var RemoteProviders = list.List{
	string(ProviderLDAP),
}

// LocalProviders contains local auth providers.
var LocalProviders = list.List{
	string(ProviderLocal),
}

// Method2FAProviders contains auth providers that support Method2FA.
var Method2FAProviders = list.List{
	string(ProviderDefault),
	string(ProviderLocal),
	string(ProviderLDAP),
}

// ClientProviders contains all client auth providers.
var ClientProviders = list.List{
	string(ProviderClient),
	string(ProviderApplication),
	string(ProviderAccessToken),
}

// Provider casts a string to a normalized provider type.
func Provider(s string) ProviderType {
	s = clean.TypeLowerUnderscore(s)
	switch s {
	case "", "_", "-", "null", "nil", "0", "false":
		return ProviderDefault
	case "token", "url":
		return ProviderLink
	case "pass", "passwd", "password":
		return ProviderLocal
	case "app", "application":
		return ProviderApplication
	case "ldap", "ad", "ldap/ad", "ldap\\ad":
		return ProviderLDAP
	case "client", "client_credentials", "oauth2":
		return ProviderClient
	default:
		return ProviderType(s)
	}
}

// Pretty returns the provider identifier in an easy-to-read format.
func (t ProviderType) Pretty() string {
	switch t {
	case ProviderLDAP:
		return "LDAP/AD"
	case ProviderClient:
		return "Client"
	case ProviderAccessToken:
		return "Access Token"
	default:
		return txt.UpperFirst(t.String())
	}
}

// String returns the provider identifier as a string.
func (t ProviderType) String() string {
	switch t {
	case "":
		return string(ProviderDefault)
	case "token":
		return string(ProviderLink)
	case "password":
		return string(ProviderLocal)
	case "client", "client credentials", "client_credentials", "oauth2":
		return string(ProviderClient)
	default:
		return string(t)
	}
}

// Equal checks if the type matches the specified string.
func (t ProviderType) Equal(s string) bool {
	return t == Provider(s)
}

// NotEqual checks if the type does not match the specified string.
func (t ProviderType) NotEqual(s string) bool {
	return !t.Equal(s)
}

// Is compares the provider with another type.
func (t ProviderType) Is(providerType ProviderType) bool {
	return t == providerType
}

// IsNot checks if the provider is not the specified type.
func (t ProviderType) IsNot(providerType ProviderType) bool {
	return t != providerType
}

// IsUndefined checks if the provider is undefined.
func (t ProviderType) IsUndefined() bool {
	return t == ""
}

// IsRemote checks if the provider is external.
func (t ProviderType) IsRemote() bool {
	return list.Contains(RemoteProviders, string(t))
}

// IsLocal checks if local authentication is possible.
func (t ProviderType) IsLocal() bool {
	return list.Contains(LocalProviders, string(t))
}

// Supports2FA checks if the provider supports two-factor authentication with a passcode.
func (t ProviderType) Supports2FA() bool {
	return list.Contains(Method2FAProviders, string(t))
}

// IsClient checks if the authentication is provided for a client.
func (t ProviderType) IsClient() bool {
	return list.Contains(ClientProviders, string(t))
}

// IsApplication checks if the authentication is provided for an application.
func (t ProviderType) IsApplication() bool {
	return t == ProviderApplication
}

// IsDefault checks if this is the default provider.
func (t ProviderType) IsDefault() bool {
	return t.String() == ProviderDefault.String()
}
