# Queue Implementation

This source code is implementation of `Queue` interface.

```go
// Queue interface.
type Queue interface {
    Push(key interface{})
    Pop() interface{}
    Contains(key interface{}) bool
    Len() int
    Keys() []interface{}
}
```

running test with coverage:

```sh
go test -cover
```

result:

```sh
PASS
coverage: 85.7% of statements
ok      _/Users/macbook/Desktop/msa-challenge/queue      0.006s
````
