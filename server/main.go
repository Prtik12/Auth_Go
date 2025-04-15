package main

import (
	"log"

	"github.com/Prtik12/auth_Golang/controllers"
	"github.com/Prtik12/auth_Golang/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := models.InitPrisma(); err != nil {
		log.Fatalf("Failed to initialize Prisma: %v", err)
	}

	defer models.PrismaClient.Prisma.Disconnect()

	r := gin.Default()

	store := cookie.NewStore([]byte("secret-session-key"))
	r.Use(sessions.Sessions("mySession", store))

	r.POST("/signup", controllers.SignUp)
	r.POST("/signin", controllers.SignIn)
	r.POST("/signout", controllers.SignOut)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
