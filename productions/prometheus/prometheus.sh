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

# http://192.168.1.1:3023
curl http://192.168.1.1:3023/metrics

# grafana url: http://192.168.1.1:3022, admin PASSWORD
# add prometheus data source: http://192.168.1.1:3023/metrics

ls -alh /var/lib/docker/volumes/go-web_grafana-storage/_data

exit

# promhttp_metric_handler_requests_total{code="200"}
# promhttp_metric_handler_requests_total{code="200", instance="192.168.0.8:3023", job="prometheus"}

process_cpu_seconds_total
# process_cpu_seconds_total{instance="192.168.0.8:3023", job="prometheus"}

go_goroutines
# go_goroutines{instance="192.168.0.8:3023", job="prometheus"}

go_gc_duration_seconds{quantile="0.75"}

go_threads

go_memstats_sys_bytes
