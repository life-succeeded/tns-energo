package task

import (
	"errors"
	"fmt"
	"tns-energo/service/task"

	"github.com/sunshineOfficial/golib/goctx"
	"github.com/sunshineOfficial/golib/golog"
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

func (s *Mongo) AddOne(ctx goctx.Context, t task.Task) (string, error) {
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

func (s *Mongo) GetByBrigadeId(ctx goctx.Context, log golog.Logger, brigadeId string) ([]task.Task, error) {
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

func (s *Mongo) UpdateStatus(ctx goctx.Context, id string, status task.Status) error {
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

func (s *Mongo) GetById(ctx goctx.Context, id string) (task.Task, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return task.Task{}, fmt.Errorf("failed to convert id to ObjectID: %v", err)
	}

	var t Task
	err = s.cli.
		Database(s.database).
		Collection(s.collection).
		FindOne(ctx, bson.M{"_id": objectID}).
		Decode(&t)
	if err != nil {
		return task.Task{}, err
	}

	return MapToDomain(t), nil
}
