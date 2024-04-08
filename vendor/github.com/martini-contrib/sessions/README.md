# sessions [![wercker status](https://app.wercker.com/status/af92c7633124fffea8984e48ee0c418b "wercker status")](https://app.wercker.com/project/bykey/af92c7633124fffea8984e48ee0c418b)
Martini middleware/handler for easy session management.

[API Reference](http://godoc.org/github.com/martini-contrib/sessions)

## Usage

~~~ go
package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/sessions"
)

func main() {
	m := martini.Classic()

	store := sessions.NewCookieStore([]byte("secret123"))
	m.Use(sessions.Sessions("my_session", store))

	m.Get("/set", func(session sessions.Session) string {
		session.Set("hello", "world")
		return "OK"
	})

	m.Get("/get", func(session sessions.Session) string {
		v := session.Get("hello")
		if v == nil {
			return ""
		}
		return v.(string)
	})

  m.Run()
}

~~~

## Authors
* [Jeremy Saenz](http://github.com/codegangsta)
