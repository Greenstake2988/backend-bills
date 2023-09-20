package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// need install GCC compiler

	"backend-bills/handlers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	// Load environment variables from the .env file
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	// Conexion ala base de datos
	h := &handlers.Handler{}
	h.ConnectDB()

	// Enable CORS middleware with permissive configuration
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	// routes Auth
	r.POST("/login", h.Login)
	r.POST("/register", h.Register)
	r.POST("/verification", h.Verification)

	//Agregando el Middlware token
	//r.Use(middlewares.AuthMiddleware)
	// routes Bills
	r.GET("/bills", h.Bills)
	r.POST("/bills", h.NewBill)
	r.GET("/bills/:id", h.GetBill)
	r.DELETE("/bills/:id", h.DeleteBill)

	// routes Users
	r.GET("/users", h.Users)
	r.POST("/users", h.NewUser)
	r.GET("/users/:id", h.GetUser)
	r.PUT("/users/:id", h.UpdateUser)
	r.DELETE("/users/:id", h.DeleteUser)

	//services.SendVerificationCodeEmail("noelhpa@gmail.com")
	r.Run(fmt.Sprintf("0.0.0.0:%s", port))

}
