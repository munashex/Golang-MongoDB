package main 
import ( 
	"context" 
	"fmt" 
	"go.mongodb.org/mongo-driver/mongo"      //packages
	"go.mongodb.org/mongo-driver/mongo/options" 
	"go.mongodb.org/mongo-driver/mongo/readpref" 
) 
func main() { 
	ctx := context.TODO() 
	opts := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(ctx, opts) 
	if err != nil { 
		panic(err) 
	}

	defer client.Disconnect(ctx) 
	fmt.Printf("%T\n", client)  
	
	//pinging the mongoDB server
	if err:= client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	
	//list the current databases in mongodb server 
	
	databases, err:= client.ListDatabaseNames(ctx, bson.M{}) 
	if err != nil {
		panic(err)
	}
	
	fmt.Println(databases)
	
}
