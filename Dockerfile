FROM golang:1.22.0-alpine AS build

WORKDIR /code

COPY ./src/go.mod  /code/go.mod
COPY ./src/go.sum  /code/go.sum

RUN go mod tidy
RUN go build -v -o /code/app ./...

FROM alpine:latest

WORKDIR /code

COPY --from=build /code/app /code/app
COPY ./src /code

EXPOSE 8080

CMD ["./app"]
