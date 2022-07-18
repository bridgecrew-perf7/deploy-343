### Kafka docker
---

#### 
- site:
 - https://hub.docker.com/r/bitnami/kafka/
- pull images:
```bash
docker pull bitnami/kafka:3.2
docker pull bitnami/zookeeper:3.8
```
- up
```bash
curl -sSL https://raw.githubusercontent.com/bitnami/bitnami-docker-kafka/master/docker-compose.yml > docker-compose.yml
docker-compose up -d
```
