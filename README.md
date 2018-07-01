# Go Windows ENV
This package can be used to load Environment variables from a .env file into your application's environment.  
This allows you to keep your env clean across multiple apps.  I haven't tested this on OSX or windows, but this should work on those Operating systems.

## Parse
Parse accepts a variadic string slice representing a slice of filepaths to be added to the ENV as it's parameter and returns a nillable error.
```go
Parse([]paths ...string) error {}
```

## Basic Usage
1. Download the package
```shell
go get github.com/mjblack23/win-env
```
2. Require it into your module
```go
import (
    env "github.com/mjblack23/win-env"
)
```
2. Create an .env file
```shell
API_KEY=some-secure-code
SECRET_KEY=an-even-more-secure-code
```
3. Call env.Parse() and party
```go
func main () {
    err := env.Parse([]string{}...)

    fmt.Println(os.Getenv("API_KEY")) // outputs some-secure-code
}
```

## Using a custom file(path)
You don't have to use .env if you don't want to.  You can set your own filepath inside of the string slice.

### Example.
```go
customEnvFile := []string{"./.development_env"}

err := env.Parse(customEnvFile...)
```

### Using multiple files to build your environment
go-win-env supports multiple files to build your environment if needed.

### Example
```go
baseEnv := "./base_env"
devEnv := "./dev_env"
dbEnv := "./db_env"
awsEnv := "./aws_env"

usedEnvs := []string{baseEnv, devEnv, dbEnv, awsEnv}
err := env.Parse(usedEnvs...)
```

All Variables from all files are now available...