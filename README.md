# _Upgote_

|                        | Status                                                       |
| ---------------------- | ------------------------------------------------------------ |
| CI                     | [![Unit tests](https://github.com/harvestcore/upgote/actions/workflows/tests.yml/badge.svg)](https://github.com/harvestcore/upgote/actions/workflows/tests.yml) |
| Lint                     | [![Unit tests](https://github.com/harvestcore/upgote/actions/workflows/golint.yml/badge.svg)](https://github.com/harvestcore/upgote/actions/workflows/golint.yml) |
| Release                | [![Release](https://img.shields.io/github/v/release/harvestcore/upgote)](https://github.com/harvestcore/upgote/releases) |
| License                | [![License](https://img.shields.io/github/license/harvestcore/upgote)](https://www.gnu.org/licenses/gpl-3.0) |
| API reference          | [![API](https://img.shields.io/badge/-API-informational)](./doc/endpoints/endpoints.md) |

---

## What is it

This software aims to be a cloud based place in which you can define a schema, a source and a update interval; allowing the user to store, fetch and update all its precious data.

Since the interaction with this software will be made via a _REST API_, the **user** of this software must have the basic knowledge of this kind of architecture.

After this brief introduction you can:

- Learn more about the workflows and user journeys [here](doc/architecture-workflows.md).
- Check the [project roadmap](doc/roadmap.md).
- Check its [architecture](doc/architecture.md)
---

## Tools, services and others

- Main programming language: [Go](https://golang.org/)
- REST-API framework: [Gorilla Mux](https://github.com/gorilla/mux)
- DBMS: [MongoDB](https://www.mongodb.com/) with the [`mongo-driver`](https://godoc.org/go.mongodb.org/mongo-driver) for Go.
- Logs: [`log` Go package](https://golang.org/pkg/log/)
- Assertions: [`testify/assert` package](https://godoc.org/github.com/stretchr/testify/assert)

---

## Installing _Upgote_

> Before installing and running _Upgote_ please take a look at the [environment variables](doc/envvars.md) needed.

This software can be installed and executed in multiple ways:

- By downloading this repository and running:
  - For local testing:
    - `make deps && make run`
  - For local production:
    - `make build && make start`
- Using the available Docker image:
  - Building it using the `Dockerfile` file.
  - Pulling it either from [ghcr.io](https://github.com/users/harvestcore/packages/container/package/upgote).
  - Using the `docker-compose.yml` file available:
    - `docker-compose build && docker-compose up`
