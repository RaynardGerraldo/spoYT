package util

import (
    "os"
    "bufio"
    "log"
)

func Extractor(filename string){
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        Search(scanner.Text())
        //fmt.Println("Song: ", scanner.Text())
    }
}
