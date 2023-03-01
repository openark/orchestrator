module github.com/openark/orchestrator

go 1.16

require (
	github.com/armon/consul-api v0.0.0-20180202201655-eb2c6b5be1b6
	github.com/codegangsta/inject v0.0.0-20150114235600-33e0aa1cb7c0 // indirect
	github.com/cyberdelia/go-metrics-graphite v0.0.0-20161219230853-39f87cc3b432
	github.com/deathowl/go-metrics-prometheus v0.0.0-20200518174047-74482eab5bfb // indirect
	github.com/fatih/color v1.10.0 // indirect
	github.com/go-martini/martini v0.0.0-20170121215854-22fa46961aab
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/btree v1.0.0 // indirect
	github.com/hashicorp/consul/api v1.7.0
	github.com/hashicorp/go-cleanhttp v0.5.2-0.20190406162018-d3fcbee8e181 // indirect
	github.com/hashicorp/go-hclog v0.15.1-0.20201116205511-59fbd7b93270 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.0 // indirect
	github.com/hashicorp/go-rootcerts v1.0.3-0.20191216101743-c8a9a31cbd76 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/raft v0.0.0-00010101000000-000000000000
	github.com/hashicorp/serf v0.9.5 // indirect
	github.com/howeyc/gopass v0.0.0-20190910152052-7cb4b85ec19c
	github.com/martini-contrib/auth v0.0.0-20150219114609-fa62c19b7ae8
	github.com/martini-contrib/gzip v0.0.0-20151124214156-6c035326b43f
	github.com/martini-contrib/render v0.0.0-20150707142108-ec18f8345a11
	github.com/mattn/go-isatty v0.0.13-0.20200128103942-cb30d6282491 // indirect
	github.com/miekg/dns v1.1.31 // indirect
	github.com/mitchellh/go-testing-interface v1.14.0 // indirect
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/montanaflynn/stats v0.0.0-20161102194025-f8cd06f93c6c
	github.com/openark/golib v0.0.0-20210520103621-827f3ea62180
	github.com/outbrain/golib v0.0.0-20200503083229-2531e5dbcc71 // indirect
	github.com/outbrain/zookeepercli v1.0.12
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/prometheus/client_golang v1.11.0
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475
	github.com/samuel/go-zookeeper v0.0.0-20201211165307-7117e9ea2414
	github.com/sjmudd/stopwatch v0.0.0-20170103085848-637ef30077b7
	github.com/stretchr/testify v1.6.1 // indirect
	golang.org/x/crypto v0.0.0-20200930160638-afb6bcd081ae // indirect
	golang.org/x/net v0.0.0-20200930145003-4acb6c075d10 // indirect
	gopkg.in/gcfg.v1 v1.2.3
	gopkg.in/warnings.v0 v0.1.2 // indirect
)

replace github.com/hashicorp/raft => github.com/openark/raft v0.0.0-20170918052300-fba9f909f7fe
