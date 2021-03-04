package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {		
        // OK 이면 index.html파일에 JSON데이터를 넘겨서 보여줌 
		c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Home Page",
			},
		)
	})

	r.Run()

}