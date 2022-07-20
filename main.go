package main 
import ( 
	"context" 
	"fmt" 
	"go.mongodb.org/mongo-driver/mongo" 
	"go.mongodb.org/mongo-driver/mongo/options" 
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
}
