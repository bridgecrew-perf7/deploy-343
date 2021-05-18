#! /usr/bin/env bash
set -eu -o pipefail

docker pull centos:7

image="registry.cn-shanghai.aliyuncs.com/d2jvkpn/jupyter:latest"
docker build --no-cache -t $image .
docker push $image

exit

docker run --detach --name=jupyter_latest_service --env=JUPYTER_Port=9000 \
  --env=JUPYTER_Password=123456 --publish=9000:9000 $image
