# User backend

## Prerequisite

### Install Packages

```bash
go mod tidy
```

## Run

### Set environment
```bash
# bigtable
export PROJECT_ID=my-project
export INSTANCE_ID=my-instance
export TABLE_NAME=vaccine-reservation-system
# secret
export SECRET=my-secret
```

```bash
go run main.go serve
```

## Development

### Set environment

1. Set the following environment

```bash
# serve swagger spec
export SPEC_ENABLED=true
export SPEC_FILES=./docs/swagger-ui

# gin log
export ACCESS_LOG=true
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