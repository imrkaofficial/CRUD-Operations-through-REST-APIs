package controller

import (
	"github.com/imrkaofficial/CRUD-Operations-through-REST-APIs/config"
	"github.com/imrkaofficial/CRUD-Operations-through-REST-APIs/models"

	"github.com/gin-gonic/gin"
)

func GetAcc(ctx *gin.Context) {
	var accounts []models.Account
	if err := config.DB.Find(&accounts).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve accounts"})
		return
	}
	ctx.JSON(200, accounts)
}

func CreateAcc(ctx *gin.Context) {
	var accounts models.Account
	if err := ctx.BindJSON(&accounts); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}
	if err := config.DB.Create(&accounts).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create account"})
		return
	}
	ctx.JSON(201, gin.H{"message": "Account has been created Successfully"})
}

func UpdateAcc(ctx *gin.Context) {
	uid := ctx.Param("uid")
	var accounts models.Account
	if err := config.DB.Where("uid = ?", uid).First(&accounts).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Account not found"})
		return
	}
	if err := ctx.BindJSON(&accounts); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}
	if err := config.DB.Save(&accounts).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update account"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Account has been updated Successfully"})
}

func DelAcc(ctx *gin.Context) {
	uid := ctx.Param("uid")
	var accounts models.Account
	if err := config.DB.Where("uid = ?", uid).Delete(&accounts).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete account"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Account has been deleted Successfully"})
}
