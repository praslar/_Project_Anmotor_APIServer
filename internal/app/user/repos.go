package user

import (
	"context"

	"github.com/anmotor/internal/app/types"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type (
	MongoDBRepository struct {
		session *mgo.Session
	}
)

func NewMongoDBRepo(session *mgo.Session) *MongoDBRepository {
	return &MongoDBRepository{
		session: session,
	}
}

func (r *MongoDBRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("").C("users")
}

func (r *MongoDBRepository) FindUser(ctx context.Context, username string) (*types.User, error) {
	s := r.session.Clone()
	defer s.Close()
	var usr *types.User
	if err := r.collection(s).Find(bson.M{
		"username": username,
	}).One(&usr); err != nil {
		return nil, err
	}
	return usr, nil
}
