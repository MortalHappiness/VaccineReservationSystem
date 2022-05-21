# User backend

## Prerequisite

### Install Packages

```bash
go mod tidy
```

## Run

```bash
go run main.go serve
```

## Development

### Serve swagger UI

1. Set the following environment

```bash
export SPEC_ENABLED=true
export SPEC_FILES=./docs/swagger-ui
```

2. Restart the server

Go to http://localhost:7712/.spec to see the doc.

### Regenerate swagger document

Please install [go-swagger](https://goswagger.io/install.html) in advance.

```bash
swagger generate spec -o ./docs/swagger-ui/swagger.yaml
```

### Linter

Please install [golangci-lint](https://golangci-lint.run/usage/install/) in advance.

```bash
golangci-lint run
```