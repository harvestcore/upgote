# Architecture

The architecture of _Upgote_ is an microservice based one. The main reason for choosing this type of architecture is the nature of the tasks to be performed. A module of the software, called Updaters, will obtain data in the background from a specified source from time to time. Each time this action is completed, other processes, also asynchronous in turn, will be launched. This software is meant to be used integrated with other systems.

- **Core**: Main core of the software. It contains all the logic related to Updater creation and data storing.
- **Updater**: Background process that fetches the configured data and will store it as soon as it performs an update.
- **API**: Landing place for all user requests. Handles all user requests to the system. The [framework](https://github.com/gorilla/mux) used is `Gorilla Mux`.
- **DB**: It is the place where all the fetched data is stored. Since the data schema is unknown until it is defined by the user, the database management system is a non relational one ([MongoDB](https://www.mongodb.com/) to be precise). The driver used to connect to the Mongo server is `mongo-driver` ([this one](https://godoc.org/go.mongodb.org/mongo-driver)).
- **Log**: Logs all operations performed across all the system. All these functionalities are handled using the `log` [Go package](https://golang.org/pkg/log/).
