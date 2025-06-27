package pkg

import (
  "net/http"
  "spoyt/util"
  "path/filepath"
  "github.com/gin-gonic/gin"
)

func Web() {
  r := gin.Default()
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
    util.Converter(dst)

    // TODO
    // Progress Bar
    // Playlist link output (maybe generate and grab TL link directly?)
  })

  r.Run()
}
