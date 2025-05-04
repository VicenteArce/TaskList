package routes

import (
    "github.com/gin-gonic/gin"
    "todoList/controllers"
)

func RegisterTaskRoutes(router *gin.Engine) {
    // Agrupamos las rutas relacionadas a tareas bajo /tasks
    tasks := router.Group("/tasks")
    {
        tasks.POST("/", controllers.CreateTask)         // Crear tarea
        tasks.GET("/", controllers.GetTasks)            // Obtener todas las tareas
        //tasks.GET("/:id", controllers.GetTask)        // Obtener tarea por ID
        tasks.PUT("/:id", controllers.UpdateTask)     // Actualizar tarea por ID
        tasks.PUT("/:id/complete", controllers.CompleteTask) // Completar tarea por ID
        tasks.DELETE("/:id", controllers.DeleteTask)  // Eliminar tarea por ID
    }
}
