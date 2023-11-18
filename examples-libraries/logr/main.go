package main

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/go-logr/logr/funcr"
)

func ExampleNew(loggerName string) logr.Logger {
	// funcr.New expects a function with the signature:
	// func(prefix, args string)
	// funcr.Options is a struct which is defined in funcr.go
	var log logr.Logger = funcr.New(func(prefix, args string) {
		fmt.Println(prefix, args)
	}, funcr.Options{})

	log = log.WithName(loggerName)
	return log
}

func ExampleNewJSON(loggerName string) logr.Logger {
	var log logr.Logger = funcr.NewJSON(func(obj string) {
		fmt.Println(obj)
	}, funcr.Options{})

	log = log.WithName(loggerName)
	return log
}

func ExampleOptions(loggerName string, level int) logr.Logger {
	var log logr.Logger = funcr.NewJSON(
		func(obj string) { fmt.Println(obj) },
		funcr.Options{
			Verbosity: level,
		})
	log = log.WithName(loggerName)
	return log
}

func main() {
	log := ExampleNew("AppLogger")
	log.Info("INFO1", "key1", "value1", "key2", "value2")

	log_json := ExampleNewJSON("AppLogger")
	log_json.Info("INFO1", "key1", "value1", "key2", "value2")

	log_options := ExampleOptions("AppLogger", 1)
	log_options.V(0).Info("V(0) message", "key", "value")
	log_options.V(1).Info("V(1) message", "key", "value")
	log_options.V(2).Info("V(2) message", "key", "value")
}
