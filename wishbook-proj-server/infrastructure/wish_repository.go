package infrastructure

import (
	"context"
	"fmt"
	"log"

	"wishbook/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// WishRepository is interface for interaction with MongoDB
type WishRepository interface {
	InsertWish(ctx context.Context, wish domain.Wish) (id string, err error)
	GetWishByID(ctx context.Context, id string) (wishResult domain.Wish, err error)
	GetAllActualWishes(ctx context.Context) (wishesResults []*domain.Wish, err error)
	RemoveWishByID(ctx context.Context, id string) (deleted int64, err error)
	UpdateWish(ctx context.Context, wish domain.Wish) (updated int64, err error)
}

type wishRepository struct{}

// NewWishRepository returns a pointer instance structure
func NewWishRepository() WishRepository {
	return &wishRepository{}
}

func (wishRepository) InsertWish(ctx context.Context, wish domain.Wish) (id string, err error) {

	client, err := newClient(ctx)
	if err != nil {
		return id, err
	}

	defer func(ctx context.Context, client *mongo.Client) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection to MongoDB closed.")
	}(ctx, client)

	fmt.Println("Connection to MongoDB.")

	// Create collection
	collection := client.Database(dbName).Collection(wishCollection)

	var wm wishMongo
	wishDomainToMongoConverter(wish, &wm)

	// Reset ID - should be autogenerated
	wm.ID = primitive.NewObjectID()

	// Collection insert one
	insertResult, err := collection.InsertOne(ctx, wm)
	if err != nil {
		return id, err
	}

	id = insertResult.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (wishRepository) GetWishByID(ctx context.Context, id string) (wishResult domain.Wish, err error) {

	client, err := newClient(ctx)
	if err != nil {
		return wishResult, err
	}

	defer func(ctx context.Context, client *mongo.Client) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection to MongoDB closed.")
	}(ctx, client)

	fmt.Println("Connection to MongoDB.")

	// Create collection
	collection := client.Database(dbName).Collection(wishCollection)

	// Create filter
	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return wishResult, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: mongoID}}

	var wm wishMongo

	err = collection.FindOne(ctx, filter).Decode(&wm)
	if err != nil {
		return wishResult, err
	}

	wishMongoToDomainConverter(wm, &wishResult)

	return wishResult, nil
}

func (wishRepository) GetAllActualWishes(ctx context.Context) (wishesResult []*domain.Wish, err error) {

	client, err := newClient(ctx)
	if err != nil {
		return wishesResult, err
	}

	defer func(ctx context.Context, client *mongo.Client) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection to MongoDB closed.")
	}(ctx, client)

	fmt.Println("Connection to MongoDB.")

	// Create collection
	collection := client.Database(dbName).Collection(wishCollection)

	// Get many
	options := options.Find()
	filter := bson.D{primitive.E{Key: "done", Value: false}}

	cur, err := collection.Find(ctx, filter, options)
	if err != nil {
		return wishesResult, err
	}

	// Finding multiple documents returns a cursor
	// iterating through the cursor allows us to decode documents one at a time
	for cur.Next(ctx) {
		var elemMongo wishMongo
		var elem domain.Wish
		err := cur.Decode(&elemMongo)
		if err != nil {
			return wishesResult, err
		}
		wishMongoToDomainConverter(elemMongo, &elem)
		wishesResult = append(wishesResult, &elem)
	}

	if err := cur.Err(); err != nil {
		return wishesResult, err
	}

	// Close the cursor once finished
	cur.Close(ctx)

	return wishesResult, nil
}

func (wishRepository) RemoveWishByID(ctx context.Context, id string) (deleted int64, err error) {

	client, err := newClient(ctx)
	if err != nil {
		return deleted, err
	}

	defer func(ctx context.Context, client *mongo.Client) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection to MongoDB closed.")
	}(ctx, client)

	fmt.Println("Connection to MongoDB.")

	//Create collection
	collection := client.Database(dbName).Collection(wishCollection)

	// Remove
	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return deleted, err
	}

	delFilter := bson.D{primitive.E{Key: "_id", Value: mongoID}}
	deleteResult, err := collection.DeleteOne(ctx, delFilter)
	if err != nil {
		return deleted, err
	}

	deleted = deleteResult.DeletedCount

	return deleted, nil
}

func (wishRepository) UpdateWish(ctx context.Context, wish domain.Wish) (updated int64, err error) {

	var wm wishMongo
	err = wishDomainToMongoConverter(wish, &wm)
	if err != nil {
		return updated, err
	}

	client, err := newClient(ctx)
	if err != nil {
		return updated, err
	}

	defer func(ctx context.Context, client *mongo.Client) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection to MongoDB closed.")
	}(ctx, client)

	fmt.Println("Connection to MongoDB.")

	// Create collection
	collection := client.Database(dbName).Collection(wishCollection)

	updFilter := bson.D{primitive.E{Key: "_id", Value: wm.ID}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "done", Value: wm.Done},
		}},
	}

	updateResult, err := collection.UpdateOne(ctx, updFilter, update)
	if err != nil {
		return updated, err
	}

	return updateResult.ModifiedCount, nil
}
