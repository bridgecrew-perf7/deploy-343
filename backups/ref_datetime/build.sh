#! /usr/bin/env bash
set -eu -o pipefail

#### config and check
# registry=$(printenv DOCKER_Registry)
registry=registry.cn-shanghai.aliyuncs.com/d2jvkpn
image="$registry/datetime:latest"

#### build local image
echo ">>> Building image: $image"
docker pull golang:1.16-alpine
docker pull alpine
docker build -f Dockerfile --no-cache -t "$image" .
docker image prune --force --filter label=stage=datetime_builder

#### push to registry
echo ">>> Pushing image: $image"
sudo docker push $image
