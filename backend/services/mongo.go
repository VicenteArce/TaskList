package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func InitMongo() {
	// Cargar las variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		log.Println("Advertencia: no se pudo cargar .env, se usará configuración del entorno.")
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("La variable de entorno MONGO_URI no está definida")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	fmt.Println("Conectado a MongoDB")
}

// GetCollection permite obtener cualquier colección especificando nombre de base de datos y colección
func GetCollection(database string, collection string) *mongo.Collection {
	return Client.Database(database).Collection(collection)
}
