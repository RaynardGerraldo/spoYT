package util

import (
    "fmt"
    "net/url"
    "net/http"
    "io"
    "log"
    "math/rand"
)

func Search(song string, duration string, artist string) string{
    encoded := url.QueryEscape(song)
    client := &http.Client{}
    req_link := fmt.Sprintf("https://inv.nadeko.net/search?q=%s+%s", encoded, "Audio")
    req, err := http.NewRequest("GET", req_link, nil)
    req.Close = true

    if err != nil {
        log.Fatal(err)
    }

    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:136.0) Gecko/20100101 Firefox/136.0")
    // randomize backend every search
    server := fmt.Sprintf("COMPANION_IDD=%d", rand.Intn(6))
    req.Header.Set("Cookie", server)
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    return Parser(string(body), duration, artist)
    //fmt.Println(string(body))
}
