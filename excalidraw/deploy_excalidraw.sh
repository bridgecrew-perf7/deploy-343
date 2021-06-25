#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

PORT=$1

docker run --detach         \
  --name excalidraw_service \
  --restart=always          \
  --publish=${PORT}:80      \
  excalidraw/excalidraw:latest
