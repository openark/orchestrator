package sessions

import (
	"github.com/gorilla/sessions"
)

// CookieStore is an interface that represents a Cookie based storage
// for Sessions.
type CookieStore interface {
	// Store is an embedded interface so that CookieStore can be used
	// as a session store.
	Store
	// Options sets the default options for each session stored in this
	// CookieStore.
	Options(Options)
}

// NewCookieStore returns a new CookieStore.
//
// Keys are defined in pairs to allow key rotation, but the common case is to set a single
// authentication key and optionally an encryption key.
//
// The first key in a pair is used for authentication and the second for encryption. The
// encryption key can be set to nil or omitted in the last pair, but the authentication key
// is required in all pairs.
//
// It is recommended to use an authentication key with 32 or 64 bytes. The encryption key,
// if set, must be either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256 modes.
func NewCookieStore(keyPairs ...[]byte) CookieStore {
	return &cookieStore{sessions.NewCookieStore(keyPairs...)}
}

type cookieStore struct {
	*sessions.CookieStore
}

func (c *cookieStore) Options(options Options) {
	c.CookieStore.Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}
