# Executing via command line

Also consult the [Orchestrator first steps](first-steps.md) page.

`orchestrator` supports two ways of running operations from command line:

- Using the `orchestrator` binary (topic of this document)
  - You will deploy `orchestrator` on ops/app boxes, but not run it as a service.
  - You will deploy the configuration file for the `orchestrator` binary to be able to
    connect to the backend DB.
- Using the [orchestrator-client](using-orchestrator-client.md) script.
  - You will only need the `orchestrator-client` script on your ops/app boxes.
  - You will not need any config file nor binaries.
  - You will need to specify the `ORCHESTRATOR_API` environment variable.

The two are (mostly) compatible. This document discusses the first option.

Following is a synopsis of command line samples. For simplicity, we assume `orchestrator` is in your path.
If not, replace `orchestrator` with `/path/to/orchestrator`.

> Samples below use a test `mysqlsandbox` topology, where all instances are on same host `127.0.0.1` and on different ports. `22987` is master,
> and `22988`, `22989`, `22990` are replicas.

Show currently known clusters (replication topologies):

    orchestrator -c clusters

> The above looks for configuration in `/etc/orchestrator.conf.json`, `conf/orchestrator.conf.json`, `orchestrator.conf.json`, in that order.
> Classic is to put configuration in `/etc/orchestrator.conf.json`. Since it contains credentials to your MySQL servers you may wish to limit access to that file.

You may choose to use a different location for the configuration file, in which case execute:

    orchestrator -c clusters --config=/path/to/config.file

> `-c` stands for `command`, and is mandatory.

Discover a new instance ("teach" `orchestrator` about your topology). `Orchestrator` will automatically recursively drill up the master chain (if any)
and down the replicas chain (if any) to detect the entire topology:

    orchestrator -c discover -i 127.0.0.1:22987

> `-i` stands for `instance` and must be in the form `hostname:port`.

Do the same, and be more verbose:

    orchestrator -c discover -i 127.0.0.1:22987 --debug
    orchestrator -c discover -i 127.0.0.1:22987 --debug --stack

> `--debug` can be useful in all operations. `--stack` prints code stack trace on (most) errors and is useful
> for development & testing purposed or for submitting bug reports.

Forget an instance (an instance may be manually or automatically re-discovered via `discover` command above):

    orchestrator -c forget -i 127.0.0.1:22987

Print an ASCII tree of topology instances. Pass a cluster name via `-i` (see `clusters` command above):

    orchestrator -c topology -i 127.0.0.1:22987

> Sample output:
>
>     127.0.0.1:22987
>     + 127.0.0.1:22989
>       + 127.0.0.1:22988
>     + 127.0.0.1:22990

Move the replica around the topology:

    orchestrator -c relocate -i 127.0.0.1:22988 -d 127.0.0.1:22987

> Resulting topology:
>
>     127.0.0.1:22987
>     + 127.0.0.1:22989
>     + 127.0.0.1:22988
>     + 127.0.0.1:22990

The above happens to move the replica one level up. However the `relocate` command accepts any valid destination.
`relocate` figures out the best way to move a replica. If GTID is enabled, use it. If Pseudo-GTID is available, use it.
If a binlog server is involved, use it. I `orchestrator` has further insight into the specific coordinates involved, use it. Otherwise just use plain-old binlog log file:pos math.

Similar to `relocate`, you can move multiple replicas via `relocate-replicas`. This moves replicas-of-an-instance below another server.

> Assume this:
>
>     10.0.0.1:3306
>     + 10.0.0.2:3306
>       + 10.0.0.3:3306
>       + 10.0.0.4:3306
>       + 10.0.0.5:3306
>     + 10.0.0.6:3306

    orchestrator -c relocate-replicas -i 10.0.0.2:3306 -d 10.0.0.6

> Results with:
>
>     10.0.0.1:3306
>     + 10.0.0.2:3306
>     + 10.0.0.6:3306
>       + 10.0.0.3:3306
>       + 10.0.0.4:3306
>       + 10.0.0.5:3306

> You may use `--pattern` to filter those replicas affected.

Other commands give you a more fine grained control on how your servers are relocated. Consider the _classic_ binary log file:pos
way of repointing replicas:

Move a replica up the topology (make it sbling of its master, or direct replica of its "grandparent"):

    orchestrator -c move-up -i 127.0.0.1:22988

> The above command will only succeed if the instance _has_ a grandparent, and does not have _problems_ such as replica lag etc.

Move a replica below its sibling:

    orchestrator -c move-below -i 127.0.0.1:22988 -d 127.0.0.1:22990 --debug

> The above command will only succeed if `127.0.0.1:22988` and `127.0.0.1:22990` are siblings (replicas of same master), none of them has _problems_ (e.g. replica lag),
> and the sibling _can_ be master of instance (i.e. has binary logs, has `log_slave_updates`, no version collision etc.)

Make an instance read-only or writeable:

    orchestrator -c set-read-only -i 127.0.0.1:22988
    orchestrator -c set-writeable -i 127.0.0.1:22988
