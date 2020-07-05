package main

import (
	. "./db"
	"net/http"
	"fmt"
	"path/filepath"
	"github.com/gin-gonic/gin"
	//. "./spider"
)


func setupRouter() *gin.Engine {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.Static("/public", "./public")
	router.GET("/api/user", func(c *gin.Context) {
	    c.JSON(http.StatusOK, gin.H{"user": "123"})
	})
	router.POST("/api/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		// Multipart form
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		files := form.File["files"]

		for _, file := range files {
			filename := filepath.Base(file.Filename)
			if err := c.SaveUploadedFile(file, filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
		}

		c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files with fields name=%s and email=%s.", len(files), name, email))
	})

    return router
}



func App() {

	err,engine:=Conn()
	if err!=nil{
		return
	}
	Test(engine)


	r:=setupRouter()
	r.Run(":8080")
}

func main() {
	//App()
	//GetGroupMsg()
}
