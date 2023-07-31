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

	// Load environment variables from the .env file
	godotenv.Load()
	port := os.Getenv("PORT")

	r := gin.Default()

	// Conexion ala base de datos
	h := &handlers.Handler{}
	h.ConnectDB()
	// Enable CORS middleware with permissive configuration
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	// routes Auth
	r.POST("/login", h.LoginHandler)
	r.POST("/register", h.RegisterHandler)

	//Agregando el Middlware token
	//r.Use(middlewares.AuthMiddleware)
	// routes Bills
	r.GET("/bills", h.BillsHandler)
	r.POST("/bills", h.NewBillHandler)
	r.GET("/bills/:id", h.GetBillHandler)
	r.DELETE("/bills/:id", h.DeleteBillHandler)

	// routes Users
	r.GET("/users", h.UsersHandler)
	r.POST("/users", h.NewUserHandler)
	r.GET("/users/:id", h.GetUserHandler)
	r.PUT("/users/:id", h.UpdateUser)
	r.DELETE("/users/:id", h.DeleteUserHandler)

	r.Run(fmt.Sprintf("0.0.0.0:%s", port))
}
