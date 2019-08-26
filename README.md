# Middleware - Socket - Performance Evaluation

### Prerequisites 

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

### Run

You can now run the program by typing its full path at the command line:

```bash
$ $GOPATH/bin/bin_name
```

Or, as you have added `$GOPATH/bin` to your `PATH`, just type the binary name:

```bash
$ bin_name
```
