package lib

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Client = &mongo.Client{}
var Coll *mongo.Collection

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://rootuser:rootpass@localhost:27017"))
	if err != nil {
		log.Fatalln("connection failed:", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}
	Client = client
	Coll = Client.Database("lol-infos").Collection("champions")
}

/*CreateDocument function inserts hero information to MongoDB
This snippet can be used in main func to re-fetch whole heroes

	heroes := lib.ExcelToSlice()
	for _, hero := range heroes {
		time.Sleep(500 * time.Millisecond)
		getHero := lib.GetHeroInfo(hero)
		if !lib.CreateDocument(getHero) {
			fmt.Printf("%s couldn't be added", hero)
			break
		}
	}
*/
func CreateDocument(hero *HeroInfoStruct) bool {
	document, err := bson.Marshal(hero)
	if err != nil {
		fmt.Println("bson marshalling failed:", err)
		return false
	}
	ctx, cancel := context.WithTimeout(Ctx, 3*time.Second)
	defer cancel()
	_, err = Coll.InsertOne(ctx, document)
	if err != nil {
		fmt.Println("insert to collection failed:", err)
		return false
	}
	return true
}

//RetrieveDocument function finds given hero informations from related collection and returns
func RetrieveDocument(hero string) *HeroInfoStruct {
	ctx, cancel := context.WithTimeout(Ctx, 3*time.Second)
	defer cancel()
	filter := bson.D{{"id", hero}}
	res := Coll.FindOne(ctx, filter)
	heroInfo := HeroInfoStruct{}
	err := res.Decode(&heroInfo)
	if err != nil {
		fmt.Println("decode failed:", err)
		return nil
	}

	return &heroInfo
}
