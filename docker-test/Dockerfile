# Official golang as parent image
FROM golang:1.8

# Mantainer
LABEL maintainer=arma29-github

# Set working directory to /go/src/app
WORKDIR /go/src/app

# Installing requirements 
RUN go get github.com/arma29/mid-socket-perf/socket/...
