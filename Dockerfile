FROM golang:1.22.0-alpine AS build

WORKDIR /code

COPY ./src/go.sum /code/go.sum
COPY ./src/go.mod /code/go.mod
RUN go mod tidy

COPY ./src /code
RUN go build -o /code/app

FROM alpine:3.19.1

WORKDIR /code

COPY --from=build /code/app /code/app

EXPOSE 8080

CMD ["./app"]
