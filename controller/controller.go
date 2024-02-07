package controller

import (
	"belajar-gorm/database"
	"belajar-gorm/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserReq struct {
	Email string `json:"email"`
}

type ProductReq struct {
	Brand string `json:"brand"`
	Name  string `json:"name"`
}

type UserRes struct {
	Message string                 `json:"message"`
	Data    map[string]models.User `json:"data"`
}

// CreateUser godoc
// @Summary Create new user
// @Description Create new user
// @Tag users
// @Accept json
// @Produce json
// @Param data body UserReq true "create new user"
// @Success 200 {object} models.User
// @Router /user [post]
func CreateUser(ctx *gin.Context) {
	var newUser UserReq

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	db := database.GetDB()

	User := models.User{
		Email: newUser.Email,
	}

	if err := db.Create(&User).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error_status": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data": map[string]interface{}{
			"user": newUser,
		},
	})
}

// GetUserById godoc
// @Summary Get User by Id
// @Description Get User by Id
// @Tag users
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /user [post]
func GetUserById(ctx *gin.Context) {
	userId := ctx.Param("userId")

	db := database.GetDB()

	user := models.User{}

	err := db.First(&user, "id = ?", userId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error_status":  "Data Not Found",
				"error_message": fmt.Sprintf("User with id %v not found", userId),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error_status": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": map[string]interface{}{
			"user": user,
		},
	})
}

func UpdateUserById(ctx *gin.Context) {
	userId := ctx.Param("userId")

	var updateUser UserReq

	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	db := database.GetDB()

	user := models.User{}

	err := db.Model(&user).Where("id = ?", userId).Updates(models.User{
		Email: updateUser.Email}).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error_status": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": map[string]interface{}{
			"user": user,
		},
	})
}

func CreateProduct(ctx *gin.Context) {
	userId := ctx.Param("userId")

	db := database.GetDB()

	var newProduct ProductReq

	if err := ctx.ShouldBindJSON(&newProduct); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(userId)
	Product := models.Product{
		UserID: uint(id),
		Brand:  newProduct.Brand,
		Name:   newProduct.Name,
	}

	err := db.Create(&Product).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error_status": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data": map[string]interface{}{
			"product": newProduct,
		},
	})
}

func GetUserWithProducts(ctx *gin.Context) {
	db := database.GetDB()

	users := models.User{}
	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error_status": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": map[string]interface{}{
			"user": users,
		},
	})
}

func DeleteProductById(ctx *gin.Context) {
	userId := ctx.Param("productId")

	db := database.GetDB()

	product := models.Product{}
	err := db.Where("id = ?", userId).Delete(&product).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error_status": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": map[string]interface{}{
			"product": product,
		},
	})
}
