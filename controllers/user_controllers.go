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

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	// _, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(201, gin.H{
		"age":      &User.Age,
		"email":    &User.Email,
		"id":       &User.ID,
		"username": &User.Username,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	// _, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password
	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "invalid email",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "invalid password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UpdateUsers(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)
	// _, _ = db, contentType
	User := models.User{}

	userId := int(userData["id"].(float64))

	reqId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "invalid parameter",
		})
		return
	}

	if reqId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "you have no access to edit this user",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err = db.Model(&User).Where("id = ?", userId).Updates(models.User{Email: User.Email, Username: User.Username}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}
	userResult := models.User{}
	err = db.First(&userResult, "id = ?", userId).Error

	c.JSON(201, gin.H{
		"id":         &userResult.ID,
		"email":      &userResult.Email,
		"username":   &userResult.Username,
		"age":        &userResult.Age,
		"updated_at": &userResult.UpdatedAt,
	})
}

func DeleteUsers(c *gin.Context) {
	db := database.GetDB()
	// contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)
	// _, _ = db, contentType
	User := models.User{}

	userId := int(userData["id"].(float64))

	reqId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err,
		})
		return
	}

	if reqId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "you have no access to edit this user",
		})
		return
	}

	Photo := models.Photo{}
	Comment := models.Comment{}
	SocialMedia := models.SocialMedia{}

	err = db.Model(&Comment).Where("user_id = ?", userId).Delete(&Comment).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": "failed deleting comment from this account",
			"errtype": err.Error(),
		})
		return
	}
	err = db.Model(&SocialMedia).Where("user_id = ?", userId).Delete(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": "failed deleting social media from this account",
			"errtype": err.Error(),
		})
		return
	}
	err = db.Model(&Photo).Where("user_id = ?", userId).Delete(&Photo).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": "failed deleting photo from this account : ",
			"errtype": err.Error(),
		})
		return
	}

	err = db.Model(&User).Where("id = ?", userId).Delete(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Your Account has been succesfully deleted",
	})
}
