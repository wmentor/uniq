# uniq

uniq is simple unique value generator.

# Summary

* Written on pure Go
* Thread safe
* It can be used in a distributed system
* Doesn't create goroutines
* MIT license

# Require

* Golang (version >= 1.10)

# Install

```
go get github.com/wmentor/uniq
```

# Usage

```go
package main

import (
  "fmt"

  "github.com/wmentor/uniq"
)

func main() {
	
	for i := 0 ; i < 10 ; i++ {
	  fmt.Println(uniq.New())
	}

}
```

Output is like this:

```
a801562de175015fd7d574d3129b47842
a801562de175015fd7d574d350c3c7843
a801562de175015fd7d574d35d82f7844
a801562de175015fd7d574d3816dd7845
a801562de175015fd7d574d3939407846
a801562de175015fd7d574d39feaf7847
a801562de175015fd7d574d3a89637848
a801562de175015fd7d574d3b0f2a7849
a801562de175015fd7d574d3b9569784a
a801562de175015fd7d574d3c1b01784b
```
