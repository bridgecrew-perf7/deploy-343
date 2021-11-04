#! /usr/bin/env bash
set -eu -o pipefail

# registry: https://github.com/acmesh-official/acme.sh
# cron: 0 0 * * *
# location: /${HOME}/.acme.sh/${DOMAIN}
{
  ${HOME}/.acme.sh/acme.sh --cron --home ${HOME}/.acme.sh
} >> $(dirname $0)/acme.$(date +"%Y-%m").log 2>&1
