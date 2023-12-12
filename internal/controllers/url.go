package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/codepnw/go-short-url/internal/constant"
	"github.com/codepnw/go-short-url/internal/database"
	"github.com/codepnw/go-short-url/internal/helper"
	"github.com/codepnw/go-short-url/internal/models"
	"github.com/gin-gonic/gin"
)

func ShortTheUrl(c *gin.Context) {
	var shortUrlBody models.ShortUrlBody
	err := c.BindJSON(&shortUrlBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": constant.BindError})
		return
	}

	code := helper.GenRandomString(6)

	record, _ := database.Mgr.GetUrlFromCode(code, constant.UrlCollection)

	if record.UrlCode != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": constant.ErrCodeInUse})
		return
	}

	var url models.UrlDB

	url.CreatedAt = time.Now().Unix()
	url.ExpiredAt = time.Now().Unix()
	url.UrlCode = code
	url.LongUrl = shortUrlBody.LongUrl
	url.ShortUrl = constant.BaseUrl + code

	resp, err := database.Mgr.Insert(url, constant.UrlCollection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "data": resp, "short_url": url.ShortUrl})
}

func RedirectURL(c *gin.Context) {
	code := c.Param("code")

	record, _ := database.Mgr.GetUrlFromCode(code, constant.UrlCollection)

	if record.UrlCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": constant.ErrNoURL})
		return
	}
	fmt.Println(record.LongUrl)

	c.Redirect(http.StatusPermanentRedirect, record.LongUrl)
}
