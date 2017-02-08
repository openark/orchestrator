FROM alpine

EXPOSE  3000

ENV GOPATH=/tmp/go

RUN set -ex \
    && apk add --update --no-cache --virtual .build-deps \
        bash \
        rsync \
        git \
        go \
        build-base \
    && cd /tmp \
    && { go get -d github.com/github/orchestrator ; : ; } \
    && cd $GOPATH/src/github.com/github/orchestrator \
    && bash build.sh -b \
    && cp ./docker/entrypoint.sh /entrypoint.sh \
    && rsync -av $(find /tmp/orchestrator-release -type d -name orchestrator -maxdepth 2)/ / \
    && rsync -av $(find /tmp/orchestrator-release -type d -name orchestrator-cli -maxdepth 2)/ / \
    && cd / \
    && apk del .build-deps \
    && rm -rf /tmp/*

CMD /entrypoint.sh
