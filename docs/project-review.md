# Technical Review — rediscc

## What is it

`rediscc` is a Go library that provides a simplified wrapper over the official `go-redis/v9` client. It offers a clean interface for basic Redis operations (Get, Set, Del, Keys, Publish) with optional debug logging support.

## Architecture

The project follows a simplified **ports & adapters** (hexagonal) pattern:

```
+-------------------+
|  RedisRepository  |  <-- Interface (ports.go)
|      Ports        |
+--------+----------+
         |
+--------+----------+
|  RedisDataStore   |  <-- Implementation (rediscc.go)
|  (adapter)        |
+--------+----------+
         |
+--------+----------+
|  go-redis/v9      |  <-- Underlying Redis client
+-------------------+
```

- `ports.go` defines the `RedisRepositoryPorts` interface.
- `rediscc.go` implements the interface through `RedisDataStore`.
- `structs.go` defines the shared data structures.

## Tech Stack

| Component | Technology |
|---|---|
| Language | Go 1.25 |
| Redis Client | github.com/redis/go-redis/v9 v9.12.0 |
| Hashing (indirect) | github.com/cespare/xxhash/v2 |

## Project Structure

```
rediscc-go/
  go.mod          -- Module definition and dependencies
  go.sum          -- Dependency checksums
  LICENSE         -- Project license
  README.md       -- Basic documentation
  ports.go        -- RedisRepositoryPorts interface
  structs.go      -- Structs: RedisDataStore, RedisOptions
  rediscc.go      -- Implementation: Connect + CRUD operations
```

## Public API

| Function/Method | Signature | Description |
|---|---|---|
| `Connect` | `Connect(ctx, redisUri, dbNumber, options) (*RedisDataStore, error)` | Connects to Redis by parsing URI + DB number |
| `Publish` | `(r *RedisDataStore) Publish(ctx, channel, payload) error` | Publishes a message to a pub/sub channel |
| `Get` | `(r *RedisDataStore) Get(ctx, key) (string, error)` | Gets the value of a key as a string |
| `GetRaw` | `(r *RedisDataStore) GetRaw(ctx, key) *redis.StringCmd` | Gets the raw unprocessed Redis command |
| `Set` | `(r *RedisDataStore) Set(ctx, key, value, expiration) error` | Sets a key with a value and TTL |
| `Del` | `(r *RedisDataStore) Del(ctx, key) error` | Deletes a key |
| `Keys` | `(r *RedisDataStore) Keys(ctx, pattern) ([]string, error)` | Finds keys matching a glob pattern |

## Domain Models

| Struct | Fields | Description |
|---|---|---|
| `RedisDataStore` | `Client *redis.Client`, `Options *RedisOptions` | Main store wrapping the Redis client |
| `RedisOptions` | `Debug bool`, `DebugPayload bool` | Configuration options for logging |

## Interface

```go
type RedisRepositoryPorts interface {
    Publish(ctx context.Context, channel string, payload interface{}) error
    Get(ctx context.Context, key string) (string, error)
    GetRaw(ctx context.Context, key string) *redis.StringCmd
    Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
    Del(ctx context.Context, key string) error
    Keys(ctx context.Context, pattern string) ([]string, error)
}
```

## Configuration

The connection is configured via parameters in `Connect`:

| Parameter | Type | Description |
|---|---|---|
| `redisUri` | `string` | Redis connection URI (e.g. `redis://:password@host:6379`) |
| `dbNumber` | `string` | Redis database number (0-15) |
| `options` | `*RedisOptions` | Debug options (can be `nil`) |

## Development

```bash
# Install as dependency
go get github.com/codecraftkit/rediscc

# Run tests (no tests currently exist)
go test ./...
```

## Observations and Potential Improvements

### Urgent

- **No tests**: The project has no test files whatsoever. This is critical for a library that handles data connections.
- **Inconsistent README**: The README mentions "compatible with the official Redis driver for Go (mongo-driver)" — this is a copy/paste error from a MongoDB library. The usage example is also empty (`//TODO`).
- **`Keys` in production**: The `Keys` method uses Redis `KEYS`, which is O(N) and blocks the server. In production, `SCAN` should be used instead, or at the very least the warning should be documented.

### Recommended

- **`dbNumber` as string**: The `dbNumber` parameter is a string that gets concatenated into the URL. It would be safer to accept an `int` and validate the range (0-15).
- **Logging with `fmt.Println`**: The debug system uses `fmt.Println` directly. It would be better to accept a configurable `logger` or use `log/slog`.
- **No `Subscribe` support**: `Publish` exists but there is no `Subscribe` method to complete the pub/sub pattern.
- **No `Close`**: There is no method to close the connection (`Client.Close()`). This can cause connection leaks.
- **No context timeout in Connect**: The `Ping` has no explicit timeout; it depends on the context passed by the user.
- **Exposing `Client` directly**: `RedisDataStore.Client` is public, which breaks encapsulation. If the goal is to use the interface, the client should be private.
