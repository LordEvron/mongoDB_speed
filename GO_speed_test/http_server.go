package main

import (
	"context"
	"fmt"
	"github.com/goombaio/namegenerator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/rand"
	"os"
	"time"
)
 const SIZE  int = 500000
type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	fmt.Println("Connecting to MongoDB!")
	clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:37017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	collection := client.Database("mydb").Collection("persons")
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	fmt.Println("Generating Dataset")
	dataset := createdataset(SIZE)
	dataset2 := createdataset(SIZE)
	dataset3 := createdataset(SIZE)
	insertsingles(dataset, collection, false)
	insertsingles(dataset2, collection, true)
	insertmany(dataset3, collection)
	findOnetentimes(collection)

}



 func createdataset(num int) [SIZE]Person{
	 seed := time.Now().UTC().UnixNano()
	 nameGenerator := namegenerator.NewNameGenerator(seed)
	 var data [SIZE]Person
	 r := rand.New(rand.NewSource(99))
	 for i := 0; i < num; i++ {
		 data[i] = Person{nameGenerator.Generate(), r.Intn(100), nameGenerator.Generate()}
	 }
	 fmt.Println("Done.. Example of data generated")
	 fmt.Println(data[SIZE-10])
 	return data;
 }
func insertsingles(dataset [SIZE]Person, collection *mongo.Collection, parallel bool){
	fmt.Println("Starting inserting single with parallel=", parallel)
	start := time.Now()
	for i := 0; i < SIZE; i++ {
		if parallel {
			go singleinsert(dataset[i], collection)
		} else {
			singleinsert(dataset[i], collection)
		}

		//fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
	}

	stop := time.Now()
	diff :=stop.Sub(start)

	fmt.Println("Inserted ", SIZE, " single entries in DB, parallel ", parallel, "--- Duration: ", diff.Seconds())

}
 func singleinsert(data Person, collection *mongo.Collection){
	 _, err :=  collection.InsertOne(context.TODO(), data)
	 if err != nil {
		 log.Fatal(err)
	 }
}

func insertmany(dataset[SIZE]Person, collection *mongo.Collection) {
	var ppl []interface{}
	for i := 0; i < SIZE; i++{
		ppl= append(ppl, dataset[i])
	}
	fmt.Println("Starting inserting multi")
	start := time.Now()
	_, err := collection.InsertMany(context.TODO(), ppl)
	if err != nil {
		log.Fatal(err)
	}
	stop := time.Now()
	diff :=stop.Sub(start)

	fmt.Println("Inserted ", SIZE, " bulk entries in DB,--- Duration: ", diff.Seconds())

}

 func findOnetentimes(collection *mongo.Collection){

	 var result Person
	 filter := bson.D{{"age", 50}}
	 start := time.Now()
	 for i:=0;i<10;i++{
		 err := collection.FindOne(context.TODO(), filter).Decode(&result)
		 if err != nil {
			 log.Fatal(err)
		 }
		 fmt.Printf(" document: %+v\n", result)
	 }

	 stop := time.Now()
	 diff:=stop.Sub(start)
	 fmt.Println("Finished foundOne()x10 -- Duration: ", diff.Seconds())

	 start = time.Now()
	 res, err := collection.CountDocuments(context.TODO(), filter)
	 if err != nil {
	 	log.Fatal(err)
	 }
	 fmt.Printf("Count Documents in collection: %+v\n", res )
	 stop = time.Now()
	 diff=stop.Sub(start)
	 fmt.Println("Finished count_documents(filter), ", "--- Duration: ", diff.Seconds())


	 fmt.Println("Selecting ALL the documents that match the filer and printing the first 10" )
	 start = time.Now()
	 found , err := collection.Find(context.TODO(), filter)
	 if err != nil {
		 log.Fatal(err)
	 } else {
		 i := 0
		 for found.Next(context.TODO()) {
			 // Declare a result BSON object
			 i++
			 if i > 10 {
				 break
			 }
			 var result bson.M
			 err := found.Decode(&result)
			 // If there is a cursor.Decode error
			 if err != nil {
				 fmt.Println("cursor.Next() error:", err)
				 os.Exit(1)

				 // If there are no cursor.Decode errors
			 } else {
				 fmt.Println("result:", result)
			 }
		 }
	 }
	 stop = time.Now()
	 diff=stop.Sub(start)
	 fmt.Println("Finished find(filter)--- Duration: ", diff.Seconds())

 }