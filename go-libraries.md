# Go Libraries

- [Go Libraries](#go-libraries)
      - [logr](#logr)


#### logr
[Ref v1.3.0 ](https://pkg.go.dev/github.com/go-logr/logr@v1.3.0); 
check out the [logr example folder](./examples-libraries/logr/).

Go’s log package doesn’t have leveled logs, you have to manually add prefixes like debug, 
info, warn, and error, yourself. Also, Go’s logger type doesn’t have a way to turn these 
various levels on or off on a per package basis [source](https://dave.cheney.net/2015/11/05/lets-talk-about-logging).

Check out the [usage](https://pkg.go.dev/github.com/go-logr/logr@v1.3.0#hdr-Usage).
**Functions**
```go
// With Go's standard log package, we might write:
log.Printf("setting target value %s", targetValue)

// With logr's structured logging, we'd write:
logger.Info("setting target", "value", targetValue)

// Errors are much the same. Instead of:
log.Printf("failed to open the pod bay door for user %s: %v", user, err)

// We'd write:
logger.Error(err, "failed to open the pod bay door", "user", user)
```
