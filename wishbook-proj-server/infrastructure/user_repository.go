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

// UserRepository is interface for interaction with MongoDB
type UserRepository interface {
	InsertUser(ctx context.Context, user domain.User) (id string, err error)
	GetUserByID(ctx context.Context, id string) (userResult domain.User, err error)
	GetAllUsers(ctx context.Context) (usersResults []*domain.User, err error)
	GetUserByLogin(ctx context.Context, login string) (userResult domain.User, err error)
	RemoveUserByID(ctx context.Context, id string) (deleted int64, err error)
	UpdateUser(ctx context.Context, user domain.User) (updated int64, err error)
}

type userRepository struct{}

// NewUserRepository returns a pointer instance structure
func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (userRepository) InsertUser(ctx context.Context, user domain.User) (id string, err error) {

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
	collection := client.Database(dbName).Collection(userCollection)

	var um userMongo
	userDomainToMongoConverter(user, &um)

	// Reset ID - should be autogenerated
	um.ID = primitive.NewObjectID()

	// Collection insert one
	insertResult, err := collection.InsertOne(ctx, um)
	if err != nil {
		return id, err
	}

	id = insertResult.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (userRepository) GetUserByID(ctx context.Context, id string) (userResult domain.User, err error) {

	client, err := newClient(ctx)
	if err != nil {
		return userResult, err
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
	collection := client.Database(dbName).Collection(userCollection)

	// Create filter
	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return userResult, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: mongoID}}

	var um userMongo

	err = collection.FindOne(ctx, filter).Decode(&um)
	if err != nil {
		return userResult, err
	}

	userMongoToDomainConverter(um, &userResult)

	return userResult, nil
}

func (userRepository) GetAllUsers(ctx context.Context) (usersResults []*domain.User, err error) {

	client, err := newClient(ctx)
	if err != nil {
		return usersResults, err
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
	collection := client.Database(dbName).Collection(userCollection)

	// Get many
	options := options.Find()
	filter := bson.M{}

	cur, err := collection.Find(ctx, filter, options)
	if err != nil {
		return usersResults, err
	}

	// Finding multiple documents returns a cursor
	// iterating through the cursor allows us to decode documents one at a time
	for cur.Next(ctx) {
		var elemMongo userMongo
		var elem domain.User
		err := cur.Decode(&elemMongo)
		if err != nil {
			return usersResults, err
		}
		userMongoToDomainConverter(elemMongo, &elem)
		usersResults = append(usersResults, &elem)
	}

	if err := cur.Err(); err != nil {
		return usersResults, err
	}

	// Close the cursor once finished
	cur.Close(ctx)

	return usersResults, nil
}

func (userRepository) GetUserByLogin(ctx context.Context, login string) (userResult domain.User, err error) {

	if login == "" {
		return userResult, nil
	}

	client, err := newClient(ctx)
	if err != nil {
		return userResult, err
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
	collection := client.Database(dbName).Collection(userCollection)

	// Create filter
	filter := bson.D{primitive.E{Key: "login", Value: login}}

	var um userMongo

	err = collection.FindOne(ctx, filter).Decode(&um)
	if err != nil {
		return userResult, err
	}

	userMongoToDomainConverter(um, &userResult)

	return userResult, nil
}

func (userRepository) RemoveUserByID(ctx context.Context, id string) (deleted int64, err error) {

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

	// Create collection
	collection := client.Database(dbName).Collection(userCollection)

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

func (userRepository) UpdateUser(ctx context.Context, user domain.User) (updated int64, err error) {

	var um userMongo
	err = userDomainToMongoConverter(user, &um)
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
	collection := client.Database(dbName).Collection(userCollection)

	// Login is not updatable
	updFilter := bson.D{primitive.E{Key: "_id", Value: um.ID}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "name", Value: um.Name},
			primitive.E{Key: "surname", Value: um.Surname},
			primitive.E{Key: "passwordhash", Value: um.PasswordHash},
			primitive.E{Key: "role", Value: um.Role},
		}},
	}

	updateResult, err := collection.UpdateOne(ctx, updFilter, update)
	if err != nil {
		return updated, err
	}

	return updateResult.ModifiedCount, nil
}
