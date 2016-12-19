# easel - A remote image processing server

## Aims & Background

In our datacenter, there are two kinds of machines. Machines **with GPU**, and those **without GPU**.

This library lets machines w/o GPU utilize remote GPU power via [grpc](http://www.grpc.io/) protocol.

## Sketch

![design sketch](./sketch.jpg)

 - [Server](./server) must be installed to machines **with** GPU.  
 It accepts GLSL shaders, uniform variables (including textures), and vertexes, then render an image and send it back to clients.
 - [Client CLI](./client-cli) is a CLI version of client implementation.
 - [Client Daemon](./client-daemon) is a daemon version of client implementation. It monitors a [SQL table](https://github.com/ledyba/easel/blob/master/client-daemon/db.sql) and process images according to table entries. This daemon could be useful for PHP front-ends.

## Implementation

[/easel-server](./easel-server) the entrypoint of "server" executable.  
[/easel-client-cli](./easel-client-cli): the entrypoint of "client-cli" executable.  
[/easel-client-daemon](./easel-client-daemon): the entrypoint of "client-daemon" executable.

[/image-filters](./image-filters): GLSL shaders. Currently, there is only a [lanczos10 filter](https://github.com/ledyba/easel/blob/master/image-filters/lanczos.go).  
[/server-impl](./server-impl):  
[/proto](./proto): grpc definitions.  
[/util](./util): utility functions.

[/test-images](./test-images): image materials for testing.  

## Known Issues

 - [lanczos10 filter](https://github.com/ledyba/easel/blob/master/image-filters/lanczos.go) cannot handles alpha channel correctly.

## How to run

### Common Prerequirements

 - [golang](https://golang.org/): Use latest version.

### Server

#### Prerequirements

```bash
# GLFW
sudo apt install libglfw3-dev
#
go get -u "github.com/go-gl/gl/v4.1-core/gl"
go get -u "github.com/go-gl/glfw/v3.2/glfw"
go get -u "github.com/Sirupsen/logrus"
go get -u "github.com/golang/protobuf/protoc-gen-go"
go get -u "google.golang.org/grpc"
go get -u "github.com/chai2010/webp"
go get -u "github.com/go-sql-driver/mysql"
```

#### Create Certificates

```bash
# create self-signed ca
openssl req -new -x509 -days 3650 -newkey rsa:2048 -nodes -keyout ca.key -out ca.crt
# create client key
openssl req -new -newkey rsa:2048 -nodes -keyout client.key -out client.csr
# sign
openssl x509 -req -days 365 -in client.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out client.crt
```

#### Command line flags
```
% ./bin/server -h
Usage of .bin/easel-server:
  -help
    	Print help and exit
  -listen string
    	listen addr (default ":3000")
```

### Client Daemon

#### Command line flags

```
% ./client-daemon -h
Usage of ./client-daemon:
  -db string
    	db address (default "user:password@tcp(host:port)/dbname")
  -filter string
    	applied filter name. (default "lanczos")
  -filter_lobes int
    	lobes parameter (default 10)
  -help
    	Print help and exit
  -server string
    	server to connect (default "localhost:3000")
  -workers int
    	workers to run (default 10)
```

### Client CLI

#### Command line flags

```
% ./bin/client-cli -h
Usage of ./bin/client-cli:
  -filter string
    	applied filter name. (default "lanczos")
  -filter_lobes int
    	lobes parameter (default 10)
  -help
    	Print help and exit
  -quality float
    	quality (default 95)
  -scale float
    	scale (default 2)
  -server string
    	server to connect (default "localhost:3000")
```
