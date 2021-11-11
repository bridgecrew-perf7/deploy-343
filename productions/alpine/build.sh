#! /usr/bin/env bash
set -eu -o pipefail

_wd=$(pwd)/
_path=$(dirname $0)/

####
name="registry.cn-shanghai.aliyuncs.com/d2jvkpn/alpine"
tag="3"
image=$name:$tag

####
docker pull alpine:$tag
echo ">>> Building image: $image"
docker build -f ./Dockerfile --no-cache -t "$image" .

####
docker push $image

imgs=$(docker images --filter=dangling=true --quiet $name)
[[ -z "$imgs" ]] || docker rmi $imgs
