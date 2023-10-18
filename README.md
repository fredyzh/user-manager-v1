# user-manager-v1

## Overview
- a microservice user management implemented by golang
- generated grpc objects as DTO
- grpc gateway generate http reverse object

## Code Base Structure
```
.
├── api
|    └──meta   
│    │    └── v1
│    │         └──pb
|    └──policy
|    └──role
|    └──token
│    │   └── v1
│    │         └──pb
│    │         └──service
|    └──user
│        └── v1
│             └──pb
│             └──service
├── app  //
│   ├── cmd
│   └── config
│   |    └── v1
|   ├── internal
|   |       └── business
|   |       |      └── user
|   |       |          └── v1
|   |       |        
|   |       └── data
|   |       |      └── user
|   |       |          └── v1
|   |       |              └── mongo
|   |       |── server
|   |               └── v1
|   |       |── service
|   |               └── v1
|
|── client
|       └── go
|       └── java
|       └── nodejs
|       └── python
|── deploy
|       └── docker
|             └── bin
|── docs
├── third_party
|       └── google
|               └── api
|── web

```