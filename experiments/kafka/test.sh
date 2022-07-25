#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})


####
go run producer.go

go run consumer.go

go run consumer_group.go


####
addrs="127.0.0.1:9093 127.0.0.1:9095 127.0.0.1:9097"
go run producer.go $addrs

go run consumer.go $addrs

go run consumer_group.go $addrs
