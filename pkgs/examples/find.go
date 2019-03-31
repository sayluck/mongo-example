package examples

import (
	"github.com/sayluck/mongo-example/pkgs/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

func EgFindByID() (interface{}, error) {
	mgo := Connect()
	var infio interface{}

	oID, err := mongodb.GenerateID("5ca037f278446ee5f602ffcc")
	if err != nil {
		return nil, err
	}
	ret := mgo.FindByID(oID, &infio)
	if ret.Error != nil {
		return nil, ret.Error
	}

	return infio, nil
}

func EgFind() (interface{}, error) {
	mgo := Connect()
	var infio interface{}

	filter := bson.D{} //{"name", "xiao min"}
	ret := mgo.Find(filter, &infio)
	if ret.Error != nil {
		return nil, ret.Error
	}

	return infio, nil
}
