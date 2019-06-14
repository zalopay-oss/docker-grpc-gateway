# Docker grpc-gateway

- [Docker grpc-gateway](#docker-grpc-gateway)
  - [Overview](#overview)
  - [Architecture Demo](#architecture-demo)
  - [Requirement](#requirement)
  - [Installation](#installation)
  - [Build](#build)
  - [Run](#run)
  - [Test](#test)
  - [Contribution](#contribution)
  - [Acknowledgments](#acknowledgments)

## Overview

Building grpc-gateway has caused many difficulties for developers such as the installation of the Golang programming environment, protobuf, how to build,.... There is a simpler, less time-consuming solution that uses `Docker to build grpc-gateway`.

In this project, I would like to introduce the solution to use docker and demo how to use it. For more details, see [here](https://medium.com/zalopay-engineering/docker-grpc-gateway-e2efbdcfe5c-e2efbdcfe5c).

## Architecture Demo

<p align="center">
  <img src="./images/model.png"/>
</p>

## Requirement

- Docker
  
## Installation

Install Docker:

- On [Ubuntu](https://docs.docker.com/install/linux/docker-ce/ubuntu/).
- On [MacOS](https://docs.docker.com/docker-for-mac/install/).
- On [Window](https://docs.docker.com/docker-for-windows/install/).
  
## Build

```sh
#clone project
$ git clone https://github.com/zalopay-oss/docker-grpc-gateway.git
```

## Run

- Run grpc-gateway
  
```sh
$ cd ./gateway/api/proto/gen/grpc-gateway

#docker build -t grpc-gateway directory_contains_dockerfile
$ docker build -t grpc-gateway .

#run image
$ docker run  --network host grpc-gateway
```

- Run service
  
```sh
$ cd ./service

#build image
$ docker build -t service-ping .

#run image
$ docker run  -- network host service-ping
```

## Test

- Test ServiceA
  
```sh
$ curl -X GET "http://localhost:9000/core/serviceA/ping/999999"
{"timestamp":"1560311912214","serviceName":"service A"}
```

- Test ServiceB

```sh
curl -X GET "http://localhost:9000/core/serviceB/ping/999999"
{"timestamp":"1560312187849","serviceName":"service B"}
```

## Contribution

If you find anything wrong or would like to contribute in any way, feel free to create a pull request/open an issue/send me a message. Any comments are welcome!

## Acknowledgments

- Thanks [Andy](https://github.com/anhldbk) for supporting me during the implementation project.
- Thanks [namely/docker-proto](https://github.com/namely/docker-protoc) for the open source is great.
