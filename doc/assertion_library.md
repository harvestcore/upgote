# Assertion library

The assertion library I'm using for this project is the `stretchr/testify/assert` [package](https://godoc.org/github.com/stretchr/testify/assert).

There are other packages that also provide assertion functions, like [github.com/realistschuckle/testify/assert](https://godoc.org/github.com/realistschuckle/testify/assert), and also Go has some mecanisms to perform this type of assertion testing.

After some investigation I've chosen the `stretchr/testify/assert` package due to the fact that it has:

- More than 130 assertion types.
- More than 140 functions to raise those assertions.
- Capability to perform HTTP requests, which is really useful to test the REST-API.
- Common function naming, making it understandable since it is really similar to other programming languages.

Apart from that:

- It does not require any special setup, just the package import.
- Its documentation is clear and contains a wide range of examples.
- It can be used with other packages from [`stretchr/testify`](https://github.com/stretchr/testify), like [`mock`](https://godoc.org/github.com/stretchr/testify/mock), which can also be usefull when testing endpoints.

The only downside is that it is not an independant package and makes use of the `testing` [package](https://golang.org/pkg/testing/) from Go, but this is not necessarily a downside.

This is a test example (from the documentation):

```go
import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

  var a string = "Hello"
  var b string = "Hello"

  assert.Equal(t, a, b, "The two words should be the same.")

}
```

As mentioned before, the syntax is easy to understand and it makes creating tests and assertions real quick.
