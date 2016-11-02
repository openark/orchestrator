# Agents

You may optionally install [orchestrator-agent](https://github.com/github/orchestrator-agent) on your MySQL hosts.
*orchestrator-agent* is a service which registers with your *orchestrator* server and accepts requests by *orchestrator*
via web API.

Supported requests relate to general, OS and LVM operations, such as:
- Stopping/starting MySQL service on host
- Getting MySQL OS info such as data directory, port, disk space usage
- Performing various LVM operations such as finding LVM snapshots, mounting/unmounting a snapshot
- Transferring data between hosts (e.g. via `netcat`)

`orchestrator-agent` is developed at [Outbrain](https://github.com/outbrain) for Outbrain's specific requirements, and is
less of a general solution. As such, it supports those operations required by Outbrain's architecture. For example, we
rely on LVM snapshots for backups; we use a directory service to register available snapshots; we are DC-aware and prefer
local backups over remote backups.

Nevertheless at the very least `orchestrator-agent` should appeal to most. It is configurable to some extent (directory
service command is configurable - write your own bash code; data transfer command is configurable - replace `netcat` with
your own prefered method, etc.).

The information and API exposed by *orchestrator-agent* to *orchestrator* allow *orchestrator* to coordinate and operate
seeding of new or corrupted machines by getting data from freshly available snapshots. Moreover, it allows *orchestrator*
to automatically suggest the source of data for a given MySQL machine, by looking up such hosts that actually have a
recent snapshot available, preferably in the same datacenter.

For security measures, an agent requires a token to operate all but the simplest requests. This token is randomly generated
by the agent and negotiated with *orchestrator*. *Orchestrator* does not expose the agent's token (right now some work
needs to be done on obscurring the token on error messages).
