# Golang Kafka Pipeline (Sarama) - Example

This project contains:
- Generator: produces CSV records to Kafka topic `source`.
- Sorter: consumes `source`, performs external sort per key (id, name, continent) and produces to topics `id`, `name`, `continent`.
- Docker + docker-compose for Kafka (Zookeeper + Kafka) and a buildable app image.

## Quickstart (WSL / Linux / Mac)

1. Build the app image:
   ```
   ./scripts/build.sh
   ```

2. Start Kafka, Zookeeperm, Generator and Sorter:
   ```
   ./scripts/start.sh
   ```

5. Verify:
   ```
   ./scripts/verify.sh id
   ```

Notes:
- Adjust TOTAL_MESSAGE environment present in docker compose file to 50000000 for full dataset (may take long).

