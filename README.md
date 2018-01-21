htprox
======

A TCP-over-HTTP proxy for when only outbound connections are allowed.

`htprox` can act as a "server", a "gateway" and a "client".

## Gateway
```
htprox -gateway -local 0.0.0.0:80
```
Will start an htprox gateway server on port 80.
This will handle connection requests from "clients" and communicate with "servers"

### Endpoints
* `/register?name=<name>` - Register `<name>` as an endpoint
* `/release?name=<name>` - Release endpoint
* `/open?name=<name>` - Start a new session on endpoint `<name>`. Will return a new ID on success
* `/close?name=<name>&id=<id>` - Close a session
* `/send?name=<name>&id=<id>` - PUT data to this endpoint to send data to that session buffer
* `/recv?name=<name>&id=<id>` - GET data from this session buffer
* `/poll?name=<name>` - Check if any sessions have non-empty buffers
* `/status` - View overall status of the gateway


## Server
```
htprox -server -local :22 -remote gateway-name:80 -endpoint ssh
```
This will start an htprox server session and register the name "ssh" with the gateway.

It will periodically (default every 15s) check in with the gateway server for incomming connections.
When a new connection is found, it will connect to `localhost:22` and relay traffic through.

## Client
```
htprox -client -local :2222 -remote gateway-name:80 -endpoint ssh
```
This will start the htprox client session. The client will begin listening on `:2222` for connections, 
and when it receives a connection, will relay that traffic to the gateway.
