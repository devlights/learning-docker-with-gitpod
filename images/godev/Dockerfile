# syntax=docker/dockerfile:1-labs
FROM golang:1.17-alpine

RUN <<EOF
        apk add --no-cache vim gcc make libc-dev musl-dev git
        go install github.com/go-delve/delve/cmd/dlv@latest
EOF

WORKDIR /workspace
