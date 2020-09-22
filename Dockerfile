# Build Container
FROM golang:1.14-alpine AS builder

RUN apk add git
RUN go get github.com/bubusuke/webcrawler-service-webapp

WORKDIR /go/src/github.com/bubusuke/webcrawler-service-webapp

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -trimpath -o webapp .

# Runtime Container
FROM alpine

WORKDIR /user/local/webcrawler-service-webapp
# Copy sources because tmpl file is required.
COPY ./templates ./templates
# Copy builded binary file
COPY --from=builder /go/src/github.com/bubusuke/webcrawler-service-webapp/webapp ./webapp

ENV APP_HOST=8080
ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_DATABASE_NAME=postgres
ENV DB_USER=postgres
ENV DB_PASSWORD=pass

ENTRYPOINT ["./webapp"]