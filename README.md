# go-web-config-sample

Samples of config for web application implemented by Go.

## Dependencies

- Go: 1.21
- gopkg.in/yaml.v3
- dario.cat/mergo

## Usage

```sh
# Run
go run .

# Run with environment variables
HOGE=a go run .

# Run on specified env
ENV=prod go run .

# Run on test env
go test

# Build and run
go build
./go-web-config-sample
```
