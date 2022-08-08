#### References
- https://raw.githubusercontent.com/redis/redis/7.0/redis.conf
- https://redis.io/docs/manual/persistence/
- https://redis.io/commands/acl-setuser/

#### Commandlines
```redis
config get dir
config get *

bgsave

save 300 10

info persistence

auth d2jvkpn

# reset default administrator password
config set requirepass d2jvkpn

acl setuser hello -keys -flushall -flushdb -config on >world
acl list

auth hello world
```
