// Package sessions contains middleware for easy session management in Martini.
//
//  package main
//
//  import (
//    "github.com/go-martini/martini"
//    "github.com/martini-contrib/sessions"
//  )
//
//  func main() {
// 	  m := martini.Classic()
//
// 	  store := sessions.NewCookieStore([]byte("secret123"))
// 	  m.Use(sessions.Sessions("my_session", store))
//
// 	  m.Get("/", func(session sessions.Session) string {
// 		  session.Set("hello", "world")
// 	  })
//  }
package sessions

import (
	"github.com/go-martini/martini"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

const (
	errorFormat = "[sessions] ERROR! %s\n"
)

// Store is an interface for custom session stores.
type Store interface {
	sessions.Store
}

// Options stores configuration for a session or session store.
//
// Fields are a subset of http.Cookie fields.
type Options struct {
	Path   string
	Domain string
	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'.
	// MaxAge>0 means Max-Age attribute present and given in seconds.
	MaxAge   int
	Secure   bool
	HttpOnly bool
}

// Session stores the values and optional configuration for a session.
type Session interface {
	// Get returns the session value associated to the given key.
	Get(key interface{}) interface{}
	// Set sets the session value associated to the given key.
	Set(key interface{}, val interface{})
	// Delete removes the session value associated to the given key.
	Delete(key interface{})
	// Clear deletes all values in the session.
	Clear()
	// AddFlash adds a flash message to the session.
	// A single variadic argument is accepted, and it is optional: it defines the flash key.
	// If not defined "_flash" is used by default.
	AddFlash(value interface{}, vars ...string)
	// Flashes returns a slice of flash messages from the session.
	// A single variadic argument is accepted, and it is optional: it defines the flash key.
	// If not defined "_flash" is used by default.
	Flashes(vars ...string) []interface{}
	// Options sets confuguration for a session.
	Options(Options)
}

// Sessions is a Middleware that maps a session.Session service into the Martini handler chain.
// Sessions can use a number of storage solutions with the given store.
func Sessions(name string, store Store) martini.Handler {
	return func(res http.ResponseWriter, r *http.Request, c martini.Context, l *log.Logger) {
		// Map to the Session interface
		s := &session{name, r, l, store, nil, false}
		c.MapTo(s, (*Session)(nil))

		// Use before hook to save out the session
		rw := res.(martini.ResponseWriter)
		rw.Before(func(martini.ResponseWriter) {
			if s.Written() {
				check(s.Session().Save(r, res), l)
			}
		})

		// clear the context, we don't need to use
		// gorilla context and we don't want memory leaks
		defer context.Clear(r)

		c.Next()
	}
}

type session struct {
	name    string
	request *http.Request
	logger  *log.Logger
	store   Store
	session *sessions.Session
	written bool
}

func (s *session) Get(key interface{}) interface{} {
	return s.Session().Values[key]
}

func (s *session) Set(key interface{}, val interface{}) {
	s.Session().Values[key] = val
	s.written = true
}

func (s *session) Delete(key interface{}) {
	delete(s.Session().Values, key)
	s.written = true
}

func (s *session) Clear() {
	for key := range s.Session().Values {
		s.Delete(key)
	}
}

func (s *session) AddFlash(value interface{}, vars ...string) {
	s.Session().AddFlash(value, vars...)
	s.written = true
}

func (s *session) Flashes(vars ...string) []interface{} {
	s.written = true
	return s.Session().Flashes(vars...)
}

func (s *session) Options(options Options) {
	s.Session().Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}

func (s *session) Session() *sessions.Session {
	if s.session == nil {
		var err error
		s.session, err = s.store.Get(s.request, s.name)
		check(err, s.logger)
	}

	return s.session
}

func (s *session) Written() bool {
	return s.written
}

func check(err error, l *log.Logger) {
	if err != nil {
		l.Printf(errorFormat, err)
	}
}
