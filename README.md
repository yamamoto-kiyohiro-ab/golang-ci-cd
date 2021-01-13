# CI / CD Example for Golang

## 概要

Golang で CI / CD するサンプル

Github Actions と CodeCov を利用

## Test

```
go test ./... -v 
```

## Run

```
go run ./cmd/dummyapp
```

## Build & Run

```
go build -o ./build/dummyapp ./cmd/dummyapp
./build/dummyapp
```

## Build image & Run

```
docker build -t dummyapp .
docker run dummyapp
```
