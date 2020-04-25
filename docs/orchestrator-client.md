# orchestrator-client

[orchestrator-client](https://github.com/openark/orchestrator/blob/master/resources/bin/orchestrator-client) is a script that wraps API calls with convenient command line interface.

It can auto-determine the leader of an `orchestrator` setup and in such case forward all requests to the leader.

It closely mimics the `orchestrator` command line interface.

With `orchestrator-client`, you:

- Do not need the `orchestrator` binary installed everywhere; only on hosts where the service runs
- Do not need to deploy `orchestrator` configuration everywhere; only on service hosts
- Do not need to make access to backend DB
- Need to have access to the HTTP api
- Need to set the `ORCHESTRATOR_API` environment variable.
  - Either provide a single endpoint for a proxy, e.g.
  ```shell
  export ORCHESTRATOR_API=https://orchestrator.myservice.com:3000/api
  ```
  - Or provide all `orchestrator` endpoints, and `orchestrator-client` will auto-pick the leader (no need for proxy), e.g.
  ```shell
  export ORCHESTRATOR_API="https://orchestrator.host1:3000/api https://orchestrator.host2:3000/api https://orchestrator.host3:3000/api"
  ```
- You may set up the environment in `/etc/profile.d/orchestrator-client.sh`. If this file exists, it will be inlined by `orchestrator-client`.

### Sample usage

Show currently known clusters (replication topologies):

    orchestrator-client -c clusters

Discover, forget an instance:

    orchestrator-client -c discover -i 127.0.0.1:22987
    orchestrator-client -c forget -i 127.0.0.1:22987

Print an ASCII tree of topology instances. Pass a cluster name via `-i` (see `clusters` command above):

    orchestrator-client -c topology -i 127.0.0.1:22987

> Sample output:
>
>     127.0.0.1:22987
>     + 127.0.0.1:22989
>       + 127.0.0.1:22988
>     + 127.0.0.1:22990

Move the replica around the topology:

    orchestrator-client -c relocate -i 127.0.0.1:22988 -d 127.0.0.1:22987

> Resulting topology:
>
>     127.0.0.1:22987
>     + 127.0.0.1:22989
>     + 127.0.0.1:22988
>     + 127.0.0.1:22990

etc.

### Behind the scenes

The command line interface makes for a nice wrapper to API calls, whose output is then transformed from JSON format to textual format.

As example, the command:

```shell
orchestrator-client -c discover -i 127.0.0.1:22987
```

Translates to (simplified here for convenience):

```shell
curl "$ORCHESTRATOR_API/discover/127.0.0.1/22987" | jq '.Details | .Key'
```

### Meta commands

- `orchestrator-client -c help`: list all available commands
- `orchestrator-client -c which-api`: output the API endpoint `orchestrator-client` would use to invoke a command. This is useful when multiple endpoints are provided via `$ORCHESTRATOR_API`.
- `orchestrator-client -c api -path clusters`: invoke a generic HTTP API call (in this case `clusters`) and return the raw JSON response.
