# Security

When operating in HTTP mode (API or Web), access to `orchestrator` may be restricted via either:

*  _basic authentication_

   Add the following to `orchestrator`'s configuration file:

        "AuthenticationMethod": "basic",
        "HTTPAuthUser":         "dba_team",
        "HTTPAuthPassword":     "time_for_dinner"

   With `basic` authentication there's just one single credential, and no roles.

   `Orchestrator`'s configuration file contains credentials to your MySQL servers as well as _basic authentication_
   credentials as specified above. Keep it safe (e.g. `chmod 600`).

*  _basic authentication, extended_

   Add the following to `orchestrator`'s configuration file:

        "AuthenticationMethod": "multi",
        "HTTPAuthUser":         "dba_team",
        "HTTPAuthPassword":     "time_for_dinner"

   The `multi` authentication works like `basic`, but also accepts the user `readonly` with any password. The `readonly` user
   is allowed to view all content but unable to perform write operations through the API (such as stopping a replica,
   repointing replicas, discovering new instances etc.)

*  _Headers authentication_

   Authenticates via headers forwarded by reverse proxy (e.g. Apache2 relaying requests to orchestrator).
   Requires:

        "AuthenticationMethod": "proxy",
        "AuthUserHeader": "X-Forwarded-User",

   You will need to configure your reverse proxy to send the name of authenticated user via HTTP header, and
   use same header name as configured by `AuthUserHeader`.

   For example, an Apache2 setup may look like the following:

        RequestHeader unset X-Forwarded-User
        RewriteEngine On
        RewriteCond %{LA-U:REMOTE_USER} (.+)
        RewriteRule .* - [E=RU:%1,NS]
        RequestHeader set X-Forwarded-User %{RU}e

   The `proxy` authentication allows for *roles*. Some users are *Power users* and the rest are just normal users.
   *Power users* are allowed to make changes to the topologies, whereas normal users are in read-only mode.
   To specify the list of known DBAs, use:

        "PowerAuthUsers": [
            "wallace", "gromit", "shaun"
            ],

Or, regardless, you may turn the entire `orchestrator` process to be read only via:


        "ReadOnly": "true",

You may combine `ReadOnly` with any authentication method you like.
