FROM golang:1.6-alpine

EXPOSE  3000

ADD . /go/src/github.com/github/orchestrator
WORKDIR /go/src/github.com/github/orchestrator

RUN set -ex \
    && apk add --no-cache --virtual .build-deps \
        bash \
        rsync \
        git \
    && bash build.sh -b \
    && rsync -av $(find /tmp/orchestrator-release -type d -name orchestrator -maxdepth 2)/ / \
    && rsync -av $(find /tmp/orchestrator-release -type d -name orchestrator-cli -maxdepth 2)/ / \
    && apk del .build-deps \
    && rm -rf /tmp/orchestrator-release \
    && cp ./docker/entrypoint.sh /entrypoint.sh

CMD /entrypoint.sh
