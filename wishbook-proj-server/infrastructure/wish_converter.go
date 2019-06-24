package infrastructure

import (
	"time"
	"wishbook/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type wishMongo struct {
	ID     primitive.ObjectID `bson:"_id"`
	Title  string
	UserID string
	Done   bool
	Date   time.Time
}

func wishMongoToDomainConverter(wm wishMongo, w *domain.Wish) {
	w.ID = wm.ID.Hex()
	w.Title = wm.Title
	w.UserID = wm.UserID
	w.Done = wm.Done
	w.Date = wm.Date
}

func wishDomainToMongoConverter(w domain.Wish, wm *wishMongo) error {
	if w.ID != "" {
		ID, err := primitive.ObjectIDFromHex(w.ID)
		if err != nil {
			return err // ObjectID is not convertable
		}
		wm.ID = ID
	}

	wm.Title = w.Title
	wm.UserID = w.UserID
	wm.Done = w.Done
	wm.Date = w.Date

	return nil
}
