FROM golang:1.15.3 AS build
WORKDIR /go/src/github.com/pmaterer/makemake
COPY go.mod .

RUN go mod download

COPY . .
RUN go test -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM alpine
COPY --from=build /go/src/github.com/pmaterer/makemake/makemake /makemake
EXPOSE 8080
ENTRYPOINT ["/makemake"]