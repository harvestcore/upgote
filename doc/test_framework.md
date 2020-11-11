# Test framework

The testing framework I'm using is the native one from Go, which makes use of the `go test` command along with the `testing` package from Go. The main reason for this is that no additional software or extra packages are needed so it is only a matter of writting the tests.

There are other [frameworks](https://bmuschko.com/blog/go-testing-frameworks/), and even some of them include web UI, but that's not necessary at all, since a simple output in a terminal is ok for me.

> Note: As mentioned [here](assertion_library), the assertion library used is [`stretchr/testify/assert`](https://godoc.org/github.com/stretchr/testify/assert).

The workflow to write and execute tests is:

- Create as many `*_test.go` files as you want. In my case one file for each one of the structs and functions that requiere tests.
- Create the tests.
- Run `go test`, or run the command configured in the task manager (in my case `make test`).
