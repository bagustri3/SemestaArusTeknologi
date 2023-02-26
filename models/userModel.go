package models

import(
"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id"`
	Username	*string `json:"username" validate:"required", min=2, max=64"`
	Password	*string `json:"password" validate:"required,min=8`
}
