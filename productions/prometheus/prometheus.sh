#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

# https://gabrieltanner.org/blog/collecting-prometheus-metrics-in-golang/
# https://hub.docker.com/r/prom/prometheus
# https://hub.docker.com/r/grafana/grafana/tags

docker pull prom/prometheus:main
docker pull grafana/grafana:main
docker-compose up -d

#### reset grafana password
docker exec -it grafana grafana-cli admin reset-admin-password PASSWORD

# grafana url: http://192.168.1.1:3022, admin PASSWORD
# add prometheus data source: http://192.168.1.1:3023

ls -alh /var/lib/docker/volumes/go-web_grafana-storage/_data
