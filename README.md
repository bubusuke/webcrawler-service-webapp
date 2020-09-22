# webcrawler-service-webapp

## Overview
See [README of webcrawler-service](https://github.com/bubusuke/webcrawler-service/blob/master/README.md).

## How to build and run at local.
### 1. Preparing.
* Prepare a postgres database.
* Execute [DDL](https://github.com/bubusuke/webcrawler-service/tree/master/initdb.d). 

### 2. Build
```
# Build
go get github.com/bubusuke/webcrawler-service-webapp
cd ${your go path}/src/github.com/bubusuke/webcrawler-service-webapp
go build -trimpath -o webcrawler-service main.go

# Run
./webcrawler-service
```
### Environments
See Webapp in [HERE of webcrawler-service](https://github.com/bubusuke/webcrawler-service#environments).
