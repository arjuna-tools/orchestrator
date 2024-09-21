FROM golang:1.23-alpine as build-stage

WORKDIR /tmp/build

COPY . .

# Build the project
RUN go build main.go

FROM alpine:3

WORKDIR /app

# Install needed deps
RUN apk add --no-cache tini

COPY --from=build-stage /tmp/build/main main

ENTRYPOINT ["tini", "--"]
CMD ["/app/main"]