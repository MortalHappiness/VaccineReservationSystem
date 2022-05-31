FROM golang:1.16
WORKDIR /go/src
ADD go.mod go.mod
RUN ["go", "mod", "tidy"]
ADD internal internal
ADD .golangci.yml .golangci.yml
ADD main.go main.go
ENTRYPOINT ["go", "run", "main.go", "serve"]