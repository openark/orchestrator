# Configuration: discovery, Pseudo-GTID

`orchestrator` will identify magic hints in the binary logs, making it able to manipulate a non-GTID topology as if it had GTID, including relocation of replicas, smart failovers end more.

See [Pseudo GTID](pseudo-gtid.md)

### Automated Pseudo-GTID injection

`orchestrator` can inject Pseudo-GTID entries for you and save you the trouble. You will configure:
```json
{
  "AutoPseudoGTID": true,
}
```
And you may ignore any other Pseudo-GTID related configuration (they will all be implicitly overriden by `orchestrator`).

You will further need to grant the following on your MySQL servers:
```sql
GRANT DROP ON _pseudo_gtid_.* to 'orchestrator'@'orch_host';
```

**NOTE**: the `_pseudo_gtid_` schema doesn't need to exist. There is no need to create it. `orchestrator` will run queries of the form:

```sql
drop view if exists `_pseudo_gtid_`.`_asc:5a64a70e:00000001:c7b8154ff5c3c6d8`
```

Those statements will do nothing but will serve as magic markers in the binary logs.

`orchestrator` will only attempt injecting Pseudo-GTID where it is allowed to do so. If you want to limit Pseudo-GTID injection to specific clusters, you may do so by only granting the privilege on those clusters you want `orchestrator` to inject Pseudo-GTID. You may disable Pseudo-GTID injection on a specific cluster via:

```sql
REVOKE DROP ON _pseudo_gtid_.* FROM 'orchestrator'@'orch_host';
```

Automated Pseudo-GTID injection is a newer development which supersedes the need for you to run your own Pseudo-GTID injection.

If you wish to enable auto-Pseudo-GTID injection having run manual Pseudo-GTID injection, you'll be happy to note that:

- You will no longer need to manage a pseudo-GTID service / event scheduler.
- And in particular you will not need to disable/enable Pseudo-GTID on old/promoted master upon master failover.

### Manual Pseudo-GTID injection

[Automated Pseudo-GTID](#automated-pseudo-gtid-injection) is the recommended method.

If you wish to inject Pseudo-GTID yourself, we suggest you should configure as follows:

```json
{
  "PseudoGTIDPattern": "drop view if exists `meta`.`_pseudo_gtid_hint__asc:",
  "PseudoGTIDPatternIsFixedSubstring": true,
  "PseudoGTIDMonotonicHint": "asc:",
  "DetectPseudoGTIDQuery": "select count(*) as pseudo_gtid_exists from meta.pseudo_gtid_status where anchor = 1 and time_generated > now() - interval 2 hour",
}
```

The above assumes:

- You have a `meta` schema
- You will inject Pseudo-GTID entries via [this sample script](https://github.com/openark/orchestrator/tree/master/resources/pseudo-gtid)
