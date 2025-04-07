package registry

import (
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
	"tns-energo/service/registry"

	"go.mongodb.org/mongo-driver/bson"
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

func (s *Mongo) AddOne(ctx libctx.Context, item registry.Item) error {
	_, err := s.cli.
		Database(s.database).
		Collection(s.collection).
		InsertOne(ctx, mapToDb(item))

	return err
}

func (s *Mongo) AddMany(ctx libctx.Context, items []registry.Item) error {
	docs := make([]interface{}, 0, len(items))
	for _, item := range items {
		docs = append(docs, mapToDb(item))
	}

	_, err := s.cli.
		Database(s.database).
		Collection(s.collection).
		InsertMany(ctx, docs)

	return err
}

func (s *Mongo) GetByAccountNumber(ctx libctx.Context, accountNumber string) (registry.Item, error) {
	var item Item
	err := s.cli.
		Database(s.database).
		Collection(s.collection).
		FindOne(ctx, bson.M{"account_number": accountNumber}).
		Decode(&item)

	return mapToDomain(item), err
}

func (s *Mongo) GetByAccountNumberRegular(ctx libctx.Context, log liblog.Logger, accountNumber string) ([]registry.Item, error) {
	cursor, err := s.cli.
		Database(s.database).
		Collection(s.collection).
		Find(ctx, bson.M{"account_number": bson.M{"$regex": accountNumber, "$options": "i"}})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			log.Errorf("failed to close cursor: %v", err)
		}
	}()

	var items []Item
	err = cursor.All(ctx, &items)
	if err != nil {
		return nil, err
	}

	return mapSliceToDomain(items), nil
}

func (s *Mongo) UpdateOne(ctx libctx.Context, item registry.Item) error {
	_, err := s.cli.
		Database(s.database).
		Collection(s.collection).
		UpdateOne(ctx, bson.M{"account_number": item.AccountNumber}, bson.M{"$set": mapToDb(item)})

	return err
}
