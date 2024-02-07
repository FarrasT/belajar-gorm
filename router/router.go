package router

import (
	"belajar-gorm/controller"

	"github.com/gin-gonic/gin"

	_ "belajar-gorm/docs"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

// @title User and Product API
// @version 1.0
// @description This is a sample for managing user and product
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email farrastimorremboko@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/license/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func StartServer() *gin.Engine {
	router := gin.Default()

	// Create
	router.POST("/user", controller.CreateUser)

	// Update
	router.PUT("/user/:userId", controller.UpdateUserById)

	// Read
	router.GET("/user/:userId", controller.GetUserById)

	// Read
	router.GET("/user/product", controller.GetUserWithProducts)

	// Create
	router.POST("/product", controller.CreateProduct)

	// Delete
	router.DELETE("/product", controller.DeleteProductById)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
