package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type MongoDB struct {
	Client      *mongo.Client
	Ctx         context.Context
	Option      *option
	Error       error `json:"error:default nil"`
	ctxTime     uint
	serviceAddr string
}

type mongoDBer interface {
	NewDBClient(opts ...optFunc) *MongoDB
	PingTest() *MongoDB
	InseartOne(doc interface{}, opts ...optFunc) *MongoDB
	FindByID(id primitive.ObjectID, outDecode *interface{}, opts ...optFunc) *MongoDB
	Find(filter interface{}, outF interface{}, opts ...optFunc) *MongoDB
	QueryWithRelation(sCollname, rCollName, localF, foreignField, as string, opts ...optFunc) *MongoDB
}

type option struct {
	dbName     string
	collection string
}

type optFunc func(opt *option)

func SetDBName(dbName string) optFunc {
	return func(opt *option) {
		opt.dbName = dbName
	}
}

func SetColl(collection string) optFunc {
	return func(opt *option) {
		opt.collection = collection
	}
}

// 1. id is "", GenerateID will return NewObjectID
func GenerateID(id string) (primitive.ObjectID, error) {
	if id == "" {
		return primitive.NewObjectID(), nil
	}
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return oID, err
	}
	return oID, nil
}

func (mongoDB *MongoDB) prepareDB(opts ...optFunc) *MongoDB {
	opt := mongoDB.Option
	for _, f := range opts {
		f(opt)
	}
	return mongoDB
}

func Create(sAddr string, ctxTime uint, dbName, collection string) mongoDBer {
	return &MongoDB{
		serviceAddr: sAddr,
		ctxTime:     ctxTime,
		Option: &option{
			dbName:     dbName,
			collection: collection,
		},
	}
}
func (mongoDB *MongoDB) NewDBClient(opts ...optFunc) *MongoDB {
	mongoDB.prepareDB(opts...)
	mongoDB.Ctx, _ = context.WithTimeout(context.TODO(), time.Duration(mongoDB.ctxTime)*time.Second)
	mongoDB.Client, mongoDB.Error = mongo.Connect(mongoDB.Ctx, options.Client().ApplyURI(mongoDB.serviceAddr))
	return mongoDB
}

func (mongoDB *MongoDB) PingTest() *MongoDB {
	mongoDB.Error = mongoDB.Client.Ping(mongoDB.Ctx, readpref.Primary())
	return mongoDB
}

func (mongoDB *MongoDB) InseartOne(doc interface{}, opts ...optFunc) *MongoDB {
	mongoDB.prepareDB(opts...)
	result, err := mongoDB.Client.Database(mongoDB.Option.dbName).
		Collection(mongoDB.Option.collection).InsertOne(mongoDB.Ctx, doc)
	if err != nil {
		mongoDB.Error = err
		return mongoDB
	}
	log.Println("inseart succ:", result.InsertedID)
	return mongoDB
}

func (mongoDB *MongoDB) FindByID(id primitive.ObjectID, outDecode *interface{}, opts ...optFunc) *MongoDB {
	mongoDB.prepareDB(opts...)
	doc := bson.M{"_id": id}
	mongoDB.Error = mongoDB.Client.Database(mongoDB.Option.dbName).
		Collection(mongoDB.Option.collection).FindOne(mongoDB.Ctx, doc).Decode(&outDecode)
	return mongoDB
}

func (mongoDB *MongoDB) Find(filter interface{}, outF interface{}, opts ...optFunc) *MongoDB {
	mongoDB.prepareDB(opts...)
	cur, err := mongoDB.Client.Database(mongoDB.Option.dbName).
		Collection(mongoDB.Option.collection).Find(mongoDB.Ctx, filter)
	mongoDB.Error = err
	if err != nil {
		fmt.Println("error:", err)
		return mongoDB
	}
	defer cur.Close(mongoDB.Ctx)

	var out interface{}
	var outs []interface{}
	for cur.Next(mongoDB.Ctx) {
		cur.Decode(&out)
		fmt.Printf("detail : %+v\n", outs)
		outs = append(outs, out)
	}
	outF = &outs
	fmt.Printf("detail : %+v\n", outF)
	return mongoDB
}

// TODO update in next time
func (mongoDB *MongoDB) QueryWithRelation(sCollname, rCollName, localF, foreignField, as string, opts ...optFunc) *MongoDB {
	mongoDB.prepareDB(opts...)
	cur, err := mongoDB.Client.Database(mongoDB.Option.dbName).
		Collection(mongoDB.Option.collection).Aggregate(mongoDB.Ctx, mongo.Pipeline{
		{
			{
				"$match",
				bson.M{
					"_id": "",
				},
			},
		},
		{
			{
				"$lookup",
				bson.D{
					{"from", rCollName},
					{"localField", localF},
					{"foreignFiled", foreignField},
					{"as", as},
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

	var out interface{}
	var outs []interface{}
	for cur.Next(mongoDB.Ctx) {
		cur.Decode(&out)
		outs = append(outs, out)
	}
	fmt.Printf("detail : %+v\n", outs)
	return mongoDB
}
