FROM ubuntu:20.04
ENV LANG=C.UTF-8 LC_ALL=C.UTF-8

WORKDIR /data/source/
RUN apt-get update --fix-missing
ARG DEBIAN_FRONTEND=noninteractive 
RUN apt-get install -y wget pkg-config build-essential wget vim curl
RUN wget https://go.dev/dl/go1.18.2.linux-amd64.tar.gz
RUN tar xzvf go1.18.2.linux-amd64.tar.gz
ENV GOROOT=/data/source/go
ENV PATH=$PATH:$GOROOT/bin

RUN mkdir -p /app
WORKDIR /app

COPY *.go ./
COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go build -o go-was

ENTRYPOINT ["/app/go-was"]
