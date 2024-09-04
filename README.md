# Simple _.env_ Reader

## Installation
Run `go get github.com/strash/dotenvgo`

## Usage
```go
import (
    "log"
    env "github.com/strash/dotenvgo"
)

func main() {
    // Path to your '.env' file
    path := "./some/vars/.env_dev"

    // Load it into memory
    // If the path is omitted (nil), it will look for a '.env' file in the root of your project
    if err := env.Load(&path); err != nil {
        log.Fatal(err.Error())
    }

    // Once loaded, you can read from it
    host, err := env.Get("HOST")
    if err != nil {
        // ... the variable doesn't exist
    }
}
```
