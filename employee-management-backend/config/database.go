package config

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ConnectDB() *mongo.Client {
	// Ton URI MongoDB Atlas avec le mot de passe
	uri := "mongodb+srv://benrahosamira2:kaulZ4tpCCwNcNH0@cluster0.udrky.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Erreur lors de la création du client MongoDB : ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Erreur lors de la connexion à MongoDB : ", err)
	}

	log.Println("Connected to MongoDB Atlas!")
	return client
}

// Variable globale pour utiliser le client MongoDB
var DB *mongo.Client = ConnectDB()

// Retourne une collection MongoDB spécifique
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("employeeDB").Collection(collectionName)
}
