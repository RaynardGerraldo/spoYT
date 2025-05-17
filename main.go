package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.LoadHTMLFiles("templates/index.tmpl")
  r.GET("/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
        "title": "SpoYT",
    })
  })
  r.Run()
}
