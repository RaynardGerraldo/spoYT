package pkg

import (
    "fmt"
    "spoyt/util"
)

func Cli(arg string){
    fmt.Println(arg)
    data := util.Converter(arg)
    playlist := util.Builder(data)
    fmt.Println(util.Final(playlist))
}
