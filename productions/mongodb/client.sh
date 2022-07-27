#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

docker exec -it mongo_db mongo -u root -p

exit

```js
db = db.getSiblingDB('admin')
db.changeUserPassword("root", passwordPrompt())

use products
db.changeUserPassword("accountUser", passwordPrompt())
```
