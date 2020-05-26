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

func (r *MongoDBRepository) Update(ctx context.Context, id string, req *types.UpdateBike) error {
	s := r.session.Clone()
	defer s.Clone()

	return r.collection(s).Update(bson.M{"bike_id": id}, bson.M{
		"$set": bson.M{
			"updated_at": time.Now(),
			"number":     req.Number,
			"color":      req.Color,
			"cost":       req.Cost,
			"status":     req.Status,
			"name":       req.Name,
		},
	},
	)

}

func (r *MongoDBRepository) Delete(ctx context.Context, id string) error {
	s := r.session.Clone()
	defer s.Close()
	if err := r.collection(s).Remove(bson.M{"bike_id": id}); err != nil {
		return err
	}
	return nil
}

//====================================================
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

func (r *MongoDBRepository) FindByBikeID(ctx context.Context, bikeID string) (*types.Bike, error) {
	selector := bson.M{"bike_id": bikeID}
	s := r.session.Clone()
	defer s.Close()
	var bike *types.Bike
	if err := r.collection(s).Find(selector).One(&bike); err != nil {
		return nil, err
	}
	return bike, nil
}

func (r *MongoDBRepository) FindAll(ctx context.Context) ([]*types.Bike, error) {
	s := r.session.Clone()
	defer s.Close()
	var bikes []*types.Bike
	if err := r.collection(s).Find(nil).All(&bikes); err != nil {
		return nil, err
	}
	return bikes, nil
}

func (r *MongoDBRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("").C("bike")
}
