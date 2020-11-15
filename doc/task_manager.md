# Task manager

The task manager I'm using is [`make`](https://www.gnu.org/software/make/). There are some reasons why `make` is the *goto* task manager to use:

- It is easy to use it, since the declaration of the rules is quite simple and easy to understand.
- I've used it in the past in other personal projects.
- It does not require any extra configuration file, just a Makefile.

There are other alternatives, like the ones below:

- [Gofer](https://github.com/chuckpreslar/gofer)
- [Gogradle](https://github.com/gogradle/gogradle), which is Gradle with a plugin for Go.
- [xmake](https://xmake.io/#/)
- [Cmake](https://cmake.org/)

Since `make` is more like a generic tool it can be also used for other tasks, not only related to the programming language ones. For example it can be used to generate documentation.

The `makefile` I've made is the following and can be found [here](../Makefile). It contains the four tasks that are needed for now:

- `all`, install all the dependencies and builds the software.
- `run`, run the software.
- `test`, run all the unit tests.
- `deps`, add dependencies and install them.
- `install`, compile and install dependencies.
- `build`, builds the software, generating an executable file.
- `lint`, runs the linter across all Go files.
- `clean`, remove cache and object files.

```Makefile
all: install

run:
    go run src/main.go

test:
    go test ./src...

deps:
    go get ./src...

install:
    go install ./src...

build:
    go build ./src...

lint:
    go vet ./src...

clean:
    go clean ./src...
```
