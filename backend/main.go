package main

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "todoList/routes"
    "todoList/services"
    "time"
)

func main() {
    // Inicializar conexión Mongo
    services.InitMongo()

    router := gin.Default()

    // Middleware CORS configurado correctamente
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    // Registrar rutas
    routes.RegisterTaskRoutes(router)

    // Iniciar servidor
    router.Run(":8080")
}
