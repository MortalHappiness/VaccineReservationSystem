FROM golang:1.16

# Install packages
WORKDIR /app
ADD go.mod go.mod
RUN ["go", "mod", "tidy"]

ADD . /app/
ENTRYPOINT ["go", "run", "main.go", "serve"]