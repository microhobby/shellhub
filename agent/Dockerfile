# base stage
FROM golang:1.13-alpine3.11 AS base

RUN apk add --update git ca-certificates build-base bash util-linux setpriv

RUN ln -sf /bin/bash /bin/sh

WORKDIR $GOPATH/src/github.com/shellhub-io/shellhub

COPY ./go.mod ./

WORKDIR $GOPATH/src/github.com/shellhub-io/shellhub/agent

COPY ./agent/go.mod ./agent/go.sum ./

RUN go mod download

# builder stage
FROM base AS builder

ARG SHELLHUB_VERSION=latest

COPY ./pkg $GOPATH/src/github.com/shellhub-io/shellhub/pkg
COPY ./agent .

WORKDIR $GOPATH/src/github.com/shellhub-io/shellhub

RUN go mod download

WORKDIR $GOPATH/src/github.com/shellhub-io/shellhub/agent

RUN go build -tags docker -ldflags "-X main.AgentVersion=${SHELLHUB_VERSION}"

# development stage
FROM base AS development

RUN apk add --update openssl openssh-client
RUN go get github.com/markbates/refresh && go get github.com/mgechev/revive

WORKDIR $GOPATH/src/github.com/shellhub-io/shellhub

RUN go mod download

#RUN cp -a $GOPATH/src/github.com/shellhub-io/shellhub/vendor /vendor

COPY ./agent/entrypoint-dev.sh /entrypoint.sh
COPY ./.revive.toml /

WORKDIR $GOPATH/src/github.com/shellhub-io/shellhub/agent

ENTRYPOINT ["/entrypoint.sh"]

# production stage
FROM alpine:3.11 AS production

WORKDIR /app
COPY --from=builder /agent /app/

ENTRYPOINT ./agent
