# Smokesignal

<p align="center">
<a href="https://godoc.org/github.com/hermannolafs/smokesignal"><img src="https://godoc.org/github.com/hermannolafs/smokesignal?status.svg" alt="Godoc" /></a>
<a href="https://goreportcard.com/report/github.com/hermannolafs/smokesignal"><img src="https://goreportcard.com/badge/github.com/hermannolafs/smokesignal" alt="Go Report Card" /></a>
</p>

A simple smoke testing library for Golang.

Currently the following things can be tested:
- Assert if port is used or not
- Assert GET request to endpoint returns HTTP OK
- All combinations of the two preceding points

Your server needs to fulfill the interface `smokesignal.Server`
For examples see the [example package](example/interface_example_test.go)
The plan is to add support for binaries so that smokesignal smoke tests can be run after building the binary, another common smoke test.

```go
type Server interface {
    Run(quit chan os.Signal)
    Stop(ctx context.Context) error
}
```

Best practice is to tag your smoke tests and run then seperately before the rest of your tests

Heavily inspired by [steinfletcher/apitest](https://github.com/steinfletcher/apitest)