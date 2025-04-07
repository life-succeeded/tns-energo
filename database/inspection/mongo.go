package inspection

import (
	libctx "tns-energo/lib/ctx"
	"tns-energo/service/inspection"

	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	cli                  *mongo.Client
	database, collection string
}

func NewStorage(cli *mongo.Client, database, collection string) *Mongo {
	return &Mongo{
		cli:        cli,
		database:   database,
		collection: collection,
	}
}

func (s *Mongo) AddOne(ctx libctx.Context, inspection inspection.Inspection) error {
	_, err := s.cli.
		Database(s.database).
		Collection(s.collection).
		InsertOne(ctx, mapToDb(inspection))

	return err
}
