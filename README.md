# hypercube-cache-redis

Redis driver for [go-hypercube](https://github.com/go-hypercube/go-hypercube)'s `cache.Cache` interface, backed by [`go-redis/v9`](https://github.com/redis/go-redis).

## Install

```bash
go get github.com/go-hypercube/hypercube-cache-redis
```

## Usage

```go
import (
	"github.com/redis/go-redis/v9"
	rediscache "github.com/go-hypercube/hypercube-cache-redis"
)

client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
cache := rediscache.New(client)

app := hypercube.New(cfg, db, cache)
```

## Notes

- Implements the full `github.com/go-hypercube/go-hypercube/cache` interface: `Get`, `Set`, `Delete`, `Has`, `Increment`, `Expire`.
- Misses map to `cache.ErrNotFound`; negative TTLs return `cache.ErrInvalidTTL`.
- `Increment` uses Redis' native `INCRBY`, which creates the key at 0 if absent — matches the interface contract exactly, no emulation needed.
- Need raw Redis data structures (hashes, lists, sets, sorted sets, pub/sub)? Bind the underlying `*redis.Client` into your app's container directly and `Resolve` it from plugins — this driver only covers the plain-cache subset.

## License

MIT
