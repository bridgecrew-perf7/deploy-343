#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

####
HTTP_Port=$1
MONGO_Port=$2
docker pull registry.cn-hangzhou.aliyuncs.com/anoy/yapi
docker pull mongo:4

mkdir -p ./mongo/data/db


####
docker run --detach \
  --name=yapi_mongo \
  --restart=always  \
  --publish=${MONGO_Port}:27017        \
  --volume=$PWD/mongo/data/db:/data/db \
   mongo:4

docker run -it --rm       \
  --link yapi_mongo:mongo \
  --entrypoint=npm        \
  --workdir=/api/vendors  \
  registry.cn-hangzhou.aliyuncs.com/anoy/yapi \
  run install-server


####
docker run --detach           \
  --name=yapi_service         \
  --restart=always            \
  --link yapi_mongo:mongo     \
  --publish=${HTTP_Port}:3000 \
  --workdir /api/vendors      \
  registry.cn-hangzhou.aliyuncs.com/anoy/yapi \
  server/app.js
# default account: admin@admin.com
# default password: ymfe.org

docker cp yapi:/api api
docker rm -f yapi_service

docker run --detach           \
  --name=yapi_service         \
  --restart=always            \
  --link yapi_mongo:mongo     \
  --publish=${HTTP_Port}:3000 \
  --volume=$PWD/api:/api      \
  --workdir=/api/vendors      \
  registry.cn-hangzhou.aliyuncs.com/anoy/yapi \
  server/app.js
