## Publisher/Subscriber Service

#### Team members: 
* [Chi Zhang](mailto:zcdirk@stanford.edu)
* [Xiao Xin](mailto:xxin@stanford.edu)

### Introduction
Publisher/Subscriber(Pub/Sub) is an asynchronous messaging-oriented service
that serves as the backbone of contemporary distributed systems. In many
cases it is used to decouple the functions of a massive monolithic application
by aggregating multiple smaller yet more cohesive services. Being such a critical 
component to distributed systems, Pub/Sub services should also be scalable
so it could fit the requirement of the services that depend on it. 

Therefore, we propose this project that aims to replicate the functions of
a pub/sub service with the capabilities to horizontally scale onto multiple
machines to better serve the distributed systems it supports. 

### Install
```
go get -u github.com/cs244b-2020-spring-pubsub/pubsub
```

### Usage
```
# Usage
pubsub --help

# Initiate the server
pubsub --config=<path_to_config>
```

### Config
Pub/Sub ingests protobuf based configuration; config structure is defined
at [proto/config.proto](proto/config.proto).

Users shall write their config as a `textproto` file. You can read more
about `textproto` from [here](https://medium.com/@nathantnorth/protocol-buffers-text-format-14e0584f70a5).
