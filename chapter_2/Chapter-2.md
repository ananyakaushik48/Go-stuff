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


## Creating a TCP PROXY
### Reader + Writer + COPY
> Here we develop a basic Reader Writer and use it with the Copy function to move data from the source to the destination
* First we define the Reader:
  ```go
    type FuckinReader struct {}
    func (f *FuckinReader) Read(b []byte) (int, error) {
        fmt.Print("in > ")
        return os.Stdin.Read(b)
    }

    type FuckinWriter struct {}
    func (f *FuckinWriter) Write(b []byte) (int, error) {
        fmt.Print("out > ")
        return os.Stdout.Write(p)
    }

    // This will give us access to the required stdin/out functions offered by the OS
    func main() {
        // Here we initialise the Reader and the Writer for the Program to access
        var (
            reader FuckinReader
            writer FuckinWriter
        )
        // Here we use io.Copy to reduce the number of steps involved in doing the mundane task of making explicit calls to the Read and Write functions for each scenario
        if _, err := io.Copy(&writer, &reader); err != nil {
            log.Fatalln("Unable to read/write data")
        }
    }
  ```

> Now we will use what we learnt to make an ECHO-Server which is the pilot project for breaking into the world of comms for any programming language
### Creating an Echo Server


