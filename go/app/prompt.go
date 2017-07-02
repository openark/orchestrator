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

        orchestrator help

    Detailed help for a given command (e.g. "relocate")

        orchestrator help relocate
`
