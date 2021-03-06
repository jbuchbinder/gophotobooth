package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kennygrant/sanitize"
)

type stringarray []string

func initApi(m *gin.Engine) {
	g := m.Group("/api")
	g.GET("/photo/:batch/:slug", apiTakePhoto)
}

func apiTakePhoto(c *gin.Context) {
	if c.Param("batch") == "" {
		log.Printf("apiTakePhoto(): parameter 'batch' not found")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	batch := sanitize.BaseName(c.Param("batch"))
	if c.Param("slug") == "" {
		log.Printf("apiTakePhoto(): parameter 'slug' not found")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	slug := strings.Replace(sanitize.BaseName(c.Param("slug")), ":", "-", -1)

	err := CapturePhoto(*StoragePath, batch, slug)
	if err != nil {
		log.Print(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, true)
}
