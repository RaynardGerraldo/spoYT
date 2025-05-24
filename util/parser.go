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

func Parser(body string, duration string, artist string) string{
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

    // check if official channel
    // dont need to remove dupes for this one, channel name and video id listing not correlated
    official_match := regexp.MustCompile(`<p class="channel-name" dir="auto">(.*\n*.*)`)
    official_res := official_match.FindAllStringSubmatch(body, -1)
    official_clean := []string{}
    for _, match := range official_res {
        if len(match) > 1 {
            //official_clean = append(official_clean, strings.ReplaceAll(match[1], "\n", "")) remember check artist name
            official_clean = append(official_clean, match[1])
        }
    }

    for i, dur := range dur_clean {
        for j, ofc := range official_clean {
            if strings.Contains(ofc,artist) {
                // check if channel has ofc checkmark
                if strings.Contains(ofc,"icon ion ion-md-checkmark-circle") {
                    return clean_res[j]
                }
            }
        }

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

        if absInt(yt_dur - sp_dur) <= 2 {
            return clean_res[i]
        }
    }
    return "No match"
}
