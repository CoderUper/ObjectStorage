#!/bin/bash
export RABBITMQ_SERVER=amqp://test:test@localhost:5672
LISTEN_ADDRESS=192.168.246.129:12341 STORAGE_ROOT=/tmp/1 go run ./dataServer/dataServer.go &
LISTEN_ADDRESS=192.168.246.129:12342 STORAGE_ROOT=/tmp/2 go run ./dataServer/dataServer.go &
LISTEN_ADDRESS=192.168.246.129:12343 STORAGE_ROOT=/tmp/3 go run ./dataServer/dataServer.go &
LISTEN_ADDRESS=192.168.246.129:12344 STORAGE_ROOT=/tmp/4 go run ./dataServer/dataServer.go &
LISTEN_ADDRESS=192.168.246.129:12345 STORAGE_ROOT=/tmp/5 go run ./dataServer/dataServer.go &
LISTEN_ADDRESS=192.168.246.129:12346 STORAGE_ROOT=/tmp/6 go run ./dataServer/dataServer.go &

LISTEN_ADDRESS=192.168.246.129:12347 go run ./apiServer/apiServer.go &
LISTEN_ADDRESS=192.168.246.129:12348 go run ./apiServer/apiServer.go &
