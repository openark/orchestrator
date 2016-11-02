# Requirements

`Orchestrator` is a standalone Go application. It requires a MySQL backend to store topologies state, maintenance status and audit history.
It is built and tested on Linux 64bit, and binaries are availably for this OS type alone. The author has not tested any other operating system.
However the [build script](https://github.com/github/orchestrator/blob/master/build.sh) is capable of building `orchestrator` for other OS/architectures.
