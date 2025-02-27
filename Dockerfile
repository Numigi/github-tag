FROM golang:1.12

RUN mkdir -p /opt/go
WORKDIR /opt/go

COPY go.mod go.sum /opt/go/
RUN go mod download

COPY pkg /opt/go/pkg
