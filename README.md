git push main v1.0.0# in-mem-cache

`in-mem-cache` is a simple in-memory cache library for Go. It provides methods to set, get, delete values in a cache
stored in memory, and save/load the cache state to/from a file. This library is useful for temporarily storing data that
needs fast access without relying on an external storage system.

## Features

- **Set**: Store a value in the cache with a specified key.
- **Get**: Retrieve a value from the cache by key.
- **Delete**: Remove a value from the cache by key.
- **SaveToFile**: Save the current state of the cache to a file.
- **LoadFromFile**: Load the cache state from a file.

## Installation

To install the package, run:

```bash
go get -u github.com/Takeso-user/in-mem-cache
```

## Usage

Below is an example of how to use the in-mem-cache package in your project:

```go
package main

import (
    "fmt"
	"time"
	"github.com/Takeso-user/in-mem-cache/cache"
)

func main() {
    // Initialize a new cache
	cache := cache.NewCache()

    // Set a value in the cache
	cache.Set("userId", 42, 10*time.Second)

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

	// Save the cache state to a file
	err := cache.SaveToFile("cache_state.json")
	if err != nil {
		fmt.Println("Error saving cache to file:", err)
	}

	// Load the cache state from a file
	err = cache.LoadFromFile("cache_state.json")
	if err != nil {
		fmt.Println("Error loading cache from file:", err)
	}
}
```

## API

### Set
```go
Set(key string, value interface{}, ttl time.Duration) bool
```

Stores a value in the cache under the specified key with a time-to-live (TTL).

### Get
```go
Get(key string) (interface{}, bool)
```
Retrieves a value from the cache by key. Returns the value and a boolean indicating if the key was found.

### Delete
```go
Delete(key string) bool
```

Removes the value associated with the key from the cache. Returns true if the key was found and deleted, otherwise
false.

### SaveToFile

```go
SaveToFile(filename string) error
```

Saves the current state of the cache to a file.

### LoadFromFile

```go
LoadFromFile(filename string) error
```

Loads the cache state from a file.

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License.


