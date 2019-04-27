package examples

import (
	"fmt"
	"github.com/sayluck/mongo-example/pkgs/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
func EgFindWithCount() (interface{}, error) {
	var infio []interface{}

	ret := FindWithCount(infio)
	if ret.Error != nil {
		return nil, ret.Error
	}

	return infio, nil
}
func FindWithCount(outF []interface{}) *mongodb.MongoDB {
	mongoDB := Connect().PrepareDB()
	cur, err := mongoDB.Client.Database(mongoDB.Option.DBName).
		Collection(mongoDB.Option.Collection).Aggregate(mongoDB.Ctx, mongo.Pipeline{
		{
			{
				"$match",
				bson.M{
					"Sex": 1,
				},
			},
		},
		{
			{
				"$limit", 2,
			},
		},
		{
			{
				"$skip", 0,
			},
		},
		{
			{
				"$group", bson.M{
					"_id":   nil,
					"count": bson.M{"$sum": 1},
					"data":  bson.M{"$push": bson.M{"id": "$id", "sex": "$Sex", "name": "$name"}},
				},
			},
		},
	}, options.Aggregate())
	mongoDB.Error = err
	if err != nil {
		fmt.Println("error:", err)
		return mongoDB
	}
	defer cur.Close(mongoDB.Ctx)

	type outT struct {
		Count uint   `bson:count`
		Data  []user `bson:data`
	}

	var out outT
	for cur.Next(mongoDB.Ctx) {
		err := cur.Decode(&out)
		if err != nil {
			fmt.Printf("out err 0: %+v\n", err)
			return nil
		}
	}
	// out detail: {Count:2 Data:[{Name:xiao hong1 ID:999999 Sex:1} {Name:xiao hong2 ID:999991 Sex:1}]}
	fmt.Printf("out detail: %+v\n", out)
	return mongoDB
}
