package util

import (
    "os"
    "log"
    "encoding/csv"
    "fmt"
    "strings"
    "bytes"
    "io"
)

// read file bytes every 32 kb, adds to count on \n
func lineCounter(r io.Reader) (int, error) {
    buf := make([]byte, 32*1024)
    count := 0
    lineSep := []byte{'\n'}

    for {
        c, err := r.Read(buf)
        count += bytes.Count(buf[:c], lineSep)

        switch {
        case err == io.EOF:
            return count, nil

        case err != nil:
            return count, err
        }
    }
}

func Converter(filename string) string{
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }

    count,err := lineCounter(file)
    if err != nil {
        log.Fatal(err)
    }

    if count > 51 {
        fmt.Println("Playlist count exceed 50, will stop at 50th song")
        count = 50
    }
    failcount := count
    
    defer file.Close()

    // reset pointer to beginning of file
    _, err = file.Seek(0, io.SeekStart)
    if err != nil {
        log.Fatal(err)
    }

    reader := csv.NewReader(file)
    // skip first row of csv file
	_, err = reader.Read()
    if err != nil {
        log.Fatal(err)
    }

    var data [][]string
    for i := 0; i < 50; i++{
        row, err := reader.Read()
        if err != nil {
           log.Fatal(err)
        }
        data = append(data,row)
    }

    //data, err := reader.ReadAll()
    //if err != nil {
        //log.Fatal(err)
    //}


    var failsongs []string
    var playlist strings.Builder
    playlist.WriteString("https://www.youtube.com/watch_videos?video_ids=")
    // song,artist and duration
    for _,j := range data {
        song := fmt.Sprintf("%s %s", j[1], j[2])
        result := Search(song, j[8], j[2])
        if result != "No match" {
            playlist.WriteString(result)
            playlist.WriteString(",")
            fmt.Printf("%s added to playlist\n", song)
        } else {
            failcount -= 1
            failsongs = append(failsongs, song)
            fmt.Printf("%s not found\n", song)
        }
    }

    if failcount != count {
        fmt.Printf("%d out of %d songs converted to youtube playlist\n", failcount, count)
        for _,fail := range failsongs {
            fmt.Printf("Fail: %s\n", fail)
        }
    }

    return playlist.String()
}
