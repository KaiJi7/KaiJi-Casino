package collection

import "go.mongodb.org/mongo-driver/bson/primitive"

type Strategy struct {
	Id          *primitive.ObjectID `json:"id" bson:"_id"`
	GamblerId   *primitive.ObjectID `json:"gambler_id" bson:"gambler_id"`
	Name        string              `json:"name" bson:"name"`
	Description string              `json:"description" bson:"description"`
}
