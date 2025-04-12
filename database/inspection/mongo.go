package inspection

import (
	"time"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
	libtime "tns-energo/lib/time"
	"tns-energo/service/inspection"

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

func (s *Mongo) AddOne(ctx libctx.Context, inspection inspection.Inspection) error {
	_, err := s.cli.
		Database(s.database).
		Collection(s.collection).
		InsertOne(ctx, MapToDb(inspection))

	return err
}

func (s *Mongo) GetByBrigadeId(ctx libctx.Context, log liblog.Logger, brigadeId string) ([]inspection.Inspection, error) {
	cursor, err := s.cli.
		Database(s.database).
		Collection(s.collection).
		Find(ctx, bson.M{"brigade.brigade_id": brigadeId})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			log.Errorf("failed to close cursor: %v", err)
		}
	}()

	var inspections []Inspection
	err = cursor.All(ctx, &inspections)
	if err != nil {
		return nil, err
	}

	return MapSliceToDomain(inspections), nil
}

func (s *Mongo) GetByInspectionDate(ctx libctx.Context, log liblog.Logger, inspectionDate time.Time) ([]inspection.Inspection, error) {
	gteDate := time.Date(inspectionDate.Year(), inspectionDate.Month(), inspectionDate.Day(), 0, 0, 0, 0, libtime.MoscowLocation())
	ltDate := gteDate.AddDate(0, 0, 1)
	cursor, err := s.cli.
		Database(s.database).
		Collection(s.collection).
		Find(ctx, bson.M{"inspection_date": bson.M{"$gte": gteDate, "$lt": ltDate}})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			log.Errorf("failed to close cursor: %v", err)
		}
	}()

	var inspections []Inspection
	err = cursor.All(ctx, &inspections)
	if err != nil {
		return nil, err
	}

	return MapSliceToDomain(inspections), nil
}
