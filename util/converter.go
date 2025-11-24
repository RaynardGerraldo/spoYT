package util

import (
    "os"
    "encoding/csv"
    "fmt"
    "strings"
    "bytes"
    "io"
    "net/http"
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

func Converter(filename string) ([][]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, fmt.Errorf("Failed to open: %w", filename)
    }

    count,err := lineCounter(file)
    if err != nil {
        return nil, fmt.Errorf("Failed to get line count")  
    }

    if count > 51 {
        fmt.Println("Playlist count exceed 50, will stop at 50th song")
        count = 50
    }

    defer file.Close()

    // reset pointer to beginning of file
    _, err = file.Seek(0, io.SeekStart)
    if err != nil {
        return nil, fmt.Errorf("Failed to reset pointer")
    }

    reader := csv.NewReader(file)
    // skip first row of csv file
	_, err = reader.Read()
    if err != nil {
        return nil, fmt.Errorf("Failed to skip first row of file")
    }

    var data [][]string
    for i := 0; i < count; i++{
        row, err := reader.Read()
        if len(row) == 0 {
            break 
        }
        if err != nil {
           return nil, fmt.Errorf("Failed to read through file")
        }
        data = append(data,row)
    }

    return data, nil
}

func Builder(data [][]string) (string,error) {
    failcount := len(data)
    var failsongs []string
    var playlist strings.Builder
    playlist.WriteString("https://www.youtube.com/watch_videos?video_ids=")
    // song,artist and duration
    for _,j := range data {
        song := fmt.Sprintf("%s %s", j[1], j[2])
        result,err := Search(song, j[9], j[2])
        if err != nil {
            return "", fmt.Errorf("Failed to search: %w", err)
        }
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

    if failcount != len(data) {
        fmt.Printf("%d out of %d songs converted to youtube playlist\n", failcount, len(data))
        for _,fail := range failsongs {
            fmt.Printf("Fail: %s\n", fail)
        }
    }
   
    if playlist.String() == "https://www.youtube.com/watch_videos?video_ids=" {
        return "", fmt.Errorf("No matches found")
    }
    return playlist.String(), nil
}

func Final(playlist string) (string,error) {
    // final playlist link
    finalURL := ""
    if playlist != "https://www.youtube.com/watch_videos?video_ids=" {
        client := &http.Client{}
        req, err := http.NewRequest("GET", playlist, nil)
        if err != nil {
           return "", fmt.Errorf("Failed to build request: %w", err)
        }

        resp, err := client.Do(req)
        if err != nil {
           return "", fmt.Errorf("Failed request: %w", err)
        }
        defer resp.Body.Close()

        finalURL = resp.Request.URL.String()
    }

    return finalURL, nil
}
