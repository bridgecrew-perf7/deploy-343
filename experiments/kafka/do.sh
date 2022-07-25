#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

####
docker pull bitnami/zookeeper:3.8
docker pull bitnami/kafka:3.2

docker-compose -f docker-compose.zookeeper.yaml up -d

addrs=127.0.0.1:9093

kafka-console-producer.sh --broker-list $addrs --topic test

kafka-console-consumer.sh --bootstrap-server $addrs --topic test --from-beginning
