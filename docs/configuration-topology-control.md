# Configuration: topology control

The following configuration affects how `orchestrator` applies changes to topology servers:

`orchestrator` will figure out the name of the cluster, data center, and more.

```json
{
  "UseSuperReadOnly": false,
}
```

### UseSuperReadOnly

By default `false`. When `true`, whenever `orchestrator` is asked to set/clear `read_only`, it will also apply the change to `super_read_only`. `super_read_only` is only available on Oracle MySQL and Percona Server, as of specific versions.
