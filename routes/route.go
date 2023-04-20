package routes

import (
	"log"
	"mini-project-alterra/controllers"
	"mini-project-alterra/repositories"
	"mini-project-alterra/usecases"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func init() {
	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "Unauthorized"
}

func New(e *echo.Echo, db *gorm.DB) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtMiddleware := middleware.JWT([]byte(os.Getenv("SECRET_JWT")))

	userRepository := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepository)
	userController := controllers.NewUserController(userUsecase)

	socialMediaRepository := repositories.NewSocialMediaRepository(db)
	socialMediaUsecase := usecases.NewSocialMediaUsecase(socialMediaRepository)
	socialMediaController := controllers.NewSocialMediaController(userUsecase, socialMediaUsecase)

	photoRepository := repositories.NewPhotoRepository(db)
	photoUsecase := usecases.NewPhotoUsecase(photoRepository)
	photoController := controllers.NewPhotoController(userUsecase, photoUsecase)

	commentRepository := repositories.NewCommentRepository(db)
	commentUsecase := usecases.NewCommentUsecase(commentRepository, photoRepository)
	commentController := controllers.NewCommentController(userUsecase, commentUsecase)

	e.POST("/users/login", userController.SignIn)
	e.POST("/users/register", userController.SignUp)
	e.GET("/users", userController.GetCredential, jwtMiddleware)
	e.PATCH("/users", userController.UpdateUser, jwtMiddleware)
	e.DELETE("/users", userController.DeleteUser, jwtMiddleware)

	e.GET("/socialmedia", socialMediaController.GetMySocialMedia, jwtMiddleware)
	e.GET("/socialmedia/:id", socialMediaController.GetMySocialMediaByID, jwtMiddleware)
	e.GET("/socialmedias", socialMediaController.GetSocialMedias, jwtMiddleware)
	e.POST("/socialmedias", socialMediaController.CreateSocialMedia, jwtMiddleware)
	e.PATCH("/socialmedias/:socialMediaId", socialMediaController.UpdateSocialMedia, jwtMiddleware)
	e.DELETE("/socialmedias/:socialMediaId", socialMediaController.DeleteSocialMedia, jwtMiddleware)

	e.GET("/photo", photoController.GetMyPhoto, jwtMiddleware)
	e.GET("/photo/:id", photoController.GetMyPhotoByID, jwtMiddleware)
	e.GET("/photos", photoController.GetPhotos, jwtMiddleware)
	e.POST("/photos", photoController.CreatePhoto, jwtMiddleware)
	e.PATCH("/photos/:id", photoController.UpdatePhoto, jwtMiddleware)
	e.DELETE("/photos/:id", photoController.DeletePhoto, jwtMiddleware)

	e.GET("/comment", commentController.GetAllMyComment, jwtMiddleware)
	e.GET("/comment/:id", commentController.GetAllMyCommenByID, jwtMiddleware)
	e.GET("/comments", commentController.GetAllComments, jwtMiddleware)
	e.POST("/comments", commentController.CreateComment, jwtMiddleware)
	e.DELETE("/comments/:id", commentController.DeleteComment, jwtMiddleware)
	e.PATCH("/comments/:id", commentController.UpdateComment, jwtMiddleware)

}
