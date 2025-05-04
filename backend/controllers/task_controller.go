package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"todoList/models"
	"todoList/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const DB_NAME = "tododatabase"
const COLLECTION = "tasks"

func CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Validar que el título no esté vacío
	if task.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El título es obligatorio"})
		return
	}
	// Validar que la descripción no esté vacía
	task.Completed = false
	task.CreatedAt = time.Now()

	// Obtener colección usando GetCollection
	collection := services.GetCollection(DB_NAME, COLLECTION)

	// Insertar en MongoDB con contexto controlado
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No se pudo guardar la tarea",
			"cause": err.Error(),
		})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusCreated, gin.H{
		"message": "Tarea creada exitosamente",
		"id":      res.InsertedID,
		"task":    task,
	})
}

func GetTasks(c *gin.Context) {
	collection := services.GetCollection(DB_NAME, COLLECTION)
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las tareas", "cause": err.Error()})
		return
	}
	defer cursor.Close(ctx)
	
	var tasks []models.Task
	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al decodificar una tarea", "cause": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}
	
	c.JSON(http.StatusOK, tasks)
	
}

func UpdateTask(c *gin.Context) {
	// Obtener el ID de la tarea desde los parámetros de la URL
	id := c.Param("id")

	// Convertir el ID a ObjectId
	objectID, err := primitive.ObjectIDFromHex(id)
	fmt.Println("ID:", id)	
	fmt.Println("ObjectID:", objectID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Obtener los datos enviados en el cuerpo de la solicitud
	var taskUpdate bson.M
	if err := c.ShouldBindJSON(&taskUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Obtener colección usando GetCollection
	collection := services.GetCollection(DB_NAME, COLLECTION)

	// Actualizar en MongoDB con contexto controlado
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": taskUpdate}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No se pudo actualizar la tarea",
			"cause": err.Error(),
		})
		return
	}

	if result.ModifiedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tarea actualizada exitosamente"})
}


func CompleteTask(c *gin.Context) {
	// Obtener el ID de la tarea desde los parámetros de la URL
	id := c.Param("id")

	// Convertir el ID a ObjectId
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Obtener colección usando GetCollection
	collection := services.GetCollection(DB_NAME, COLLECTION)

	// Actualizar en MongoDB con contexto controlado
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No se pudo completar la tarea",
			"cause": err.Error(),
		})
		return
	}

	if result.ModifiedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tarea completada exitosamente"})
}

func DeleteTask(c *gin.Context) {
	// Obtener el ID de la tarea desde los parámetros de la URL
	id := c.Param("id")

	// Convertir el ID a ObjectId
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Obtener colección usando GetCollection
	collection := services.GetCollection(DB_NAME, COLLECTION)
	// Eliminar en MongoDB con contexto controlado
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": objectID}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No se pudo eliminar la tarea",
			"cause": err.Error(),
		})
		return
	}
	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tarea eliminada exitosamente"})
}