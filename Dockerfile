ARG ARCH="amd64"
FROM ${ARCH}/alpine

ARG OS="linux"
ARG BIN_ARCH="amd64"

ENV PROJECT_NAME="aws-cwa-sns-google-chat"
ENV HOME="/app"

LABEL name="${PROJECT_NAME}" \
  org.opencontainers.image.url="https://github.com/slashdevops/aws-cwa-sns-google-chat" \
  org.opencontainers.image.source="https://github.com/slashdevops/aws-cwa-sns-google-chat"

RUN apk add --no-cache --update \
  ca-certificates \
  && rm -rf /tmp/* /var/tmp/* /var/cache/apk/*

RUN mkdir -p $HOME && \
  chown -R nobody.nobody $HOME

COPY dist/$PROJECT_NAME-$OS-$BIN_ARCH/* $HOME/

ENV PATH="${PATH}:${HOME}"

VOLUME $HOME
USER nobody:nobody
WORKDIR $HOME

CMD ["/app/idpscim", "--help"]