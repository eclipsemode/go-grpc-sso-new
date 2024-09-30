package mongo

import (
	"gopkg.in/mgo.v2"
	"log"
)

type Storage struct {
	session *mgo.Session
	db      *mgo.Database
}

func NewStorage(storagePath string, dbName string) (*Storage, error) {
	s, err := mgo.Dial(storagePath)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %s", err.Error())
	}

	s.SetMode(mgo.Primary, true)

	return &Storage{
		session: s,
		db:      s.DB(dbName),
	}, nil
}

func (s *Storage) Close() {
	s.session.Close()
}
