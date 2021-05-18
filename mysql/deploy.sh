#! /usr/bin/env bash
set -eu -o pipefail

export MYSQL_ROOT_PASSWORD=$1
envsubst < $(dirname $0)/deployment.yaml > docker-compose.yaml

docker-compose pull
docker-compose up -d
