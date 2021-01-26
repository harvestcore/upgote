# Microservices

## R1

### "System" requirements

This section contains a list of desirable (or not) features that a web framework should have in order to be used in HarvestCCode.

- No need for MVC support.
- Middleware support is welcome.
- Async/concurrency is also desirable.
- Ultra high performance & availability is desirable but not a must.
- Lightweight.
- Open source.
- Not obsolete.
- JSON support.
- Good documentation.
- Good unit-test integration.
- Scalability and Modularity.
- No ORM needed.

### Some well known frameworks and toolkits

#### [Martini](https://github.com/go-martini/martini)

This framework is inspired in Express and Sinatra, and its syntax is really similar. It allows exception handling, routing and also the possibility to include custom middlewares. It is also widely used to create quick, easy and simple API, since it supports natively both JSON and XML.

```go
m := martini.Classic()

m.Get("/hello", func(params martini.Params) string {
  return "Hello " + params["name"]
})

m.RunOnAddr(":8080")
```

#### [Gorilla](https://github.com/gorilla/mux)

Gorilla Mux implements a custom HTTP Handler that includes a lot more features than the default one from Go. It is super simple to use, and does not require any extra library or extra knowledge. Since it implements the http.Handler interface it works right away with the basic [HTTP Server](https://golang.org/pkg/net/http/) from Go, so there is no need to install extra packages.

Related to unit testing, since it makes use of the [HTTP Server](https://golang.org/pkg/net/http/) from Go it works smoothly with the [httptest](https://golang.org/pkg/net/http/httptest/) package.

Apart from Mux, Gorilla includes some other ["tools"](https://www.gorillatoolkit.org/) that are also really useful.

Code example:

```go
router := mux.NewRouter().PathPrefix("/api").Subrouter()

router.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
    file, _ := ioutil.ReadFile(logFilePath)

    w.Header().Set("Content-Type", "text/plain")
    w.Write(file)
}).Methods("GET")

server := &http.Server{
    Handler: router,
    Addr:    ":8080",

    // Read and write timeouts to avoid the server hang
    ReadTimeout:  10 * time.Second,
    WriteTimeout: 10 * time.Second,
}

server.ListenAndServe()
```

#### [Beego](https://beego.me/)

Beego is really similar to Django, and its features are more likely to be used in web applications than APIs. It includes MVC support and also includes an ORM. Due to these two points I discarded using Beego, not for being a bad choice, but for including features that I will not use and that will increase the size and complexity of the software.

#### [Gin](https://github.com/gin-gonic/gin)

It is a lightweight and minimalist framework that claims to be up to 40 times faster than Martini. It also makes use of the default http.HTTP package from Go, so it also runs every request on a different [Goroutine](https://golangbot.com/goroutines/).

Its downside is the same as Martini, both are the "Express" for Golang, and they're not meant to be used in complex APIs.

Code example:

```go
r := gin.Default()

r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "pong",
    })
})

r.Run()
```

#### [Revel](https://revel.github.io/)

This framework is similar to [Beego](https://beego.me/), since it is a full-stack web framework. It includes features like MVC support and custom modules integration. As usual it also allows to create custom middlewares. Unlike some others in this lists, due to its full-stack features for bigger web applications, I discarded this option.

### So, why Gorilla Mux?

There are other similar frameworks, like [Revel](https://revel.github.io/) or [Beego](https://beego.me/), but those discarded for being way larger than Gorilla Mux, and because they include some features that are not needed in this type of use, like MVC support.

Another important point is that HarvestCCode is not meant to be used intensively like it could happen with other APIs, so there is no need to use a fully featured framework.

As it only implements the router, the only thing left to do as a developer is to define the different handlers and middlewares needed.

Also, its "coding style" is more familiar and easier to understand (at least for me).

#### A note on performace

This [GitHub repo](https://github.com/smallnest/go-web-framework-benchmark) includes a bunch of performance tests to numerous Golang web frameworks. The tests include concurrency, cpu-bound and latency among others.

The results show that in general the fastest one is Gin, due to its lightweight design, and the slowest is Martini. But this only happens in scenarios where concurrency is not key. In these cases Gorilla turns to be way capable than the other ones, with less latency and a better request/second ratio.

The downside of Gorilla is that, in some cases, it allocates more memory than the others, but the difference is not relevant enough to discard Gorilla.

## R2

### API Structure and resources

> The endpoints documentation can be seen [HERE](endpoints/endpoints.md).

This software has four different resources that are transferred by its API. There are seven endpoints in total, whose URIs are shown in bold text:

- 1st resource. <u>Status and healtheck</u>. Managed by the [status handler](../api/handlers/status.go):
  - [GET] **/status**. Returns the status of the software.
    - Possible status codes: 200
  - [GET] **/healthcheck**. Returns a simplified status of the component.
    - Possible status codes: 200
- 2nd resource. <u>Data fetched by the updaters</u>. Managed by the [data handler](../api/handlers/data.go):
  - [POST, DELETE] **/data**. Data management.
    - Possible status codes: 200, 422
- 3rd resource. <u>Updaters</u>. Managed by the [updater handler](../api/handlers/updater.go):
  - [GET, POST, PUT, DELETE] **/updater**. Updaters management.
    - Possible status codes: 200, 201, 422
  - [POST] /updater/action. Start or stop an updater.
    - Possible status codes: 200, 422
- 4th resource. <u>Logs</u>. Managed by the [log handler](../api/handlers/log.go):
  - [GET, POST] **/log**. Log operations.
    - Possible status codes: 200, 422

The [API Server](../api/server.go) encapsulates the HTTP Server and the Gorilla/Mux router. There is only one instance of the server, and registers all the individual handlers or routers (the ones mentiones above).

#### User stories <-> Endpoints

Status endpoints:

- These endpoints were added as a way to check the status of the software in any given time. They're used in the [Dockerfile.hcc](../Dockerfile.hcc) and [docker-compose.yml](../docker-compose.yml) files.

Data endpoints:

- [HU#13 - Fetch data from the system](https://github.com/harvestcore/HarvestCCode/issues/13)

Updater endpoints:

- [HU#12 - Manage Updaters](https://github.com/harvestcore/HarvestCCode/issues/12)

Log endpoints:

- [HU#14 - Logs](https://github.com/harvestcore/HarvestCCode/issues/14)
- [HU#18 - Logs download](https://github.com/harvestcore/HarvestCCode/issues/18)

## R3

### Logs

In order to include the logging systen into the API I have created a custom middleware that integrates with Gorilla Mux. The code is [over here](../api/middlewares/logging.go) and makes use of the current logging infrastructure to log every single request that the router handles.

The implementation only requieres to define a function.

```go
// logging.go - Define middleware.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.AddRequest(r)
		next.ServeHTTP(w, r)
	})
}

// server.go - Use the middleware in the router.
router.Use(LoggingMiddleware)
```

### Distributed configuration

[Etcd3](https://github.com/etcd-io/etcd) is the tool used for distributed configuration. The [config manager](../config/config_manager.go) includes an instance of the client. The flow that the manager follows when a variable is fetched is the following:

1. Check if the variable is in the variables pool. If so, return the variable.
2. If not, try to fetch the variable from the Etcd3 server via the client. If existant, store it in the variables pool and return it.
3. If not, try to fetch it from the environment. If existant, store it in the variables pool and return it.
4. If not, use the default value and store it in the variable pool.

## R4

Each one of the handlers listed in [R2](#R2) and also the server has its own `*_test.go` file with its corresponding unit tests. The tests are design to check all the possible scenarios when performing requests to the API. They check all the error codes, structures that are returned by the endpoints and also try to perform requests by sending wrong data or parameters. More info in each of the files listed below.

- [log_api_test.go](../api/tests/log_api_test.go)
- [data_api_test.go](../api/tests/data_api_test.go)
- [status_api_test.go](../api/tests/status_api_test.go)
- [updater_api_test.go](../api/tests/updater_api_test.go)
- [server_test.go](../api/server_test.go)

Also, as mentioned [here](#User stories <-> Endpoints), each handler is related to one (or more) user story. The tests were designed in order to make sure that the operations listed in the US are always correct and all the possible errors are handled.

## R5

Some other extra work:

- [Endpoints documentation](endpoints/endpoints.md).

- [Dockerfile.hcc](../Dockerfile.hcc). I've created this Dockerfile in order to be able to run the software within a Docker container. I've also followed (or at least tried to follow) the best-practices when creating Docker images. The result is an image that only weights 30.5MB.
- [docker-compose.yml](../docker-compose.yml). Docker-compose file to run both the software and a MongoDB container orchestrated.

- [publish-hcc-image.yml](../.github/workflows/publish-hcc-image.yml). Extra pipeline to deploy the built image to DockerHub and GitHub Registry.