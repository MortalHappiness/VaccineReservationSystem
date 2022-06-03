FROM golang:1.16

# Install packages
WORKDIR /app
ADD go.mod go.mod
RUN ["go", "mod", "tidy"]

ADD . /app/

RUN go build -o worker

ENTRYPOINT ["/app/worker"]
CMD ["serve"]