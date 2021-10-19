# param

Param is a wrapper of the standard `flag` package which adds a support of environment variables.

It is very simular to `flag`, but expects yet another parameter (environment variable name).

Example:

```go
// flag
myParam := flag.String("my-param", "default value", "Description")

// param
myParam := param.String("MY_PARAM", "my-param", "default value", "Description")
```

It support `Parse` and `PrintDefaults` functions too.

Available types: `bool`, `int`, `int64`, `uint`, `uint64`, `float64`, `time.Duration`, `string`.
