# Configuration: discovery, Pseudo-GTID

`orchestrator` will identify magic hints in the binary logs, making it able to manipulate a non-GTID topology as if it had GTID, including relocation of replicas, smart failovers end more.

See [Pseudo GTID](pseudo-gtid.md)

If you wish to use Pseudo-GTID, we suggest you should configure as follows:
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
- You will inject Pseudo-GTID entries via [this sample script](https://github.com/github/orchestrator/tree/master/resources/pseudo-gtid)
