# First Go

## Introduction

## Setup

```bash
make postgres
make createdb
make migrateup
```

## Install

- Step 1: make query by cmd:

```bash
make sqlc
```

- Step 2: make mock store by cmd:

```bash
make mock
```

- Step 3: run the server by cmd:

```bash
make server
```

## test

- By psql:

```bash
docker exec -it postgres16 psql -U root -d testGo
```
