# syntax=docker/dockerfile:1-labs
FROM golang:1.17-alpine

RUN <<EOF
    apk add --no-cache make gcc libc-dev musl-dev
    go install github.com/go-delve/delve/cmd/dlv@latest
EOF

WORKDIR /workspace
