#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

docker pull bitnami/zookeeper:3.8
docker pull bitnami/kafka:3.2

path=/home/hello/Apps/kafka_2.13-3.2.0
addrs=127.0.0.1:9093

$path/bin/kafka-console-producer.sh --broker-list $addrs --topic test

$path/bin/kafka-console-consumer.sh --bootstrap-server $addrs --topic test --from-beginning
