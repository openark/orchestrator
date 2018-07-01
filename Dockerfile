# Requires Docker 17.09 or later (multi stage builds)

FROM alpine

ENV GOPATH=/tmp/go

RUN apk update
RUN apk add --no-cache libcurl ; apk add libcurl
RUN apk add --no-cache rsync
RUN apk add --no-cache go
RUN apk add --no-cache gcc
RUN apk add --no-cache g++
#RUN apk add --no-cache build-base

RUN mkdir -p $GOPATH/src/github.com/github/orchestrator
WORKDIR $GOPATH/src/github.com/github/orchestrator
COPY . .
RUN \
    && bash build.sh -b \
    && rsync -av $(find /tmp/orchestrator-release -type d -name orchestrator -maxdepth 2)/ / \
    && rsync -av $(find /tmp/orchestrator-release -type d -name orchestrator-cli -maxdepth 2)/ / \
    && cd / \
    && apk del .build-deps \
    && rm -rf /tmp/*

FROM alpine

EXPOSE 3000

COPY --from=0 /usr/local/orchestrator /usr/local/orchestrator

WORKDIR /usr/local/orchestrator
ADD docker/entrypoint.sh /entrypoint.sh
CMD /entrypoint.sh
