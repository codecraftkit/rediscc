# rediscc

The `rediscc` package provides a simple and efficient way to manage multiple Redis connections in Go projects. This module abstracts the configuration and connection setup, allowing developers to focus on directly interacting with collections and documents.

#### Key Features:
* Enables the instantiation of multiple independent Redis connections.
* Simplifies connection management with configurable structures.
* Provides direct access to collections within a database.
* Fully compatible with standard Redis operations such as inserts, queries, updates, and deletions.

---
### Install

```bash
go get github.com/codecraftkit/rediscc
```

### Usage
Here’s a practical example of how to use the `rediscc` package:
```go
package main

import (
	"fmt"
	"context"
	"github.com/codecraftkit/rediscc"
)

func main() {

	MyRedisDbDataStore := rediscc.RedisDataStore{}
	
	if err := rediscc.Connect("mongodb://localhost:27017", "my_db", &MyRedisDbDataStore); err != nil {
		panic(err)
	}

	// Collection
	MyCollection := MyRedisDbDataStore.DB.Collection("users")

	type User struct {
		ID    string `bson:"_id,omitempty"`
		Email string `bson:"email"`
		Name  string `bson:"name"`
	}

	// Insert
	_, err := MyCollection.InsertOne(context.Background(), bson.M{"_id": "asdqwe123", "name": "John Doe", "email": "johndoe@example.com"})
	if err != nil {
		panic(err)
	}

	// Find
	var user User
	err = MyCollection.FindOne(context.Background(), bson.M{"_id": "asdqwe123"}).Decode(&user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
	
}
```
---

#### Why Use rediscc?
Modularity: Ideal for projects requiring multiple connections to different databases.
Ease of Use: Reduces the initial complexity of setting up Redis connections.
Seamless Integration: Compatible with the official Redis driver for Go (mongo-driver).

#### Best Suited For:
Developers seeking a straightforward solution to manage Redis connections in applications that need to efficiently and cleanly interact with multiple databases.






