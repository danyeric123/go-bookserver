##
## Build
##

FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY main.go ./
COPY ./pkg ./

RUN go build -o /go-bookserver

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /go-bookserver /go-bookserver

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/go-bookserver"]
