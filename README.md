# CI / CD Example for Golang

[![CI](https://github.com/yamamoto-kiyohiro-ab/golang-ci-cd/workflows/CI/badge.svg)](https://github.com/yamamoto-kiyohiro-ab/golang-ci-cd/actions?query=workflow%3ACI)
[![codecov](https://codecov.io/gh/yamamoto-kiyohiro-ab/golang-ci-cd/branch/main/graph/badge.svg?token=F2zbZucQiz)](https://codecov.io/gh/yamamoto-kiyohiro-ab/golang-ci-cd)

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
