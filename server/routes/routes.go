package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/maulerrr/book-addict-server/server/controllers"
	"github.com/maulerrr/book-addict-server/server/middlewares"
)

func InitRoutes(app *gin.Engine) {
	router := app.Group("api/v1")

	authRouter := router.Group("/auth")
	authRouter.POST("/login", controllers.Login)
	authRouter.POST("/signup", controllers.Signup)

	usersRouter := router.Group("/users")
	usersRouter.GET("/", controllers.GetAllUsers)
	usersRouter.GET("/:id", controllers.GetUserById)
	usersRouter.POST("/", controllers.CreateUser)
	usersRouter.DELETE("/:id", controllers.DeleteUserById)

	readListRouter := router.Group("/readlist")
	readListRouter.GET("/", middlewares.AuthMiddleware(), controllers.GetAllBooks)
	readListRouter.POST("/", middlewares.AuthMiddleware(), controllers.AddBook)
	readListRouter.GET("/:id", middlewares.AuthMiddleware(), controllers.GetBookById)
	readListRouter.DELETE("/:id", middlewares.AuthMiddleware(), controllers.DeleteBookById)

	bookTabRouter := router.Group("/booktabs")
	bookTabRouter.GET("/", middlewares.AuthMiddleware(), controllers.GetAllBookTabs)
	bookTabRouter.GET("/favorites/:id", middlewares.AuthMiddleware(), controllers.GetFavorites)
	bookTabRouter.GET("/finished/:id", middlewares.AuthMiddleware(), controllers.GetFinishedBooks)
	bookTabRouter.POST("/add", middlewares.AuthMiddleware(), controllers.AddBookTab)
	bookTabRouter.DELETE("/delete/:id", middlewares.AuthMiddleware(), controllers.DeleteBookTabById)
	bookTabRouter.PUT("/update", middlewares.AuthMiddleware(), controllers.UpdateBookTab)
}
