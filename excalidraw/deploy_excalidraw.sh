#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

PORT=$1

docker run --rm --detach    \
  --name excalidraw_service \
  --publish=${PORT}:80      \
  excalidraw/excalidraw:latest
