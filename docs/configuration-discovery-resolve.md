# Configuration: discovery, name resolving

Let orchestrator know how to resolve hostnames.

Most people will want this:

```json
{
  "HostnameResolveMethod": "default",
  "MySQLHostnameResolveMethod": "@@hostname",
}
```

Your hosts may refer to one another by IP addresses, and/or short names, and/or fqdn, and/or VIPs. `orchestrator` needs to uniquely and consistently identify a host. It does so by resolving the target hostname.

Many find `"MySQLHostnameResolveMethod": "@@hostname"` to be easiest. Your options are:

- `"HostnameResolveMethod": "cname"`: do CNAME resolving on hostname
- `"HostnameResolveMethod": "default"`: no special resolving via net protocols
- `"MySQLHostnameResolveMethod": "@@hostname"`: issue a `select @@hostname`
- `"MySQLHostnameResolveMethod": "@@report_host"`: issue a `select @@report_host`, requires `report_host` to be configured
- `"HostnameResolveMethod": "none"` and `"MySQLHostnameResolveMethod": ""`: do nothing. Never resolve. This may appeal to setups where everything uses IP addresses at all times.
