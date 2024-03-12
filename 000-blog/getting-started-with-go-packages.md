# Learning Go from Scractch.

## Initializing, building, and running a Go package

1. Open your terminal or command prompt and navigate to the directory where you want to create your Go project.
2. Run the following command to initialize Go modules:

```bash
go mod init leroy.africa/leroy/hello
```

3. Create a new file named leroy.go and add your code. Here is an example of the code you can add:

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Leroy rocks")
}

```

4. Run the following command to build your go packages:

```bash
go build -o hello_leroy main.go
```

5. Run the following command to execute your program:

```bash
./hello_leroy
```
