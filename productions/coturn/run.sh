#! /usr/bin/env bash
set -eu -o pipefail

_wd=$(pwd)/
_path=$(dirname $0)/


docker-compose pull

mkdir -p logs

docker-compose up -d
