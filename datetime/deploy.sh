#! /usr/bin/env bash
set -eu -o pipefail

#### config
PORT="$1"
image="registry.cn-shanghai.aliyuncs.com/d2jvkpn/datetime:latest"

#### deploy
docker pull $image
export PORT=${PORT}
envsubst < deployment.yaml > docker-compose.yaml

docker-compose pull
docker-compose up -d

echo ">>> HTTP server port: $PORT"
docker logs datetime_service
