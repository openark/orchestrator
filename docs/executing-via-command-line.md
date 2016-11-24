# Executing via command line

Also consult the [Orchestrator first steps](first-steps.md) page.

Following is a synopsis of command line samples. For simplicity, we assume `orchestrator` is in your path.
If not, replace `orchestrator` with `/path/to/orchestrator`.

> Samples below use a test `mysqlsandbox` topology, where all instances are on same host `127.0.0.1` and on different ports. `22987` is master,
> and `22988`, `22989`, `22990` are replicas.

Show currently known clusters (replication topologies):

    orchestrator -c clusters cli

> The above looks for configuration in `/etc/orchestrator.conf.json`, `conf/orchestrator.conf.json`, `orchestrator.conf.json`, in that order.
> Classic is to put configuration in `/etc/orchestrator.conf.json`. Since it contains credentials to your MySQL servers you may wish to limit access to that file.

You may choose to use a different location for the configuration file, in which case execute:

    orchestrator -c clusters --config=/path/to/config.file cli

> `-c` stands for `command`, and is mandatory.

Discover a new instance ("teach" `orchestrator` about your topology). `Orchestrator` will automatically recursively drill up the master chain (if any)
and down the replicas chain (if any) to detect the entire topology:

    orchestrator -c discover -i 127.0.0.1:22987 cli

> `-i` stands for `instance` and must be in the form `hostname:port`.

Do the same, and be more verbose:

    orchestrator -c discover -i 127.0.0.1:22987 --debug cli
    orchestrator -c discover -i 127.0.0.1:22987 --debug --stack cli

> `--debug` can be useful in all operations. `--stack` prints code stack trace on (most) errors and is useful
> for development & testing purposed or for submitting bug reports.

Forget an instance (an instance may be manually or automatically re-discovered via `discover` command above):

    orchestrator -c forget -i 127.0.0.1:22987 cli

Print an ASCII tree of topology instances. Pass a cluster name via `-i` (see `clusters` command above):

    orchestrator -c topology -i 127.0.0.1:22987 cli

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

The above happens to move the replica one level up. However the `relocate` command accepts any valid destination. `relocate`
figures out the best way to move a replica. If GTID is enabled, use it. If Pseudo-GTID is available, use it. If a binlog server is
involved, use it. I `orchestrator` has further insight into the specific coordinates involved, use it. Otherwise just use
plain-old binlog log file:pos math.

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

Other command sgive you a more fine grained control on how your servers are relocated. Consider the _classic_ binary log file:pos
way of repointing replicas:

Move a replica up the topology (make it sbling of its master, or direct replica of its "grandparent"):

    orchestrator -c move-up -i 127.0.0.1:22988 cli

> The above command will only succeed if the instance _has_ a grandparent, and does not have _problems_ such as replica lag etc.

Move a replica below its sibling:

    orchestrator -c move-below -i 127.0.0.1:22988 -d 127.0.0.1:22990 --debug cli

> `-s` stands for `sibling`.

> The above command will only succeed if `127.0.0.1:22988` and `127.0.0.1:22990` are siblings (replicas of same master), none of them has _problems_ (e.g. replica lag),
> and the sibling _can_ be master of instance (i.e. has binary logs, has `log_slave_updates`, no version collision etc.)

Promote a replica to be co-master with its master, making for a circular Master-Master topology:

    orchestrator -c make-co-master -i 127.0.0.1:22988 cli

> The above command will only succeed if `127.0.0.1:22988`'s master is root of topology (is not itself a replica)
> and is not associated in another co-master ring.

Reset a replica, effectively breaking down the replication (destructive action):

    orchestrator -c reset-slave -i 127.0.0.1:22988 cli

> *A note on topology refactoring commands*
>
> `move-up`, `move-below`, `make-co-master` and `reset-slave` are the building blocks of _classic_ topology refactoring.
> With the first two actions one can make any change to the topology, with the exception of moving the master.
> The last two allow replacing a master by promoting one of its replicas to be a co-master (MySQL master-master
> replication), then resetting the newly promoted co-master, effectively making it the master of all topology.

> These actions are also as atomic as possible, by only affecting two replication servers per action (e.g. `move-up` affects
> the instance and its master; `move-below` affect the instance and its sibling).

> The word _classic_ relates to the method of using an up-and-alive topology, where all connections are good and
> instances can be queried for their replication status.

> However `orchestrator` also supports topology refactoring in situations where servers are inaccessible.
> This could made to work via GTID and Pseudo-GTID.
>
> It may allow promoting a replica up the topology even as its master is dead, or
> matching and synching the replicas of a failed master even though they all stopped replicating in different
> positions.

The following are Pseudo-GTID specific commands:

Match a replica below another instance (we expect the other instance to be as advanced or more advanced than the moved replica)

    orchestrator -c match-below -i 127.0.0.1:22988 -d 127.0.0.1:22990 --debug cli

> The above required Pseudo GTID to be present and configured. It may take more time to execute as it needs to
> look up entires in the servers binary log.
> See [Pseudo GTID](#pseudo-gtid) for more details.

Make an instance read-only or writeable:

    orchestrator -c set-read-only -i 127.0.0.1:22988 cli
    orchestrator -c set-writeable -i 127.0.0.1:22988 cli

Begin maintenance mode on an instance. While in maintenance mode, `orchestrator` will not allow this instance to
be moved or participate in another instance's move:

    orchestrator -c begin-maintenance -i 127.0.0.1:22988 --reason="load testing; do not disturb" cli

End maintenance mode on an instance:

    orchestrator -c end-maintenance -i 127.0.0.1:22988 cli

Make an infinite, continuous discovery and investigation of known instances. Typically this is what the web service executes.

    orchestrator -c continuous --debug cli

Just check if an instance can be connected: attempt to resolve hostname and make a TCP connection to host & port:

    orchestrator -c resolve -i myhost.mydomain:3306
