FROM alpine:latest

RUN apk add --no-cache tzdata vim make musl-dev go curl
RUN cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
RUN echo "Asia/Tokyo" > /etc/timezone

ENV TZ Asia/Tokyo
ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP.UTF-8
ENV LC_ALL ja_JP.UTF-8

ENV GOPATH /root/go
ENV PATH ${GOPATH}/bin:/usr/local/go/bin:${PATH}

CMD [ "sh" ]
