package util

import (
    "os"
    "log"
    "encoding/csv"
    "fmt"
)

func Extractor(filename string){
    file, err := os.Open(filename)

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    reader := csv.NewReader(file)
    data, err := reader.ReadAll()

    if err != nil {
        log.Fatal(err)
    }

    // song,artist and duration
    for _,j := range data {
        song := fmt.Sprintf("%s %s", j[1], j[2])
        Search(song, j[8])
    }
}
