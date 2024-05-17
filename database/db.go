package db

import (
	"context"
	"fmt"

	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitiateDb() {

	//get database login and password from env file
	err := godotenv.Load("data.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	// Getting and using a value from .env
	login := os.Getenv("DBLOGIN")
	password := os.Getenv("DBPASSWORD")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	// Construct the MongoDB URI string using fmt.Sprintf
	uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.8sck7u0.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0", login, password)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
