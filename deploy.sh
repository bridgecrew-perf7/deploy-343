#! /usr/bin/env bash
set -eu -o pipefail

docker-compose --env-file ./docker-env up -d

# docker inspect mysql_service | jq -r ".[0].NetworkSettings.IPAddress"172.17.0.2
