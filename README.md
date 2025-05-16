# TaskAlley

## Running the service locally

---

## Problem statement

Build a simple task management microservice using Go

---

## Inferred requirements

1. User should be able to create, update, delete and list tasks  
2. User should be able filter task list be certain criteria  
3. Each task should be part of a project  
4. The service should be horizontally scalable  
5. The service should provide a mechanism to authorize user reaquests in a maintainable manner  
6. The service should have mechanism/s to communicate with other services  

---

## Overview

The service is build as a REST microservice which exposes http APIs for managing projects and tasks.

---

## The concept of Project allows us to

1. Group a set tasks together for easier management  
2. Easily extend to a state where a group of users can participate in a set of tasks  
3. Execute IAM checks on project level  

Projects and Tasks are highly cohesive entities hence it makes sense to keep them in one microservice. Other services in the system could be Identity and Access Management (IAM) service, Notification Service, Analytics service etc.

---

## Database

The service uses a simple Postgres database to persist Projects and Tasks, schema is present in `sql/schema.sql` file

---

## API spec

rendered at https://sumityadav29.github.io/taskalley/

---

## Communication with other services

Other services in the system might be interested in the events that happen in this microservice like Task is created, Task is deleted, Project is created etc. Moreover, few cross cutting concerns in the service might also be interested in these events like audit trail logging, etc.

However, the core logic of the service should be independent of all of this for the purpose of easy extensability and maintainability. For this reason the service provides and extendable, mantainable pattern - application events.

Application Events are emitted by the service on key changes TaskCreated, ProjectCreated, etc to be consumed withing the service. There can be multiple independent, parallel consumers of these events and they can take care of one thing each like sending a message to kafka topic for other microservices, noting audit trail, etc. 

This design ensures core logic and cross cutting concerns stay independent of each other leading to easier maintainability of the service. An example of KafkaApplicationEventHandler is present the code

---

## Authentication/Authorization

Service uses middleware pattern for authentication/authorization. This again ensure the core business logic is independent of it. This middleware can talk to IAM service, do local JWT checks or enforce auth through cached auth rules depending on use case and the core logic will remain unaffected.

The middleware pattern can also be extended for logging and cross cutting concerns

---

## Scaling horizontally

Since the service is stateless it can be scaled horizontally by adding more replicas, however, this should be done only when service is the bottleneck for throughput. If bottleneck is at the database level then adding more service replicas will be detrimnental instead of helpful.

**If database is bottleneck:**

1. Identify what exactly leads to low throughput  
2. Add correct indexes  
3. Remove soft deleted/unused data  
4. Consider adding a cache layer to reduce db load  
5. Consider read replicas to distribute read load  
6. Consider sharding database - using projectId as the shard key, this ensures all data related to particular project is stored in one node and reduces cross shard queries  

---

## Directory Structure

The service divides core fucntionality files into groups by domain like task and projects and cross cutting concern files by concerns like applicationevents and middlewares. Rest of the diectory structure ensures Go best practices are followed for understandability like entryppoints under `cmd`.

### Annotated directory structure
```
.
├── cmd                         # Service entry points, can be extended to consumer application, command line application, etc without affecting other parts
│   └── server                  # Web service entry point
│       └── main.go
├── config
│   └── database.go             # Contains application environment confiuration code
├── Dockerfile
├── go.mod
├── go.sum
├── internal                    # Contains core logic of the service, ensured no one can use outside by compiler
│   ├── applicationevents       # package for application event processing
│   │   ├── eventbus.go         # Defines EventBus, to where all Application Events are published
│   │   ├── eventhandler.go     # Defines interface for handling application events
│   │   ├── events.go           # Defineds events emitted by service
│   │   └── kafkahandler.go     # Concrete event handler which publishes to kafka topic for other microservices
│   ├── middlewares             # package for defininf middlewares
│   │   └── auth.go             # Auth middleware
│   ├── project                 # project domain files
│   │   ├── handler.go
│   │   ├── model.go
│   │   ├── repository.go
│   │   └── service.go
│   └── task                    # task domain files
│       ├── handler.go
│       ├── model.go
│       ├── repository.go
│       ├── service.go
│       └── taskfilters
│           └── taskfilters.go
├── README.md
├── sql                         # database schema and mock data
│   ├── mockdata.sql
│   └── schema.sql
└── static                      # docs
    └── docs
        ├── index.html
        ├── openapi.yaml
        └── redoc-static.html
```