package routes

import (
	"github.com/alghoziii/cms-backend-technical/controllers"
	"github.com/alghoziii/cms-backend-technical/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.GET("/docs/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1),
	))

	authController := controllers.NewAuthController(db)
	router.POST("/login", authController.Login)

	categoryController := controllers.NewCategoryController(db)
	categoryRoutes := router.Group("/categories")
	{
		categoryRoutes.GET("/", categoryController.FindCategories)
		categoryRoutes.GET("/:id", categoryController.FindCategoryById)
		categoryRoutes.Use(middlewares.AuthMiddleware())
		categoryRoutes.POST("/", categoryController.CreateCategory)
		categoryRoutes.PUT("/:id", categoryController.UpdateCategory)
		categoryRoutes.DELETE("/:id", categoryController.DeleteCategory)
	}

	newsController := controllers.NewNewsController(db)
	newsRoutes := router.Group("/news")
	{
		newsRoutes.GET("/", newsController.FindNews)
		newsRoutes.GET("/:id", newsController.FindNewsById)
		newsRoutes.Use(middlewares.AuthMiddleware())
		newsRoutes.POST("/", newsController.CreateNews)
		newsRoutes.PUT("/:id", newsController.UpdateNews)
		newsRoutes.DELETE("/:id", newsController.DeleteNews)
	}

	
	commentController := controllers.NewCommentController(db)
	commentRoutes := router.Group("/news/:id/comments")
	{
		commentRoutes.GET("/", commentController.FindCommentsByNewsId)
		commentRoutes.POST("/", commentController.CreateComment)
	}

	customPageController := controllers.NewCustomPageController(db)
	pageRoutes := router.Group("/pages")
	{
		pageRoutes.GET("/", customPageController.FindCustomPages)
		pageRoutes.GET("/:id", customPageController.FindCustomPageById)
		pageRoutes.Use(middlewares.AuthMiddleware())
		pageRoutes.POST("/", customPageController.CreateCustomPage)
		pageRoutes.PUT("/:id", customPageController.UpdateCustomPage)
		pageRoutes.DELETE("/:id", customPageController.DeleteCustomPage)
	}

	return router
}
