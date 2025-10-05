#!/usr/bin/env bash
set -euo pipefail

TOPIC=${1:-id}
echo "Consuming first 20 messages from topic $TOPIC"

docker run --network golang_kafka_pipeline_sarama_goapp --rm \
  confluentinc/cp-kafka:7.2.1 \
  kafka-console-consumer --bootstrap-server kafka:29092 \
  --topic "$TOPIC" --from-beginning --max-messages 20  --group test

scho "-----------------------------------------------------------------------------"

echo "Consuming first 20 messages from topic name"

docker run --network golang_kafka_pipeline_sarama_goapp --rm \
  confluentinc/cp-kafka:7.2.1 \
  kafka-console-consumer --bootstrap-server kafka:29092 \
  --topic name --from-beginning --max-messages 20  --group test

scho "-----------------------------------------------------------------------------"

echo "Consuming first 20 messages from topic continent"

docker run --network golang_kafka_pipeline_sarama_goapp --rm \
  confluentinc/cp-kafka:7.2.1 \
  kafka-console-consumer --bootstrap-server kafka:29092 \
  --topic continent --from-beginning --max-messages 20  --group test



