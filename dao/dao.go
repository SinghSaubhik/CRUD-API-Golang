package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
	"todo-app/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

import . "todo-app/models"

var client *mongo.Client

func getCollection() (*mongo.Collection, context.Context, context.CancelFunc) {
	collection := client.Database(utils.DB_NAME).Collection(utils.COLLECTION_NAME)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return collection, ctx, cancel
}

func ConnectDB() (*mongo.Client, context.Context) {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic("Unable to create DB client")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic("Unable to connect to DB server")
	}

	log.Println("Connected to DB")
	return client, ctx
}

func InsertOne(todo Todo) error {
	collection, ctx, cancel := getCollection()
	defer cancel()
	_, err := collection.InsertOne(ctx, bson.M{"_id": todo.ID, "title": todo.Title, "completed": todo.Completed})
	if err != nil {
		return Error{}.NewError("Hello")
	}

	return nil
}

func FindAll() []Todo {
	todos := []Todo{}
	collection, ctx, cancel := getCollection()
	defer cancel()
	cur, _ := collection.Find(ctx, bson.D{})
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result Todo
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return todos
}

func FindOne(id string) (Todo, error) {
	var todo Todo
	collection, ctx, cancel := getCollection()
	defer cancel()

	filter := bson.D{{"_id", id}}
	err := collection.FindOne(ctx, filter).Decode(&todo)
	return todo, err
}

func FindOneAndDelete(id string) (Todo, error) {
	var deletedTodo Todo
	collection, ctx, cancel := getCollection()
	defer cancel()
	filter := bson.D{{"_id", id}}
	err := collection.FindOneAndDelete(ctx, filter).Decode(&deletedTodo)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return deletedTodo, err
		}
		log.Fatal(err)
	}

	return deletedTodo, nil
}

func FindOneAndUpdate(id string, newTodo Todo) (Todo, error) {
	var updatedDoc Todo
	collection, ctx, cancel := getCollection()
	defer cancel()
	filter := bson.D{{"_id", id}}
	//bson.D{{"$set", bson.D{{"email", "newemail@example.com"}}}}
	newDoc := bson.M{"$set": bson.M{"_id": id, "title": newTodo.Title, "completed": newTodo.Completed}}
	err := collection.FindOneAndUpdate(ctx, filter, newDoc).Decode(&updatedDoc)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return updatedDoc, err
		}
		log.Fatal(err)
	}
	return updatedDoc, nil
}
