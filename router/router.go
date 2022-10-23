package router

import (
	"net/http"

	"twojsomsiad/controller"
	_ "twojsomsiad/docs"
	"twojsomsiad/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.GZIP())

	r.Use(middleware.CORS())

	// Controller
	api := controller.Controller{DB: db}

	// Authorization
	authorization := middleware.Auth(&api)
	// Authorization middleware for later use
	authMiddleware := authorization.MiddlewareFunc()

	// Swagger
	sg := r.Group("/")
	{
		sg.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		sg.GET("/swagger/", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/swagger-ui/index.html")
		})
	}

	// Remove CSP for testing
	// r.Use(middleware.CSP())

	// Endpoints
	auth := r.Group("/auth")
	{
		auth.POST("/login", authorization.LoginHandler)
		auth.GET("/refresh", authorization.RefreshHandler)
		auth.POST("/register", api.Register)
	}

	user := r.Group("/user")
	{
		user.GET("/:id", api.GetUser)
		protected := user.Group("/")
		{
			protected.Use(authMiddleware)
			protected.GET("/", api.GetMyUser)
			protected.POST("/", api.UpdateUser)
			protected.GET("/adverts", api.GetUserAdverts)
		}
	}

	advert := r.Group("/advert")
	{
		advert.GET("", api.Adverts)
		advert.GET("/:id", api.Advert)
		protected := advert.Group("/")
		{
			protected.Use(authMiddleware)
			protected.POST("/", api.CreateAdvert)
			protected.DELETE("/:id", api.RemoveAdvert)
			protected.GET("/:id/apply", api.Apply)
			protected.GET("/:id/application", api.GetApplications)
			protected.GET("/:id/application/:apid", api.VerifyApplication)
		}
	}

	return r
}
