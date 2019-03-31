package main

import "github.com/sayluck/mongo-example/pkgs/examples"

type user struct {
	Name string
	ID   string
	Sex  int
}

func main() {
	// eg 1
	mgo := examples.Connect()

	// eg 2 inseart one: use struct
	doc := &user{
		Name: "xiao min",
		ID:   "0123456789",
		Sex:  1,
	}
	ret := mgo.InseartOne(doc)
	if ret.Error != nil {
		panic(ret.Error)
	}
}
