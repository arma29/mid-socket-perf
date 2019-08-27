# Official golang as parent image
FROM golang:1.8

# Set working directory to /go/src/app
WORKDIR /go/src/app

# Copy everything to container (only need dockerfile)
COPY . .

# Installing requirements 
RUN go get github.com/arma29/mid-socket-perf/socket/...

# # Execute the binary when the container launches
# CMD ["ServerUDP"]