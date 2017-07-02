/*
   Copyright 2016 GitHub Inc.
	 See https://github.com/github/orchestrator/blob/master/LICENSE
*/

package app

const AppPrompt string = `
orchestrator [-c command] [-i instance] [-d destination] [--verbose|--debug] [... cli ] | http

Cheatsheet:
    Run orchestrator in HTTP mode:

        orchestrator --debug http

    See all possible commands:

        orchestrator -c help

    Detailed help for a given command (e.g. "relocate")

        orchestrator help relocate
      or
        orchestrator -c relocate help

    Usage for most commands:
        orchestrator -c <command> [-i <instance.fqdn>] [-d <destination.fqdn>] [--verbose|--debug]

    -i (instance):
        instance on which to operate, in "hostname" or "hostname:port" format.
        Default port is 3306 (or DefaultInstancePort in config)
        For some commands this argument can be ommitted altogether, and the
        value is implicitly the local hostname.
    -d (Destination)
        destination instance (used when moving replicas around or when failing over)
    -s (Sibling/Subinstance/deStination) - synonym to "-d"

    -c (command):
        Listed below are all available commands; all of which apply for CLI execution (ignored by HTTP mode).
        Different flags are required for different commands; see specific documentation per commmand.
`
