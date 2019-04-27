package examples

import (
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestEgFindByID(t *testing.T) {
	ret, err := EgFindByID()
	if err != nil {
		t.Error(err)
	}
	t.Logf("info:%+v", ret)
	u := user{}
	bData, err := bson.Marshal(ret)
	if err != nil {
		t.Error(err)
	}
	err = bson.Unmarshal(bData, &u)
	if err != nil {
		t.Error(err)
	}

	t.Logf("info:%+v", u)
}

func TestEgFind(t *testing.T) {
	ret, err := EgFind()
	if err != nil {
		t.Error(err)
	}
	t.Logf("info:%+v", ret)
	u := user{}
	bData, err := bson.Marshal(ret)
	if err != nil {
		t.Error(err)
	}
	err = bson.Unmarshal(bData, &u)
	if err != nil {
		t.Error(err)
	}

	t.Logf("info:%+v", u)
}

func TestEgFindWithCount(t *testing.T) {
	ret, err := EgFindWithCount()
	if err != nil {
		t.Error(err)
	}
	t.Logf("info:%+v", ret)
}
