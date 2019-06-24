package infrastructure

import (
	"wishbook/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userMongo struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string
	Surname      string
	Login        string
	PasswordHash string
	Role         int
}

func userMongoToDomainConverter(um userMongo, u *domain.User) {
	u.ID = um.ID.Hex()
	u.Name = um.Name
	u.Surname = um.Surname
	u.Login = um.Login
	u.PasswordHash = um.PasswordHash
	u.Role = domain.UserRole(um.Role)
}

func userDomainToMongoConverter(u domain.User, um *userMongo) error {
	if u.ID != "" {
		ID, err := primitive.ObjectIDFromHex(u.ID)
		if err != nil {
			return err // "ObjectID is not convertable"
		}
		um.ID = ID
	}

	um.Name = u.Name
	um.Surname = u.Surname
	um.Login = u.Login
	um.PasswordHash = u.PasswordHash
	um.Role = int(u.Role)

	return nil
}
