package models

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoUri        = "mongodb://mongo:27017"
	mongoDB         = "stats"
	mongoCollection = "infor"
)

type StatsList struct {
	Name   string `bson:"name"`
	Number string `bson:"number"`
}

func Query() []StatsList {
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri).SetConnectTimeout(time.Second*5))
	if err != nil {
		log.Fatal(err)
	}

	database := mongoClient.Database(mongoDB)
	collection := database.Collection(mongoCollection)
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	defer cursor.Close(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	record := &StatsList{}
	list := make([]StatsList, 0)

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(record)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, *record)
	}
	return list
}

func HandleList() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {

		tmp := template.Must(template.ParseFiles("views/list.html"))
		tmp.ExecuteTemplate(response, "list.html", struct {
			Lists []StatsList
		}{Query()})
	})

}

func HandleAdd() {
	http.HandleFunc("/add/", func(response http.ResponseWriter, request *http.Request) {
		var (
			lists  StatsList
			errors = make(map[string]string)
		)

		if request.Method == http.MethodGet {

		} else if request.Method == http.MethodPost {
			name := request.PostFormValue("name")
			number := request.PostFormValue("number")

			lists = StatsList{
				Name:   name,
				Number: number,
			}

			mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri).SetConnectTimeout(time.Second*5))
			if err != nil {
				log.Fatal(err)
			}
			database := mongoClient.Database(mongoDB)
			collection := database.Collection(mongoCollection)
			_, err = collection.InsertOne(context.TODO(), lists)
			if err != nil {
				log.Fatal(err)
			}

			http.Redirect(response, request, "/", http.StatusFound)

		}

		tmp := template.Must(template.ParseFiles("views/add.html"))
		tmp.ExecuteTemplate(response, "add.html", struct {
			List   StatsList
			Errors map[string]string
		}{lists, errors})
	})
}
