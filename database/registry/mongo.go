package registry

import (
	"errors"
	"tns-energo/service/registry"

	"github.com/sunshineOfficial/golib/goctx"
	"github.com/sunshineOfficial/golib/golog"
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

func (s *Mongo) AddOne(ctx goctx.Context, item registry.Item) error {
	_, err := s.cli.
		Database(s.database).
		Collection(s.collection).
		InsertOne(ctx, MapToDb(item))

	return err
}

func (s *Mongo) AddMany(ctx goctx.Context, items []registry.Item) error {
	docs := make([]interface{}, 0, len(items))
	for _, item := range items {
		docs = append(docs, MapToDb(item))
	}

	_, err := s.cli.
		Database(s.database).
		Collection(s.collection).
		InsertMany(ctx, docs)

	return err
}

func (s *Mongo) GetByAccountNumber(ctx goctx.Context, accountNumber string) (registry.Item, error) {
	var item Item
	err := s.cli.
		Database(s.database).
		Collection(s.collection).
		FindOne(ctx, bson.M{"account_number": accountNumber}).
		Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return registry.Item{}, registry.ErrItemNotFound
		}

		return registry.Item{}, err
	}

	return MapToDomain(item), nil
}

func (s *Mongo) GetByAccountNumberRegular(ctx goctx.Context, log golog.Logger, accountNumber string) ([]registry.Item, error) {
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

	return MapSliceToDomain(items), nil
}

func (s *Mongo) UpdateOne(ctx goctx.Context, item registry.Item) error {
	_, err := s.cli.
		Database(s.database).
		Collection(s.collection).
		UpdateOne(ctx, bson.M{"account_number": item.AccountNumber}, bson.M{"$set": MapToDb(item)})

	return err
}
