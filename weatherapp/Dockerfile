FROM golang:alpine as build

ADD . /go/src/weatherapp
WORKDIR /go/src/weatherapp

RUN go install 

# Copy just the executable to the release image.
FROM alpine:3.6

RUN apk add --no-cache ca-certificates

WORKDIR /app
EXPOSE 8080

COPY --from=build /go/bin/weatherapp .

ENTRYPOINT ["./weatherapp"]