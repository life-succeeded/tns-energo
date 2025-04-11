package task

import (
	"errors"
	"fmt"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
	"tns-energo/service/task"

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

func (s *Mongo) AddOne(ctx libctx.Context, t task.Task) (string, error) {
	result, err := s.cli.
		Database(s.database).
		Collection(s.collection).
		InsertOne(ctx, MapToDb(t))
	if err != nil {
		return "", err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("failed to convert inserted ID")
	}

	return id.Hex(), nil
}

func (s *Mongo) GetByBrigadeId(ctx libctx.Context, log liblog.Logger, brigadeId string) ([]task.Task, error) {
	cursor, err := s.cli.
		Database(s.database).
		Collection(s.collection).
		Find(ctx, bson.M{"brigade_id": brigadeId})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			log.Errorf("failed to close cursor: %v", err)
		}
	}()

	var tasks []Task
	err = cursor.All(ctx, &tasks)
	if err != nil {
		return nil, err
	}

	return MapSliceToDomain(tasks), nil
}

func (s *Mongo) UpdateStatus(ctx libctx.Context, id string, status task.Status) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert id to ObjectID: %v", err)
	}

	_, err = s.cli.
		Database(s.database).
		Collection(s.collection).
		UpdateByID(ctx, objectID, bson.M{"$set": bson.M{"status": int(status)}})

	return err
}
