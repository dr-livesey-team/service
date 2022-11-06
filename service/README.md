# Overview #

We had to solve several problems in order to implement the service. We had to search anomalies, locate they addresses
stores and manages necessary data in convinient way. Also we had to provide API to interact with service. So that, we
has been divided our service into small parts, such as address registry, api gateway, request regiustry and anomalies
searcher. All parts will be described in details futher.

# Detailed description #

> [Address registry](address_registry/README.md)

> [API Gateway](gateway/README.md)

> [Request registry](request_registry/README.md)

# Deployment #

## Docker ##

We build and run service in docker container. It is convinient for developers because docker install necessary
dependences and set up workspace (i.e. workspace layout, configuration files, etc) in appropriate way. Also service
becames cross-platform.

## Docker usage ##

To build docker container run
```
$ docker/client/unix/build.sh
```

To start docker container run
```
$ docker/client/unix/start.sh
```

To stop docker container run
```
$ docker/client/unix/stop.sh
```

To enter in docker container run
```
$ docker/client/unix/enter.sh
```

## Control utility ##

To set up workspace run
```
$ ./control.sh init 
```

To build service service run
```
$ ./control.sh build
```

To clean up workspace run
```
$ ./control.sh clean
```

To start service work
```
$ ./control.sh start
```

To stop service work
```
$ ./control.sh start
```
