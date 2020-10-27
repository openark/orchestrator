# Gotchas

* By default, `orchestrator` only polls a server once a minute (configurable via `InstancePollSeconds`). This means that any status you see is essentially an estimation. Different instances get polled at different times. The status you see on the _cluster_ page, for example, does not necessarily reflect a given point in time, but rather a combination of different points in time in the last minute (or whatever poll interval you use).

The `problems` drop down to the right is also executed asynchronously. You may therefore see the same instance in two
places (once in the `cluster` page, once in the `problems` drop down) in two different states.

If you want to get the most up to date instance status, use the "Refresh" button on the instance's _settings dialog_.  

* It could take a couple of minutes for `orchestrator` to fully detect a cluster's topology. The time depends on the depth
of the topology (if you have replicas-of-replicas the time increases). This is due to the fact that `orchestrator` polls the instances independently, and an insight on the topology must propagate from master to replica on the next polling occasion.

* Specifically, if you fail over to a new master, you may find that for a couple minutes the topologies seem empty.
This may happen because instances used to identify themselves as belonging to a certain topology that is now being destroyed.
This is self-healing. Refresh and look at the `Clusters` menu to review the newly created cluster (names after the new master)
over time.

* Don't restart `orchestrator` while you're running a seed (only applies to working with _orchestrator-agent_)

Otherwise `orchestrator` is non-intrusive and self-healing. You can restart it whenever you like.
