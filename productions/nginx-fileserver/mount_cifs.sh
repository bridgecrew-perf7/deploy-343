#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

# */5 * * * * bash /root/cron/mount_cifs.sh # check every 5min

USERNAME=xxxx
PASSWORD=yyyy
IP=192.168.1.42
RemotePath=path/of/remote
LocalPath=/path/to/local

# exit if the path is alread mounted
mountpoint $LocalPath ||
{
    mount -t cifs -o username=$USERNAME,password=$PASSWORD,vers=2.0,sec=ntlmssp \
      //$IP/$RemotePath $LocalPath
}
