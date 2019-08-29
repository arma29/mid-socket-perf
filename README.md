# Middleware - Socket - Performance Evaluation

A simple socket project to evaluate the UDP and TCP performance. The client request a number and the server returns the respective fibonacci element. UDP and TCP implementations built in Go. 

You can either check/run the code or run the tests to generate .csv files into your machine. Feel free to analyse the data any way you want.

## To run the code

### Prerequisites 

- Go installed.

From [Golang - Getting Started](https://golang.org/doc/install)

Add /usr/local/go/bin to the PATH environment variable. You can do this by adding this line to your /etc/profile (for a system-wide installation) or $HOME/.profile:

```bash
$ export PATH=$PATH:/usr/local/go/bin
```

For convenience, add the workspace's bin subdirectory to your `PATH`:

```bash
$ export PATH=$PATH:$(go env GOPATH)/bin
```

Also a good idea add these lines:

```bash
$ export GOPATH=$(go env GOPATH);
$ export PATH=$PATH:$GOPATH/bin;
```

### Install via `go get`

```bash
$ go get github.com/arma29/mid-socket-perf/socket/...
```

This command will build and install all go programs under `mid-socket-perf/socket` directory. It then installs that binaries to the workspace's bin directory as the folder name.

### Usage

Now, you can run the program by typing its full path at the command line:

```bash
$ $GOPATH/bin/bin_name
```

Or, as you have added `$GOPATH/bin` to your `PATH`, just type the binary name:

```bash
$ bin_name
```
In this case, we need to start the server, and his respective client. Each server will listening to all net interfaces. 

- UDP Server listening port: 1200
    - `$ $GOPATH/bin/ServerUDP`
- TCP Server listening port: 7171
    - `$ $GOPATH/bin/ServerTCP`

To run the clients, we need to pass the server IP and the request number as argument e.g:

```bash
$ $GOPATH/bin/ClientTCP 127.0.0.1 5
$ $GOPATH/bin/ClientUDP 127.0.0.1 5
```

This will print 10K lines containing the fibonacci number requested, sample number and the measured time between the request and response, in milliseconds.

## To run the performance tests

### Prerequisites

- Docker
- docker-compose

### Usage

Simply run:

```bash
$ ./suit.sh
```

It takes 2-3 hours to run everything. Briefly, this will test the scenarios [UDP,TCP](socket type) x [1,2,3,4,5](simultaneous clients) and will output in `/home/$USER/Output` the .csv files in this format:

| Fibonacci | Sample | Time |
| -------- | -------- | -------- |
| requested number     | Sample number     | RTT time     |
