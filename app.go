package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

/*
curl -X POST http://localhost:8080/upload \
  -F "upload[]=@/Users/appleboy/test1.zip" \
  -F "upload[]=@/Users/appleboy/test2.zip" \
  -H "Content-Type: multipart/form-data"
*/
func saveFile(c *gin.Context, dst string) error {
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	for _, file := range files {
		if err := c.SaveUploadedFile(file, filepath.Join(dst, file.Filename)); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	rootPath := os.Getenv("ROOT_PATH")
	port := os.Getenv("PORT")

	if _, err = os.Stat(rootPath); err != nil {
		log.Fatal(err)
	}
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		if err := saveFile(c, rootPath); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("files uploaded!"))
	})
	router.POST("/upload/:name", func(c *gin.Context) {
		dstDir := filepath.Join(rootPath, c.Param("name"))
		if _, err = os.Stat(dstDir); err != nil {
			log.Warnf("%s not exists, creating now", dstDir)
			if err := os.Mkdir(dstDir, 0775); err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
		}
		if err := saveFile(c, dstDir); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("files success uploaded!"))
	})
	router.Run(port)
}
