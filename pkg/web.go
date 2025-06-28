package pkg

import (
  "net/http"
  "spoyt/util"
  "path/filepath"
  "github.com/gin-gonic/gin"
  "fmt"
  "strings"
)

func Web() {
  r := gin.Default()
  var data [][]string
  var playlist strings.Builder
  playlist.WriteString("https://www.youtube.com/watch_videos?video_ids=")
  var progress float64
  var start int = 0

  r.LoadHTMLFiles("templates/index.tmpl")
  r.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
        "title": "SpoYT",
    })
  })

  r.POST("/upload", func(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
	    c.String(http.StatusBadRequest, "File upload error: %v", err)
		return
	}

    dst := filepath.Join("uploads", filepath.Base(file.Filename))
    if err := c.SaveUploadedFile(file, dst); err != nil {
        c.String(http.StatusInternalServerError, "Save failed: %v", err)
	    return
    }
    data = util.Converter(dst)
  })

  r.GET("/progress", func(c *gin.Context) {
    if start >= len(data) {
        return
    }

    song := fmt.Sprintf("%s %s", data[start][1], data[start][2])
    result := util.Search(song, data[start][8], data[start][2])
    if result != "No match" {
       playlist.WriteString(result)
       playlist.WriteString(",")
       fmt.Printf("%s added to playlist\n", song)
    } else {
        fmt.Printf("%s not found\n", song)
    }
	progress = progress + 100.0 / float64(len(data))
    c.JSON(http.StatusOK, gin.H{
	    "progress": fmt.Sprintf("%.2f", progress),
	})
    start++
  })

  r.GET("/final", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "link": util.Final(playlist.String()),
    })
    progress = 0
    start = 0
  })

  r.Run()
}
