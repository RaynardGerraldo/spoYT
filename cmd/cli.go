package cmd

import (
    "fmt"
    "spoyt/util"
)

func Cli(filename string){
    fmt.Println(filename)
    data := util.Converter(filename)
    playlist := util.Builder(data)
    fmt.Println(util.Final(playlist))
}
