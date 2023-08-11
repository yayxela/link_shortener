package models

import "go.mongodb.org/mongo-driver/bson/primitive"

const LinkCollection = "links"

type Link struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Base      string             `bson:"base" json:"base"`
	Short     string             `bson:"short" json:"short"`
	ValidFor  primitive.DateTime `bson:"validFor" json:"validFor"`
	CreatedAt primitive.DateTime `bson:"createdAt" json:"createdAt"`
}
