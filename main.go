package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	r.GET("/ckeditor_assets/pictures/:id/original_:file", func(c *gin.Context) {
		c.Redirect(301, "https://cdn.vashurok.ru/ckeditor_assets/pictures/"+c.Param("id")+"/content_"+c.Param("file"))
	})
	r.Run("127.0.0.1:13200")
}
