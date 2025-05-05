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
		categoryRoutes.GET("/", categoryController.AllCategory)
		categoryRoutes.GET("/:id", categoryController.CategoryById)
		categoryRoutes.Use(middlewares.AuthMiddleware())
		categoryRoutes.POST("/", categoryController.CreateCategory)
		categoryRoutes.PUT("/:id", categoryController.UpdateCategory)
		categoryRoutes.DELETE("/:id", categoryController.DeleteCategory)
	}

	newsController := controllers.NewNewsController(db)
	newsRoutes := router.Group("/news")
	{
		newsRoutes.GET("/", newsController.AllNews)
		newsRoutes.GET("/:id", newsController.NewsById)
		newsRoutes.Use(middlewares.AuthMiddleware())
		newsRoutes.POST("/", newsController.CreateNews)
		newsRoutes.PUT("/:id", newsController.UpdateNews)
		newsRoutes.DELETE("/:id", newsController.DeleteNews)
	}

	
	commentController := controllers.NewCommentController(db)
	commentRoutes := router.Group("/news/:id/comments")
	{
		commentRoutes.POST("/", commentController.CreateComment)
	}

	customPageController := controllers.NewCustomPageController(db)
	pageRoutes := router.Group("/pages")
	{
		pageRoutes.GET("/", customPageController.AllCustomPages)
		pageRoutes.GET("/:id", customPageController.PageById)
		pageRoutes.Use(middlewares.AuthMiddleware())
		pageRoutes.POST("/", customPageController.CreateCustomPage)
		pageRoutes.PUT("/:id", customPageController.UpdateCustomPage)
		pageRoutes.DELETE("/:id", customPageController.DeleteCustomPage)
	}

	return router
}
