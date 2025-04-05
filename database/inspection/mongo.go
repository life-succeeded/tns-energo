package inspection

import (
	libctx "tns-energo/lib/ctx"

	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	cli                  *mongo.Client
	database, collection string
}

func NewRepository(cli *mongo.Client, database, collection string) *Mongo {
	return &Mongo{
		cli:        cli,
		database:   database,
		collection: collection,
	}
}

func (r *Mongo) CreateOne(ctx libctx.Context, inspection Inspection) error {
	_, err := r.cli.
		Database(r.database).
		Collection(r.collection).
		InsertOne(ctx, inspection)

	return err
}

func (r *Mongo) CreateMany(ctx libctx.Context, inspections []Inspection) error {
	docs := make([]interface{}, 0, len(inspections))
	for _, inspection := range inspections {
		docs = append(docs, inspection)
	}

	_, err := r.cli.
		Database(r.database).
		Collection(r.collection).
		InsertMany(ctx, docs)

	return err
}
