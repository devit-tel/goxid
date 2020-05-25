package goxid

import (
	"github.com/rs/xid"
	mongo "go.mongodb.org/mongo-driver/bson/primitive"
)

type ID struct {
	freezeID string
}

var instance *ID

func New() *ID {
	if instance == nil {
		instance = &ID{}
	}

	return instance
}

func (i *ID) Gen() string {
	if i.freezeID != "" {
		return i.freezeID
	}

	return xid.New().String()
}

func (i *ID) Freeze(id string) {
	i.freezeID = id
}

func (i *ID) GenObjectID() string {
	if i.freezeID != "" {
		return i.freezeID
	}

	return mongo.NewObjectID().Hex()
}
