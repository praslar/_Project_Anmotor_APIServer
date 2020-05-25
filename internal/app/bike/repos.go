package bike

import (
	"context"
	"time"

	"github.com/anmotor/internal/app/types"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type (
	MongoDBRepository struct {
		session *mgo.Session
	}
)

func NewMongoDBRespository(session *mgo.Session) *MongoDBRepository {
	return &MongoDBRepository{
		session: session,
	}
}

func (r *MongoDBRepository) Create(ctx context.Context, bike *types.Bike) error {
	s := r.session.Clone()
	defer s.Close()
	bike.CreatedAt = time.Now()
	bike.UpdateAt = bike.CreatedAt

	if err := r.collection(s).Insert(bike); err != nil {
		return err
	}
	return nil
}

func (r *MongoDBRepository) FindByNumber(ctx context.Context, number string) (*types.Bike, error) {
	selector := bson.M{"number": number}
	s := r.session.Clone()
	defer s.Close()
	var user *types.Bike
	if err := r.collection(s).Find(selector).One(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *MongoDBRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("").C("bike")
}
