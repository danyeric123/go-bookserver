##
## Build
##

FROM golang:alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /go-bookserver

##
## Deploy
##

FROM alpine

WORKDIR /

COPY --from=build /go-bookserver /go-bookserver

EXPOSE 8080

# Add user that will be used in the container.
RUN adduser --disabled-password davidnagar
USER davidnagar

ENTRYPOINT ["/go-bookserver"]
