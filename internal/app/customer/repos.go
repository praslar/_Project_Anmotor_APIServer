package customer

import "github.com/globalsign/mgo"

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
	return s.DB("").C("customers")
}