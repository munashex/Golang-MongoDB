package main 
import ( 
	"context" 
	"fmt" 
	"go.mongodb.org/mongo-driver/mongo" 
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
	
	if err:= client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
}
