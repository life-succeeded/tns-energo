package report

import (
	libctx "tns-energo/lib/ctx"
	domain "tns-energo/service/analytics"

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

func (s *Mongo) AddOne(ctx libctx.Context, r domain.Report) error {
	_, err := s.cli.
		Database(s.database).
		Collection(s.collection).
		InsertOne(ctx, mapToDb(r))

	return err
}
