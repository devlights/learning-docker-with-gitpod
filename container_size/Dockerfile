FROM golang:1.17-alpine as build

WORKDIR /app
COPY main.go .
RUN go build -o app main.go

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/app .

ENTRYPOINT [ "./app" ]
