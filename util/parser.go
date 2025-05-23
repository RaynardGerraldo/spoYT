package util

import (
    "regexp"
    "slices"
    "strings"
    "strconv"
)

func absInt(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func Parser(body string, duration string) string{
    // get video ids
    match := regexp.MustCompile(`.*href="/watch\?v=([^"&]*).*`)
    res := match.FindAllStringSubmatch(body, -1)
    clean_res := []string{}

    // take matching group result 
    for _, match := range res {
        if len(match) > 1 {
            clean_res = append(clean_res, match[1])
        }
    }

    // remove dupes
    //slices.Sort(clean_res)
    clean_res = slices.Compact(clean_res)
    // get video duration
    dur_match := regexp.MustCompile(`.*<p class="length">([0-9][0-9]?:[0-9][0-9]:?[0-9]?[0-9]?)</p>`)
    dur_res := dur_match.FindAllStringSubmatch(body, -1)
    dur_clean := []string{}
    for _, match := range dur_res {
        if len(match) > 1 {
            dur_clean = append(dur_clean, match[1])
        }
    }

    for i, dur := range dur_clean {
        if string(duration[1:]) == dur {
            return clean_res[i]
        }

        yt_split := strings.Split(dur, ":")
        sp_split := strings.Split(duration, ":")

        min,_ := strconv.Atoi(yt_split[0])
        sec,_ := strconv.Atoi(yt_split[1])
        yt_dur := min * 60 + sec

        min,_ = strconv.Atoi(sp_split[0])
        sec,_ = strconv.Atoi(sp_split[1])
        sp_dur := min * 60 + sec

        if absInt(yt_dur - sp_dur) <= 5 {
            return clean_res[i]
        }

    }

    return "No match"
}
