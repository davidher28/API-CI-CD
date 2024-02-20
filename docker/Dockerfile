FROM golang:1.22.0

WORKDIR /go/src

COPY .. .

RUN go mod tidy
RUN go build -v ./...

EXPOSE 8080

CMD ["go run src/main.go"]
