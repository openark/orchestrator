# Orchestrator configuration in larger environments

If monitoring a large number of servers the backend database can become a bottleneck.   The following comments refer to using MySQL as the orchestrator backend.

Some configuration options allow you to control the throughput.  These settings are:
- `BufferInstanceWrites`
- `InstanceWriteBufferSize`
- `InstanceFlushIntervalMilliseconds`
- `DiscoveryMaxConcurrency`

Limit the number of concurrent discoveries made by orchestrator using `DiscoveryMaxConcurrency` and ensure that the backend server's `max_connections` setting is high enough to allow orchestrator to make as many connections as it needs.

By setting `BufferInstanceWrites: True` in orchestrator when a poll completes the results will be buffered until `InstanceFlushIntervalMilliseconds` has elapsed or `InstanceWriteBufferSize` buffered writes have been made.

The buffered writes are ordered by the time of the write using a single `insert ... on duplicate key update ...` call.  If the same host appears twice only the last write will be written to the database for that host.

`InstanceFlushIntervalMilliseconds` *should be* well below `InstancePollSeconds` as making this value too high will mean the data is not being written to the orchestrator db backend.  This can lead to `not recently checked` problems.  Also the different health checks are run against the backend database state so not updating it frequently enough could lead to Orchestrator not detecting the different failure scenarios correctly.

Suggested values to start with for larger Orchestrator environments might be:
```
  ...
  "BufferInstanceWrites": true,
  "InstanceWriteBufferSize": 1000, 
  "InstanceFlushIntervalMilliseconds": 50,
  "DiscoveryMaxConcurrency": 1000,
  ...
```
