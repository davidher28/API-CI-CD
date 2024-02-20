FROM golang:1.22.0-alpine

WORKDIR /code

COPY ./src/go.mod  /code/go.mod
COPY ./src/go.sum  /code/go.sum

RUN go mod tidy
RUN go build -v ./...

COPY ./src /code

EXPOSE 8080

CMD ["go", "run", "main.go"]
