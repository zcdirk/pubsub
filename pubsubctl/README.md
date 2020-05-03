## Publisher/Subscriber Service Client

### Intro
A primitive command-line interface with main service.

### Install
```
go get -u github.com/cs244b-2020-spring-pubsub/pubsub/pubsubctl
```

### Usage
```
# Usage
pubsubctl help

# To publish a message
pubsubctl publish --topic=<topic> message

# To subscribe to topics
pubsubctl subscribe <topic1> <topic2> ...
```
