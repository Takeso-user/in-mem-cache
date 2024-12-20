# in-mem-cache

`in-mem-cache` is a simple in-memory cache library for Go. It provides methods to set, get, and delete values in a cache
stored in memory. This library is useful for temporarily storing data that needs fast access without relying on an
external storage system.

## Features

- **Set**: Store a value in the cache with a specified key.
- **Get**: Retrieve a value from the cache by key.
- **Delete**: Remove a value from the cache by key.

## Installation

To install the package, run:

```bash
go get -u github.com/Takeso-user/in-mem-chache
```

## Usage

Below is an example of how to use the in-mem-cache package in your project:

```go
package main

import (
    "fmt"
    "github.com/Takeso-user/in-mem-cache/cache"
)

func main() {
    // Initialize a new cache
    cache := cache.NewCache(10 * time.Second)

    // Set a value in the cache
    cache.Set("userId", 42)

    // Get a value from the cache
    userId, found := cache.Get("userId")
    if found {
        fmt.Println("UserID:", userId) // Output: UserID: 42
    } else {
        fmt.Println("UserID not found")
    }

    // Delete a value from the cache
    cache.Delete("userId")

    // Try to get the deleted value
    userId, found = cache.Get("userId")
    if found {
        fmt.Println("UserID:", userId)
    } else {
        fmt.Println("UserID not found") // Output: UserID not found
    }
}
```

## API

### Set

```go
Set(key string, value interface{})
```

Stores a value in the cache under the specified key.

### Get

```go
Get(key string) (interface{}, bool)
```

Retrieves a value from the cache by key. Returns the value and a boolean indicating if the key was found.

### Delete

```go
Delete(key string) bool
```

Removes the value associated with key from the cache. Returns true if the key was found and deleted, otherwise false.

Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

License

This project is licensed under the MIT License.