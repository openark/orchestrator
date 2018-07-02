FROM alpine

ENV GOPATH=/tmp/go

RUN set -ex \
    && apk add --update --no-cache bash \
    && apk add --update --no-cache --virtual .build-deps \
        rsync \
        git \
        go \
        build-base \
    && cd /tmp \
    && { go get -d github.com/github/orchestrator ; : ; } \
    && cd $GOPATH/src/github.com/github/orchestrator \
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
