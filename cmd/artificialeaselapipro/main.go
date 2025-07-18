// cmd/artificialeaselapipro/main.go
package main

import (
"flag"
"log"
"os"

"artificialeaselapipro/internal/artificialeaselapipro"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialeaselapipro.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
