# Using orchestrator-client

[orchestrator-client](https://github.com/github/orchestrator/blob/master/resources/bin/orchestrator-client) is a script that wraps API calls with convenient command line interface.

In fact, it closely mimics the `orchestrator` command line interface.

With `orchestrator-client`, you:

- Do not need the `orchestrator` binary installed everywhere; only on hosts where the service runs
- Do not need to deploy `orchestrator` configuration everywhere; only on service hosts
- Do not need to make access to backend DB
- Need to have access to the HTTP api
- Need to set the `ORCHESTRATOR_API` environment variable.
  e.g. `export ORCHESTRATOR_API=https://orchestrator.myservice.com:3000/api`

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

    orchestrator-client -c discover -i 127.0.0.1:22987

Translates to (simplified here for convenience):

    curl "$orchestrator_api/discover/127.0.0.1/22987" | jq '.Details | .Key'
