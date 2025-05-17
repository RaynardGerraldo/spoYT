package util

import (
    "fmt"
    "net/url"
    "net/http"
    "io"
    "log"
)

func Search(song string){
    encoded := url.QueryEscape(song)
    req_link := fmt.Sprintf("https://www.youtube.com/results?search_query=%s", encoded)
    resp, err := http.Get(req_link)

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    fmt.Println(string(body))
}
