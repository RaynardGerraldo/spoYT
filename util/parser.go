package util

import (
    "fmt"
    "regexp"
    "slices"
)

func Parser(body string, duration string){
    // get video ids
    match := regexp.MustCompile(`.*href="/watch\?v=([^"&]*).*`)
    res := match.FindAllStringSubmatch(body, -1)
    clean_res := []string{}
    fmt.Printf("Type of res: %T\n", res)

    // take matching group result 
    for _, match := range res {
        if len(match) > 1 {
            clean_res = append(clean_res, match[1])
        }
    }

    // remove dupes
    slices.Sort(clean_res)
    clean_res = slices.Compact(clean_res)
    fmt.Println(clean_res)
    fmt.Println(len(clean_res))

    // get video duration
    fmt.Println(duration)
}
