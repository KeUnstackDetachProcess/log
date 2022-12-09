# log
Simple golang logging package

## Usage example
```go
package main

import (
	"test/pkg/color"
	"test/pkg/log"

	"github.com/fatih/color"
)

func main() {
	log.Initialize("LOG", color.FgMagenta)
  
	log.Ok("Simple example using *colored* format")
	log.Info("Just an *information* log, *2* + *2* = *%d*", 2+2)
	log.Warn("Here's a warning...")
	log.Err("And here we have an *error*, error log also uses *os.Exit(1)*")
}

```
