FROM golang:1.17-alpine

RUN apk add --no-cache curl

WORKDIR /app

COPY main.go .

RUN go build -o app main.go

CMD [ "./app" ]
