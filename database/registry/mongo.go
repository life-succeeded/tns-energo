package registry

import (
	libctx "tns-energo/lib/ctx"

	"go.mongodb.org/mongo-driver/bson"
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

func (r *Mongo) AddOne(ctx libctx.Context, item Item) error {
	_, err := r.cli.
		Database(r.database).
		Collection(r.collection).
		InsertOne(ctx, item)

	return err
}

func (r *Mongo) AddMany(ctx libctx.Context, items []Item) error {
	docs := make([]interface{}, 0, len(items))
	for _, inspection := range items {
		docs = append(docs, inspection)
	}

	_, err := r.cli.
		Database(r.database).
		Collection(r.collection).
		InsertMany(ctx, docs)

	return err
}

func (r *Mongo) GetByAccountNumber(ctx libctx.Context, accountNumber string) (Item, error) {
	var item Item
	err := r.cli.
		Database(r.database).
		Collection(r.collection).
		FindOne(ctx, bson.M{"account_number": accountNumber}).
		Decode(&item)

	return item, err
}
