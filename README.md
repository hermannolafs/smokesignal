# Smokesignal

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
