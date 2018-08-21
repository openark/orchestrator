apk add -U sqlite
apk add -U tzdata

cp /usr/share/zoneinfo/America/Los_Angeles /etc/localtime

time_start=$(date)

rules_at_start=$(sqlite3 orchestrator.sqlite3 "SELECT hostname, port, promotion_rule, last_suggested FROM candidate_database_instance")
echo 'Rules at start:'
echo $rules_at_start

echo 'Tag host'
promotion_rule='prefer'
./orchestrator -c register-candidate --promotion-rule $promotion_rule

prom_rules=$(sqlite3 orchestrator.sqlite3 "SELECT hostname, port, promotion_rule, last_suggested FROM candidate_database_instance" | wc -l)
while [ $prom_rules -ne 0 ]; do
  prom_rules=$(sqlite3 orchestrator.sqlite3 "SELECT hostname, port, promotion_rule, last_suggested FROM candidate_database_instance" | wc -l)
  sleep 1
done

echo 'Started at:'
echo $time_start
echo 'Exiting at:'
date
