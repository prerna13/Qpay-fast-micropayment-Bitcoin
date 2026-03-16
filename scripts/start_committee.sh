#!/bin/bash

echo "Starting committee nodes..."

PORTS=(8000 8001 8002)

for PORT in "${PORTS[@]}"
do
   PORT=$PORT go run cmd/committee_node/main.go &
done

wait
