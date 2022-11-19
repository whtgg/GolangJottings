FROM golang:1.19.3-alpine

#RUN go get github.com/hashicorp/consul/api

ENV SOURCES /go/src/github.com/baoqger/service-discovery-demo/

COPY . ${SOURCES}

ENV CONSUL_HTTP_ADDR localhost:8500

WORKDIR ${SOURCES}

CMD ${SOURCES}server





