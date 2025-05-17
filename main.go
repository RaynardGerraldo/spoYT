package main

import (
    "os"
    "spoyt/pkg"
)

func main() {
    if len(os.Args) == 2 {
        pkg.Cli(os.Args[1])
    } else{
        pkg.Web()
    }
}
