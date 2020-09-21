# test

## test


# Overview
* 

# Process sequence

## Preparing
**Since the dynamic site scraping is not implemented yet, download the theme list page by following the steps below.** 

1. Open [web page](https://folio-sec.com/theme).
2. 

# How to build and run

## CASE 1: Using docker image(Recommended).
### Build
Execute [build.sh](https://github.com/bubusuke/webcrawler-service/blob/master/build.sh) and create docker-image named webcrawler-service.

### Run
Execute [run.sh](https://github.com/bubusuke/webcrawler-service/blob/master/run.sh).

## CASE 2 Using webapp binary file.
```
# Build
go get github.com/bubusuke/webcrawler-service
cd ${your go path}/src/github.com/bubusuke/webcrawler-service
go build -trimpath -o webcrawler-service main.go

# Run
./webcrawler-service
```

# Environment variables (Setting Parameters)

| Name     | Value   | Description              |
| :------- | :------ | :----------------------- |
| PORT     | 8080    | Webapplication Port      |
| GIN_MODE | release | GIN flamework log level. |



- [ ] test, comment
- [ ] style sheet
- [ ] docker build
  - [ ] logging web�A�v���P�[�V�����̃��M���O���@���āH
  - [ ] html read
  - [ ] env
- [ ] readme
  - [ ] operation
  - [ ] setup
  - [ ] sequence diagram
  - [ ] how to read html update
