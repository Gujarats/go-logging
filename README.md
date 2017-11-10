## Golang logging library

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/op/go-logging) [![build](https://img.shields.io/travis/op/go-logging.svg?style=flat)](https://travis-ci.org/op/go-logging)

Please see the original readme link [here ](https://github.com/op/go-logging/blob/master/README.md).

## Modification

I hate to initialize the log file for debugging purpose. so I create some method and you can just call it like this : 
```go
// this will print the debug message with the prefix and the message with spefied format
logging.Debug("prefix", "message")

```

I also save the debug message in a file called `logger.log` in the directory where you called this function.

