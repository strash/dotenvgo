# Simple _.env_ Reader

## Installation
Run `go get github.com/strash/dotenvgo`

## Usage
```go
import (
  env "github.com/strash/dotenvgo"
)

func main() {
  // Path to your '.env' file
  path := "/some/path/to/.env_dev"

  // Load it into memory
  // If the path is omitted (nil), it will look for a '.env' file in the root of your project
  env.Load(&path)

  // Once loaded, you can read from it
  host := "localhost"
  if env.Get("HOST") != nil {
    host = *env.Get("HOST")
  }
}
```
