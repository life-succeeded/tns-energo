package brigade

import (
	"errors"
	"fmt"
	libctx "tns-energo/lib/ctx"
	domain "tns-energo/service/brigade"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (s *Mongo) AddOne(ctx libctx.Context, b domain.Brigade) (string, error) {
	result, err := s.cli.
		Database(s.database).
		Collection(s.collection).
		InsertOne(ctx, MapToDb(b))
	if err != nil {
		return "", err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("failed to convert inserted ID")
	}

	return id.Hex(), nil
}

func (s *Mongo) GetById(ctx libctx.Context, id string) (domain.Brigade, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Brigade{}, fmt.Errorf("failed to convert id to ObjectID: %v", err)
	}

	var brigade Brigade
	err = s.cli.
		Database(s.database).
		Collection(s.collection).
		FindOne(ctx, bson.M{"_id": objectID}).
		Decode(&brigade)
	if err != nil {
		return domain.Brigade{}, err
	}

	return MapToDomain(brigade), nil
}
