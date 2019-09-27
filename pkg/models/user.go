package models

import (
	"context"
	"github.com/Blockchainpartner/hagentrust/pkg/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// A User is a user of the app, with his own credentials and token balance
type User struct {
	ID               *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	ClientID         *primitive.ObjectID `json:"-" bson:"clientID"`
	OnChainDeployed  *bool               `json:"onChainDeployed" bson:"onChainDeployed"`
	Name        *string             `json:"name" bson:"name"`
	Email            *string             `json:"email" bson:"email"`
	Address          *string             `json:"address" bson:"address"`
}

// UserFilter is the filter type on User
type UserFilter struct {
	ID               *primitive.ObjectID    `json:"-" bson:"_id,omitempty" csv:"-"`
	ClientID         *primitive.ObjectID    `json:"-" bson:"clientID,omitempty" csv:"-"`
	OnChainDeployed  *bool                  `json:"onChainDeployed,omitempty" bson:"onChainDeployed,omitempty"`
	Name        *string                `json:"name,omitempty" bson:"name,omitempty" csv:"name"`
	Email            *string                `json:"email,omitempty" bson:"email,omitempty" csv:"email"`
	Address          *string                `json:"address,omitempty" bson:"address,omitempty" csv:"-"`
}

// Post pushes a new User to the DB
func (u User) Post() error {
	// store user in the db
	_, err := db.UserCollection.InsertOne(
		context.Background(),
		u,
	)
	return err
}