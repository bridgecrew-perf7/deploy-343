#! /usr/bin/env bash
set -eu -o pipefail

#### deploy
export PORT=$1
envsubst < $(dirname $0)/deployment.yaml > docker-compose.yaml

docker-compose pull
docker-compose up -d
