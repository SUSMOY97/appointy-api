package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



type Article struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title" bson:"title,omitempty"`
	SubTitle string             `json:"subtitle" bson:"subtitle,omitempty"`
	Content  string             `json:"content" bson:"content,omitempty"`
	Creation time.Time          `json:"creation" bson:"creation,omitempty"`
}


var collection = ConnectDB()

func main() {
	
	r := mux.NewRouter()

	
	r.HandleFunc("/articles", getArticles).Methods("GET")
	r.HandleFunc("/articles", createArticle).Methods("POST")
	r.HandleFunc("/articles/{id}", getArticle).Methods("GET")
	r.HandleFunc("/articles/search?q=title", getArticle).Methods("GET")

	
	log.Fatal(http.ListenAndServe(":8000", r))
}
