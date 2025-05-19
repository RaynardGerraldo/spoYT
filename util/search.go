package util

import (
    "fmt"
    "net/url"
    "net/http"
    "io"
    "log"
)

func Search(song string, duration string){
    encoded := url.QueryEscape(song)
    client := &http.Client{}
    req_link := fmt.Sprintf("https://yewtu.be/search?q=%s", encoded)
    req, err := http.NewRequest("GET", req_link, nil)

    if err != nil {
        log.Fatal(err)
    }

    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:136.0) Gecko/20100101 Firefox/136.0")
    resp, err := client.Do(req)

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    fmt.Println(req_link)
    Parser(string(body), duration)
    //fmt.Println(string(body))
}
