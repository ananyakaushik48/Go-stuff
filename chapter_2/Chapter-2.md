# CHAPTER 2

## How to run a go program

```go
package main
import (
 "fmt"
)
func main() {
 fmt.Println("Hello, Black Hat Gophers!")
}
```

```bash
$ go build hello.go
$ ./hello
Hello, Black Hat Gophers!
```
> The following command will reduce the binary size by approximately 30 percent:
> 
> ```$ go build -ldflags "-w -s"```


## Creating a 