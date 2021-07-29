FROM golang as build

ENV GOPROXY=https://goproxy.io

ADD . /soda

WORKDIR /soda

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server

FROM alpine:latest

RUN echo "http://mirrors.aliyun.com/alpine/latest-stable/main/" > /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf && \
    mkdir -p /www

WORKDIR /www

COPY --from=build /soda/api_server /usr/bin/api_server

RUN chmod +x /usr/bin/api_server

ENTRYPOINT ["api_server"]