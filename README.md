# HarvestCCode

[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0) ![Tests](https://github.com/harvestcore/HarvestCCode/workflows/Tests/badge.svg)

---

## What is it

This software aims to be a cloud based place in which you can define a schema, a source and a update interval, allowing the user to store, fetch and update all its precious data.

After this brief introduction you can:

- Learn more about the event workflows and user journeys [here](doc/architecture-workflows.md).
- Check the [project roadmap](doc/roadmap.md).

---

## Architecture

The architecture of HarvestCCode is an event based one. The main reason for choosing this type of architecture is the nature of the tasks to be performed. A module of the software, called Updaters, will obtain data in the background from a specified source from time to time. Each time this action is completed, other processes, also asynchronous in turn, will be launched.

It has different parts that interact with each other via events. Those parts are:

- **Core**: Main core of the software. It contains all the logic related to Updater creation and data storing.
- **Handler**: It is the central part that handles all the events sent and received by the rest of the parts.
- **Updater**: Background process that fetches the configured data and will send events as soon as it performs an update.
- **API**: Landing place for all user requests. Handles all user requests to the system. The [framework](https://github.com/gorilla/mux) used is `Gorilla Mux`.
- **DB**: It is the place where all the fetched data is stored. Since the data schema is unknown until it is defined by the user, the database management system is a non relational one ([MongoDB](https://www.mongodb.com/) to be precise). The driver used to connect to the Mongo server is `mongo-driver` ([this one](https://godoc.org/go.mongodb.org/mongo-driver)).
- **Log**: Logs all operations performed across all the system. All these functionalities are handled using the `log` [Go package](https://golang.org/pkg/log/).

Since the operations to be performed are mostly asynchronous it does not make sense to use a non-event architecture. In addition, these tasks are programmed to run in the background, so there is no need to provide more than a confirmation that its request has been processed correctly or not.

---

## Why

The main reason is that I wanted to learn Go, use queues, background processes and event based architectures.

---

## Tools, services and others

- Main programming language: [Go](https://golang.org/)
- REST-API framework: [Gorilla Mux](https://github.com/gorilla/mux)
- DBMS: [MongoDB](https://www.mongodb.com/) with the [`mongo-driver`](https://godoc.org/go.mongodb.org/mongo-driver) for Go.
- Logs: [`log` Go package](https://golang.org/pkg/log/)
- Communication between processes: [`rpc` Go package](https://golang.org/pkg/net/rpc/)

---

## Documentation

- [Basic repo, git and GitHub configuration](doc/milestones/basic-git-github-config.md)
- [Currently open and closed user stories](https://github.com/harvestcore/HarvestCCode/issues?q=is%3Aopen+is%3Aclosed+is%3Aissue+label%3Auser-stories+)
