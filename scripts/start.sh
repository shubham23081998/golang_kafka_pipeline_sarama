#!/bin/bash
set -euo pipefail

echo "Stopping any existing containers..."
docker-compose -f docker-compose.yml down

echo "Building Docker images..."
docker-compose -f docker-compose.yml build

echo "Starting Zookeeper and Kafka..."
docker-compose -f docker-compose.yml up -d zookeeper kafka kafka-ui

echo "Waiting for Kafka to be ready..."
sleep 10

echo "Starting generator container..."
docker-compose -f docker-compose.yml up -d generator

echo "Waiting for generator to finish producing..."
# optional: sleep some seconds if generator finishes quickly, or monitor logs
sleep 15

echo "Starting sorter container..."
docker-compose -f docker-compose.yml up -d sorter

echo "Pipeline started. Use 'docker-compose logs -f' to see output."
