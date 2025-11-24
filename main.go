package main

import (
    "os"
    "spoyt/cmd"
)

func main() {
    if len(os.Args) == 2 {
        cmd.Cli(os.Args[1])
    } else{
        cmd.Web()
    }
}
