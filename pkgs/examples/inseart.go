package examples

import (
	"go.mongodb.org/mongo-driver/bson"
)

type user struct {
	Name string `json:"name",bson:"name"`
	ID   string `json:"id",bson:"id"`
	Sex  int    `json:"sex",bson:"sex"`
}

// eg 1 inseart one: use struct
func EgInseartOneWithStruct() error {
	mgo := Connect()
	doc := &user{
		Name: "xiao min",
		ID:   "0123456789",
		Sex:  1,
	}
	ret := mgo.InseartOne(doc)
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}

// eg 2 inseart one: use bson.M
func EgInseartOneWithBson() error {
	mgo := Connect()
	doc := bson.M{"name": "xiao hong", "id": "999993", "Sex": 2}
	ret := mgo.InseartOne(doc)
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}
