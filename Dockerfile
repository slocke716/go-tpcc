FROM bitnami/golang:1.17

WORKDIR /opt/bitnami/go/src/go-tpcc
COPY src .
RUN go build .
ENTRYPOINT ["/opt/bitnami/go/src/go-tpcc/go-tpcc"]
