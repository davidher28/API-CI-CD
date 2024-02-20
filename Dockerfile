FROM golang:1.22.0-alpine AS build

WORKDIR /code

COPY ./src/. /code/.

RUN go mod tidy
RUN go build -o /code/app

FROM alpine

WORKDIR /code

COPY --from=build /code/app /code/app
COPY ./src /code

EXPOSE 8080

CMD ["./app"]
