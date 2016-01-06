# auth [![wercker status](https://app.wercker.com/status/8e5237b01b52f169a1274fad9a89617b "wercker status")](https://app.wercker.com/project/bykey/8e5237b01b52f169a1274fad9a89617b)
Martini middleware/handler for http basic authentication.

[API Reference](http://godoc.org/github.com/martini-contrib/auth)

## Simple Usage

Use `auth.Basic` to authenticate against a pre-defined username and password:

~~~ go
import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/auth"
)

func main() {
  m := martini.Classic()
  // authenticate every request
  m.Use(auth.Basic("username", "secretpassword"))
  m.Run()
}
~~~

## Advanced Usage

Using `auth.BasicFunc` lets you authenticate on a per-user level, by checking
the username and password in the callback function:

~~~ go
import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/auth"
)

func main() {
  m := martini.Classic()
  // authenticate every request
  m.Use(auth.BasicFunc(func(username, password string) bool {
    return username == "admin" && password == "guessme"
  }))
  m.Run()
}
~~~

Note that checking usernames and passwords with string comparison might be
susceptible to timing attacks. To avoid that, use `auth.SecureCompare` instead:

~~~ go
  m.Use(auth.BasicFunc(func(username, password string) bool {
    return auth.SecureCompare(username, "admin") && auth.SecureCompare(password, "guessme")
  }))
}
~~~

Upon successful authentication, the username is available to all subsequent
handlers via the `auth.User` type:

~~~ go
  m.Get("/", func(user auth.User) string {
    return "Welcome, " + string(user)
  })
}
~~~

## Authors
* [Jeremy Saenz](http://github.com/codegangsta)
* [Brendon Murphy](http://github.com/bemurphy)
