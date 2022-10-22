package controllers

import (
	"final/database"
	"final/helpers"
	"final/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userId

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               &SocialMedia.ID,
		"name":             &SocialMedia.Name,
		"social_media_url": &SocialMedia.SocialMediaUrl,
		"user_id":          &SocialMedia.UserID,
		"created_at":       &SocialMedia.CreatedAt,
	})
}

func ShowSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := []models.SocialMedia{}
	respSocialMedia := []models.SocialMediaGetResponse{}
	err := db.Find(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": "social media data cannot accessed",
		})
	}
	User := models.User{}

	for _, v := range SocialMedia {
		db.First(&User, "id = ?", v.UserID)
		input := models.SocialMediaGetResponse{}
		input.SocialMedia.ID = v.ID
		input.SocialMedia.Name = v.Name
		input.SocialMedia.SocialMediaUrl = v.SocialMediaUrl
		input.SocialMedia.UserID = v.UserID
		input.SocialMedia.CreatedAt = v.CreatedAt
		input.SocialMedia.UpdatedAt = v.UpdatedAt
		input.SocialMedia.User.ID = User.ID
		input.SocialMedia.User.Username = User.Username
		respSocialMedia = append(respSocialMedia, input)
	}
	c.JSON(http.StatusOK, respSocialMedia)
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}

	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "invalid parameter",
		})
		return
	}

	socialMediaResult := models.SocialMedia{}
	err = db.First(&socialMediaResult, "id = ?", socialMediaId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "bad request",
			"message": "not found",
		})
		return
	}
	if socialMediaResult.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "you have no access to edit this Social Media",
		})
		return
	}
	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err = db.Model(&SocialMedia).Where("id = ?", socialMediaId).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaUrl: SocialMedia.SocialMediaUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	err = db.First(&socialMediaResult, "id = ?", socialMediaId).Error
	c.JSON(http.StatusOK, gin.H{
		"id":               &socialMediaResult.ID,
		"name":             &socialMediaResult.Name,
		"social_media_url": &socialMediaResult.SocialMediaUrl,
		"user_id":          &socialMediaResult.UserID,
		"updated_at":       &socialMediaResult.UpdatedAt,
	})
}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))

	SocialMedia := models.SocialMedia{}

	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "invalid parameter",
		})
		return
	}

	socialMediaResult := models.SocialMedia{}
	err = db.First(&socialMediaResult, "id = ?", socialMediaId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "bad request",
			"message": "not found",
		})
		return
	}
	if socialMediaResult.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "you have no access to edit this Social Media",
		})
		return
	}

	err = db.Model(&SocialMedia).Where("id = ?", socialMediaId).Delete(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(201, gin.H{
		"message": "Your Social Media has been succesfully deleted",
	})
}
