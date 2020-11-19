# Tags

`orchestrator` supports tagging of instances and searching by tags.

Tagging is provided as a service to the user and is not used internally by `orchestrator`.

### Tag commands

The following commands are supported. A breakdown follows:

- `orchestrator-client -c tag -i some.instance --tag name=value`
- `orchestrator-client -c tag -i some.instance --tag name`
- `orchestrator-client -c untag -i some.instance -t name`
- `orchestrator-client -c untag-all -t name=value`
- `orchestrator-client -c tags -i some.instance`
- `orchestrator-client -c tag-value -i some.instance -t name`
- `orchestrator-client -c tagged -t name`
- `orchestrator-client -c tagged -t name=value`
- `orchestrator-client -c tagged -t name=`

and these API endpoints:

- `api/tag/:host/:port/:tagName/:tagValue`
- `api/tag/:host/:port?tag=name`
- `api/tag/:host/:port?tag=name%3Dvalue`
- `api/untag/:host/:port/:tagName`
- `api/untag/:host/:port?tag=name`
- `api/untag-all/:tagName/:tagValue`
- `api/untag-all?tag=name%3Dvalue`
- `api/tags/:host/:port`
- `api/tag-value/:host/:port/:tagName`
- `api/tag-value/:host/:port?tag=name`
- `api/tagged?tag=name`
- `api/tagged?tag=name%3Dvalue`
- `api/tagged?tag=name%3D`

### Tags, general

A tag can be of the form `name=value`, or can be of the form `name`, in which case the value is implicitly set as the empty string.
A `name` can be of the following formats:

- `word`
- `some-other-word`
- `some_word_word_yet`

Though not strictly enforced, avoid using special characters/punctuations.

### Tagging

`-c tag` or `api/tag` adds or replaces an existing tag to an instance. `orchestrator` will not indicate whether the tag pre-existed beforehand, nor will it offer the previous value, if any.

Example:
```shell
$ orchestrator-client -c tag -i db-host-01:3306 --tag vttablet_alias=dc1-0123456789
```
In the above we chose to create a tag named `vttablet_alias` with a value.

Tagging is per instance. The instance itself is unaffected by this operation. `orchestrator` maintains tags as metadata. The instance needs not be available.

### Untagging: single instance

`-c untag` or `api/untag` removes a tag, if exists, from a given instance. `orchestrator` outputs the instance name if the tag did in fact exist, or empty output if the tag did not exist.

You may tags:

- Specify tag name and tag value: tag is removed only if it equals that value.
- Specify tag name only: tag is removed regardless of its value.

Example:
```shell
$ orchestrator-client -c untag -i db-host-01:3306 --tag vttablet_alias
```

### Untagging: multiple instances

`-c untag-all` or `api/untag-all` removes a tag from all instances where the value matches. Note that tag value must be provided.

Example:
```shell
$ orchestrator-client -c untag-all --tag vttablet_alias=dc1-0123456789
```

### Listing instance tags

For a given instance `-c tags` or `api/tags` lists all known tags.

Example:
```shell
$ orchestrator-client -c tag -i db-host-01:3306 --tag vttablet_alias=dc1-0123456789
$ orchestrator-client -c tag -i db-host-01:3306 --tag old-hardware

$ orchestrator-client -c tags -i db-host-01:3306
old-hardware=
vttablet_alias=dc1-0123456789
```

listed tags are sorted by name.
Note that we added `old-hardware` tag without value. It exports as `old-hardware=`, with implicit empty value.


### Listing instance tags for a cluster

For a given instance or cluster alias `-c topology-tags` or `api/topology-tags` lists the cluster topology with all known tags for each instance.

Example:
```shell
$ orchestrator-client -c tag -i db-host-01:3306 --tag vttablet_alias=dc1-0123456789
$ orchestrator-client -c tag -i db-host-01:3306 --tag old-hardware

$ orchestrator-client -c topology-tags -alias mycluster
db-host-01:3306     [0s,ok,5.7.23-log,rw,ROW,>>,GTID,P-GTID] [vttablet_alias=dc1-0123456789, old-hardware]
+ db-host-02:3306   [0s,ok,5.7.23-log,ro,ROW,>>,GTID,P-GTID] []

$ orchestrator-client -c topology-tags -i db-host-01:3306
db-host-01:3306     [0s,ok,5.7.23-log,rw,ROW,>>,GTID,P-GTID] [vttablet_alias=dc1-0123456789, old-hardware]
+ db-host-02:3306   [0s,ok,5.7.23-log,ro,ROW,>>,GTID,P-GTID] []
```


### Getting the value of a specific tag

`-c tag-value` or `api/tag-value` return the value of a specific tag on an instance.

Example:
```shell
$ orchestrator-client -c tag -i db-host-01:3306 --tag vttablet_alias=dc1-0123456789
$ orchestrator-client -c tag -i db-host-01:3306 --tag old-hardware
$ orchestrator-client -c tag-value -i db-host-01:3306 --tag vttablet_alias
dc1-0123456789
$ orchestrator-client -c tag-value -i db-host-01:3306 --tag old-hardware

# <empty value>
$ orchestrator-client -c tag-value -i db-host-01:3306 --tag no-such-tag
tag no-such-tag not found for db-host-01:3306
# in stderr
```

### Searching instances by tags

`-c tagged` or `api/tagged` lists instances by tags, as follows:

- `-c tagged -tag name=value`: list instances where `name` exists and equals `value`.
- `-c tagged -tag name`: list instances where `name` exists, regardless of the value.
- `-c tagged -tag name=`: list instances where `name` exists and has an empty value.
- `-c tagged -tag name,role=backup`: list instances tagged by `name` (regardless of its value) and are _also_ tagged with `role=backup`
- `-c tagged -tag !name`: list instances where no tag called `name` exists, regardless of its value
- `-c tagged -tag ~name`: `~` is a synonym to `!`.
- `-c tagged -tag name,~role`: list instances tagged by `name` (regardless of its value) and are _not_ tagged by `role` (regardless of its value)
- `-c tagged -tag ~role=backup`: list instances that _are_ tagged with `role`, but with value other than `backup`.
  Notice how this differs from `-c tagged -tag ~role` which will list instances which don't have the `role` tag in the first place.

### Tags, internal

Tags are associated with instances, but the association is internal to `orchestrator` and does not affect the actual MySQL instances.
Tags are stored internally on a backend table. Tags are advertised by the `raft` protocol; a tagging operation (`tag`, `untag`) carried out by the `raft` leader, will be applied on the `raft` followers.

### Use cases

The need for tags has come up from multiple users with differing use cases.

- A common use case for [Vitess](http://github.com/vitess.io/vitess) users is the need to associate an instance with a `vttablet` alias.
- Users may wish to apply promotion logic based on tags. While `orchestrator` does not use tagging internally in any decision making, users may set `promotion-rule` based on tags, or apply different failover operations.
