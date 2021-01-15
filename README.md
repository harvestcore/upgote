# HarvestCCode

[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0) ![CI](https://github.com/harvestcore/HarvestCCode/workflows/Tests/badge.svg) ![Lint](https://github.com/harvestcore/HarvestCCode/workflows/Go%20linter/badge.svg) [![Build Status](https://travis-ci.com/harvestcore/HarvestCCode.svg?branch=master)](https://travis-ci.com/harvestcore/HarvestCCode)

---

## What is it

This software aims to be a cloud based place in which you can define a schema, a source and a update interval, allowing the user to store, fetch and update all its precious data.

Since the interaction with this software will be made via a REST-API, the **user** of this software must have the basic knowledge of this kind of architecture.

After this brief introduction you can:

- Learn more about the event workflows and user journeys [here](doc/architecture-workflows.md).
- Check the [project roadmap](doc/roadmap.md).
- Check its [architecture](doc/architecture.md)

---

## Why

The main reason is that I wanted to learn Go, use queues, background processes and event based architectures.

---

## Tools, services and others

- Main programming language: [Go](https://golang.org/)
- REST-API framework: [Gorilla Mux](https://github.com/gorilla/mux)
- DBMS: [MongoDB](https://www.mongodb.com/) with the [`mongo-driver`](https://godoc.org/go.mongodb.org/mongo-driver) for Go.
- Logs: [`log` Go package](https://golang.org/pkg/log/)
- Assertions: [`testify/assert` package](https://godoc.org/github.com/stretchr/testify/assert)

---

## Project structure

```yml
lenguaje: Go
fichero_tareas: Makefile
entidad:
  - src/config/config_manager.go
  - src/core/core.go
  - src/db/item.go
  - src/db/mongo_engine.go
  - src/db/responses.go
  - src/event/event.go
  - src/updater/updater.go
  - src/log/log.go
  - src/utils/eventtype.go
test:
  - src/config/config_manager_test.go
  - src/db/item_test.go
  - src/db/mongo_engine_test.go
  - src/event/event_test.go
  - src/updater/updater_test.go
  - src/log/log_test.go

```

---

## Documentation

- [Basic repo, git and GitHub configuration](doc/milestones/basic-git-github-config.md)
- [Currently open and closed user stories](https://github.com/harvestcore/HarvestCCode/issues?q=is%3Aopen+is%3Aclosed+is%3Aissue+label%3Auser-stories+)

## Other documentation

- [Task manager](doc/task_manager.md)
- [Assertion package](doc/assertion_library.md)
- [Test framework](doc/test_framework.md)
- [Progress on Milestone 2](doc/milestones/tests-hu.md)
- [Dockerfile, DockerHub and GH Registry](doc/dockerf.tests.md)
- [Progress on Milestone 3](doc/milestones/m3.md)
- [Continuous integration](doc/ci.md)
- [Progress on Milestone 4](doc/milestones/m4.md)
