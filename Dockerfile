# Build Container
FROM golang:1.14-alpine AS builder

RUN apk add git
RUN go get github.com/bubusuke/webcrawler-service

WORKDIR /go/src/github.com/bubusuke/webcrawler-service

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -trimpath -o webcrawler-service main.go

# Runtime Container
FROM alpine

WORKDIR /user/local/webcrawler-service
# Copy sources because tmpl file is required.
COPY . .
# Copy builded binary file
COPY --from=builder /go/src/github.com/bubusuke/webcrawler-service/webcrawler-service ./webcrawler-service

ENTRYPOINT [ "./webcrawler-service" ]