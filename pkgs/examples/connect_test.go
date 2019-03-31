package examples

import (
	"testing"
)

func TestConnect(t *testing.T) {
	mgo := Connect()
	if mgo.Error != nil {
		t.Error(mgo.Error)
	}
}
