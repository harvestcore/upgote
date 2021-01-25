# Microservices

For this project I've used [Gorilla Mux](https://github.com/gorilla/mux) to create a REST API to manage the system.

## Why Gorilla Mux?

Gorilla Mux implements a custom HTTP Handler that includes a lot more features than the default one from Go. It is super simple to use, and does not require any extra library or extra knowledge. Since it implements the http.Handler interface it works right away with the basic [HTTP Server](https://golang.org/pkg/net/http/) from Go, so there is no need to install extra packages.

There are other similar frameworks, like [Revel](https://revel.github.io/) or [Beego](https://beego.me/), but those discarded for being way larger than Gorilla Mux, and because they include some features that are not needed in this type of use, like MVC support.

Another important point is that HarvestCCode is not meant to be used intensively like it could happen with other APIs, so there is no need to use a fully featured framework.

Related to unit testing, since it makes use of the [HTTP Server](https://golang.org/pkg/net/http/) from Go it works smoothly with the [httptest](https://golang.org/pkg/net/http/httptest/) package.

## API Structure

> The endpoints documentation can be seen [HERE](endpoints/endpoints.md).

There are seven endpoints in total:

- Managed by the [status handler](../api/handlers/status.go):
  - [GET] /status. Returns the status of the software.
  - [GET] /healthcheck. Returns a simplified status of the component.

- Managed by the [log handler](../api/handlers/log.go):
  - [GET, POST] /log. Log operations.

- Managed by the [data handler](../api/handlers/data.go):
  - [POST, DELETE] /data. Data management.

- Managed by the [updater handler](../api/handlers/updater.go):
  - [GET, POST, PUT, DELETE] /updater. Updaters management.
  - [POST] /updater/start. Start an updater.
  - [POST] /updater/stop. Stop an updater.

The [API Server](../api/server.go) encapsulates the HTTP Server and the Gorilla/Mux router. There is only one instance of the server, and registers all the individual handlers or routers (the ones mentiones above).

Each one of the handlers and also the server has its own `*_test.go` file with its corresponding unit tests.

- [log_api_test.go](../api/tests/log_api_test.go)
- [data_api_test.go](../api/tests/data_api_test.go)
- [status_api_test.go](../api/tests/status_api_test.go)
- [updater_api_test.go](../api/tests/updater_api_test.go)
- [server_test.go](../api/server_test.go)

## User stories <-> Endpoints

Log endpoints:

- [HU#14 - Logs](https://github.com/harvestcore/HarvestCCode/issues/14)
- [HU#18 - Logs download](https://github.com/harvestcore/HarvestCCode/issues/18)

Data endpoints:

- [HU#13 - Fetch data from the system](https://github.com/harvestcore/HarvestCCode/issues/13)

Updater endpoints:

- [HU#12 - Manage Updaters](https://github.com/harvestcore/HarvestCCode/issues/12)

Status endpoints:

- These endpoints were added as a way to check the status of the software in any given time. They're used in the [Dockerfile.hcc](../Dockerfile.hcc) and [docker-compose.yml](../docker-compose.yml) files.
