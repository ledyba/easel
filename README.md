# easel - A remote image processing server

## Aims & Background

In our datacenter, there are two kinds of machines. Machines with GPU, and those without GPU.

This library lets machines w/o GPU utilize GPU power via [grpc](http://www.grpc.io/) protocol.

## Sketch

 - [Server](./server) will be installed to machines **with** GPU.  
 It accepts GLSL shaders, uniform variables (including textures), and vertexes and execute it.
 - [Client CLI](./client-cli) is a client implementation
 - [Client Daemon](./client-daemon)

## Implementation

 - [/client-cli](./client-cli): the entrypoint of "client-cli" executable.
 - [/client-daemon](./client-daemon): the entrypoint of "client-daemon" executable.
 - [/image-filters](./image-filters): GLSL shaders. Currently, there is a [lanczos10 filter](https://github.com/ledyba/easel/blob/master/image-filters/lanczos.go) only.
 - [/proto](./proto):

## Known Issues

 - [lanczos10 filter](https://github.com/ledyba/easel/blob/master/image-filters/lanczos.go) cannot handles alpha channel correctly.

## How to run

### Prerequirement

 - [golang](https://golang.org/): Use latest version.
 - GLFW (for server)
   - `sudo apt install libglfw3-dev`

### Server

### Client
