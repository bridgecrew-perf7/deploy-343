#! /usr/bin/env bash
set -eu -o pipefail

# cron: 0 0 1 * *
{
  date +">>> %FT%T%z"

# $ crontab -e
# *	minute, 0-59
# *	hour, 0-23
# *	day of month, 1- 31
# *	month, 1-12
# *	day of week, 0-6

# append new line
# $ crontab -l
  yesterday=$(date -d 'yesterday' '+%Y-%m-%d')
  pid=$(cat /var/run/nginx.pid)

  for f in $(ls ${HOME}/Work/nginx/log/*.log); do
    test -s $f || continue || true
    out=${f%\.log}.${yesterday}.log
    echo "    saving $out"
    mv $f $out
    pigz $out &
  done

  kill -USR1 $pid

  wait
} >> $(dirname $0)/nginx_log_rotate.$(date +"%Y-%m").log 2>&1
