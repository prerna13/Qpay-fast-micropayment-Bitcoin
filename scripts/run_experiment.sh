#!/bin/bash

echo "Starting Q_pay experiment"

echo "Launching committee nodes..."
./scripts/start_committee.sh &

sleep 2

echo "Starting aggregator..."
go run cmd/aggregator/main.go &

sleep 2

echo "Starting seller..."
go run cmd/seller/main.go &

sleep 2

echo "Starting buyer..."
go run cmd/buyer/main.go
