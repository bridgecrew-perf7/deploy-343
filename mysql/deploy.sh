#! /usr/bin/env bash
set -eu -o pipefail

export MYSQL_ROOT_PASSWORD=$1
export PORT=$2
envsubst < $(dirname $0)/deployment.yaml > docker-compose.yaml

docker-compose pull
docker-compose up -d
