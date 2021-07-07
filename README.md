# piglatin
A package/CLI written in Golang that translates any sentence in the English alphabet to Pig Latin :pig:

# Installation

`go get github.com/SeanMcGoff/piglatin`

# Example

```go

package main

import (
	"fmt"

	"github.com/SeanMcGoff/piglatin"
)

func main() {
	text := piglatin.ToPigLatin("hello world")
	fmt.Println(text) //ellohay orldway 
}
```
