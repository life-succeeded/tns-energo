package report

import (
	"errors"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
	domain "tns-energo/service/analytics"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (s *Mongo) AddOne(ctx libctx.Context, r domain.Report) (string, error) {
	result, err := s.cli.
		Database(s.database).
		Collection(s.collection).
		InsertOne(ctx, MapToDb(r))
	if err != nil {
		return "", err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("failed to convert inserted ID")
	}

	return id.Hex(), nil
}

func (s *Mongo) GetAll(ctx libctx.Context, log liblog.Logger) ([]domain.Report, error) {
	cursor, err := s.cli.
		Database(s.database).
		Collection(s.collection).
		Find(ctx, bson.M{}, options.Find().SetSort(bson.D{{"created_at", -1}}))
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			log.Errorf("failed to close cursor: %v", err)
		}
	}()

	var reports []Report
	err = cursor.All(ctx, &reports)
	if err != nil {
		return nil, err
	}

	return MapSliceToDomain(reports), nil
}
