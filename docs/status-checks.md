# Status Checks

There is a status endpoint located at `/api/status` that does a healthcheck of the system and reports back
with HTTP status code 200 if everything is ok.  Otherwise it reports back HTTP status code 500.

#### Custom Status Checks
Since there are various standards that companies might use for their status check endpoints, you can
customize this by setting:

```json
{
  "StatusEndpoint": "/my_status"
}
```

Or whatever endpoint you want.  

#### Lightweight Health Check

This status check is a very lightweight check because we assume your load balancer might be hitting it
frequently or some other frequent monitoring.  If you want a richer check that actually makes changes
to the database you can set that with:

```json
{
  "StatusSimpleHealth": false
}
```

#### SSL Verification

Lastly if you run with SSL/TLS we *don't* require the status check to have a valid OU or client cert to be
presented.  If you're using that richer check and would like to have the verification turned on you can set:

```json
{
  "StatusOUVerify": true
}
```
