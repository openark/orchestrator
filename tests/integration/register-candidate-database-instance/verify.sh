#!/bin/bash

num_results=$(sqlite3 /tmp/orchestrator.db "SELECT hostname, port, promotion_rule, last_suggested FROM candidate_database_instance where last_suggested > datetime('now', ('-60 MINUTE'))" | wc -l)
if [ $num_results -ne 1 ] ; then
    echo "ERROR did not find 1 row in candidate_database_instance, found $num_results"
    return 1
fi
