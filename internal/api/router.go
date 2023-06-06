package api

import (
	"filemanger/internal/api/v1/auth"
	"filemanger/internal/api/v1/fileapi"
	"filemanger/internal/api/v1/userapi"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		authRouter := r.Group("/auth")
		authRouter.POST("/login", auth.LoginHandler)
		authRouter.POST("/logout", auth.LogoutHandler)

		// v1.Use(auth.AuthMiddleware())
		userGroup := v1.Group("/users")
		{
			userGroup.GET("/:id", userapi.GetUser)
			userGroup.POST("", userapi.AddUser)
			userGroup.PUT("/:id", userapi.UpdateUser)
			userGroup.DELETE("/:id", userapi.DeleteUser)
		}

		fileGroup := v1.Group("/files")
		{
			fileGroup.GET("/:id", fileapi.DeleteFile) // WARN: Edit later
			fileGroup.POST("/upload", fileapi.UploadFile)
			fileGroup.DELETE("/:id", fileapi.DeleteFile)
			fileGroup.POST("/file",fileapi.GetFolderByID)
		}
		folderGroup := v1.Group("/folders")
		{
			folderGroup.POST("/create",fileapi.CreateFolder)
			folderGroup.GET("/:id",fileapi.GetFolderByID)

		}
	}

	r.Run(":8080")
}
