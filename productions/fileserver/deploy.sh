#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

# https://www.yanxurui.cc/posts/server/2017-03-21-NGINX-as-a-file-server/

# docker build -f Dockerfile -t fileserver .

export PORT="8100"
envsubst < ${_path}/deploy.yaml > docker-compose.yaml
docker-compose up -d
