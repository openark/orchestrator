# Requires Docker 17.09 or later (multi stage builds)
#
# Orchestrator will look for a configuration file at /etc/orchestrator.conf.json.
# It will listen on port 3000.
# If not present a minimal configuration will be generated using the following environment variables:
#
# Default variables which can be used are:
#
# ORC_TOPOLOGY_USER (default: orchestrator): username used by orchestrator to login to MySQL when polling/discovering
# ORC_TOPOLOGY_PASSWORD (default: orchestrator):  password needed to login to MySQL when polling/discovering
# ORC_DB_HOST (default: orchestrator):  orchestrator backend MySQL host
# ORC_DB_PORT (default: 3306):  port used by orchestrator backend MySQL server
# ORC_DB_NAME (default: orchestrator): database named used by orchestrator backend MySQL server
# ORC_USER (default: orc_server_user): username used to login to orchestrator backend MySQL server
# ORC_PASSWORD (default: orc_server_password): password used to login to orchestrator backend MySQL server

FROM alpine:3.6

ENV GOPATH=/tmp/go

RUN apk update
RUN apk upgrade
RUN apk add --update libcurl
RUN apk add --update rsync
RUN apk add --update gcc
RUN apk add --update g++
RUN apk add --update go
RUN apk add --update build-base
RUN apk add --update bash
RUN apk add --update git

RUN mkdir -p $GOPATH/src/github.com/github/orchestrator
WORKDIR $GOPATH/src/github.com/github/orchestrator
COPY . .
RUN bash build.sh -b
RUN rsync -av $(find /tmp/orchestrator-release -type d -name orchestrator -maxdepth 2)/ /
RUN rsync -av $(find /tmp/orchestrator-release -type d -name orchestrator-cli -maxdepth 2)/ /
RUN cp /usr/local/orchestrator/orchestrator-sample-sqlite.conf.json /etc/orchestrator.conf.json

FROM alpine:3.6

EXPOSE 3000

COPY --from=0 /usr/local/orchestrator /usr/local/orchestrator
COPY --from=0 /etc/orchestrator.conf.json /etc/orchestrator.conf.json

WORKDIR /usr/local/orchestrator
ADD docker/entrypoint.sh /entrypoint.sh
CMD /entrypoint.sh
