package pkg

import (
    "fmt"
    "spoyt/util"
)

func Cli(arg string){
    fmt.Println(arg)
    util.Extractor(arg)
}
