package cmd

import (
    "fmt"
    "spoyt/util"
    "os"
)

func Cli(arg string){
    fmt.Println(arg)
    data,err := util.Converter(arg)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to read csv: %v\n", err)
        os.Exit(1)
    }
    playlist,err := util.Builder(data)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to build playlist link: %v\n", err)
        os.Exit(1)
    }
    final,err := util.Final(playlist)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to get final link: %v\n", err)
        os.Exit(1)
    }
    fmt.Println(final)
}
