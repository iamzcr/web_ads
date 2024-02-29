package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"sync"
)

var (
	client     *mongo.Client
	collection *mongo.Collection
	wg         sync.WaitGroup
	done       chan bool
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://192.168.18.93:25001")
	var err error
	client, err = mongo.Connect(nil, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("web_ads")
	collection = database.Collection("click_log")

	http.HandleFunc("/insert", handleInsert)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleInsert(w http.ResponseWriter, r *http.Request) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		document := bson.D{
			{Key: "name", Value: "John"},
			{Key: "age", Value: 30},
		}
		_, err := collection.InsertOne(nil, document)
		if err != nil {
			log.Println(err)
		}
		done <- true
	}()
	fmt.Fprint(w, "Insert request received")

	select {
	case <-done:
		// 插入操作已完成
	default:
		// 超时或其他处理
	}
}
