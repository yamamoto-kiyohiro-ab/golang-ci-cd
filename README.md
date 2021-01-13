# CI / CD Example for Golang

[![CI](https://github.com/yamamoto-kiyohiro-ab/golang-ci-cd/workflows/CI/badge.svg)](https://github.com/yamamoto-kiyohiro-ab/golang-ci-cd/actions?query=workflow%3ACI)
[![codecov](https://codecov.io/gh/yamamoto-kiyohiro-ab/golang-ci-cd/branch/main/graph/badge.svg?token=F2zbZucQiz)](https://codecov.io/gh/yamamoto-kiyohiro-ab/golang-ci-cd)

## 概要

Golang で CI / CD するサンプル

Github Actions, CodeCov, Github Container Registry などを利用し、以下を実現しています

1. PR 作成時に自動テスト & カバレッジ可視化

2. master マージ時に image を作成し GCR に登録

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

## Run (registered image)

```
docker pull ghcr.io/yamamoto-kiyohiro-ab/golang-ci-cd:latest
docker run golang-ci-cd
```
