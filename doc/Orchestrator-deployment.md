This text discusses deployment options for _orchestrator_.

It is assumed you already know how to install _orchestrator_ on a production machine and how to configure it with a backend database. It is also assumed you have configured your MySQL servers to allow connections from _orchestrator_.
If not, please refer to the [installation section](Orchestrator-Manual#installation)

The following hints should clear up some questions:

- _Orchestrator_ can run as a linux service
  - When running as a service, it assumes the role of continuous discovery: a never ending polling of your MySQL topologies
  - When running as a service, it also provides HTTP API/Web interface
- _Orchestrator_ can run in command line mode
  - Useful for us CLI geeks who want to have greater control and capture all the lovely debug messages

###### In terms of deployment

- _Orchestrator_ uses a MySQL backend.
- The backend database has the _state_ of your topologies. _Orchestrator_ is almost stateless.
  - _Orchestrator_ only has state for pending operations (e.g. while a slave is being moved)
- The backend databse will have the state of your multiple topologies.  
- You _can_ and _should_ have more than one _orchestrator_ running with the same MySQL backend.
- You should _not_ have more than one MySQL backend.

The MySQL backend is a master server, for which you may have slaves for redundancy or otherwise build your favorite HA solution.

The author of _orchestrator_ has deployed it in large environments of thousands of servers, with a single backend database managing them all.

###### Orchestrator service

- You _can_ and _should_ have multiple _orchestrator_ services running on different machines, all running on the same backend database.
  - When running as a service, _orchestrator_ repeatedly attempts to claim _leadership_. Should one _orchestrator_ service
    fail, another one will pick up where it left.
  - Only one service will be the leader at any given time.
  - The leader is the one polling for servers; doing database cleanup; checking for crash recovery scenarios etc.
- You may choose to have all your _orchestrator_ services load-balanced

###### Orchestrator CLI
- You _can_ and _should_ have as many deployments of orchestrator CLI as you like, on multiple servers, all configured to work
  with the _same backend MySQL_
- You _may_ concurrently issue commands from multiple CLIs, as well as working with the Web UI at the same time.
- The (single) MySQL backend has the necessary state for managing concurrent operations.
- _Orchestrator_ has "maintenance locks" which prevent destructive concurrent operations on the same instance. At worst an
  operation will be rejected due to not being able to acquire maintenance lock.

The author of _orchestrator_ has it deployed on multiple machines as a service, behind a F5 load balancer. On the same setup,
CLI is deployed and can be executed from thousands of machines.

##### A visualized example

![Orchestrator deployment](images/orchestrator-deployment.png)

In the above four _orchestrator_ services are behind an HTTP load balancer. Only one of them is the _leader_ at any given time; they compete for leadership between themselves and recognize if the leader is non doing its duty.

To the right: the many topologies polled by _orchestrator_. The leader polls each an every server in those topologies.

On top left: the _orchestrator_ MySQL backend: a master & three slaves. All 4 services use the same backend database.

Not shown in this picture (for clarity purposes), but the _orchestrator_ backend database and its slaves are themselves one of those topologies
polled by _orchestrator_ It eats its own dogfood.

On left, bottom: _orchestrator_ CLI is installed on any and all MySQL servers. All these CLI deployments use the very same
MySQL backend database.

Not shown in this picture (for clarity purposes), the MySQL servers on these hosts are being polled by _orchestrator_ just as those
on the right.

You may choose to install _orchestrator_ on non-MySQL hosts, of course. It has no MySQL dependencies on the deployed host.
