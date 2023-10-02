package scanner

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/g10z3r/archx/internal/domain/entity"
	"github.com/g10z3r/archx/internal/infrastructure/db/mongodb/scanner/model"
)

type structRepository struct {
	documentID primitive.ObjectID
	collection *mongo.Collection
}

func newStructRepository(docID primitive.ObjectID, col *mongo.Collection) *structRepository {
	return &structRepository{
		documentID: docID,
		collection: col,
	}
}

func (r *structRepository) Append(ctx context.Context, structEntity *entity.StructEntity, structIndex int, pkgPath string) error {
	filter := bson.D{
		{Key: "_id", Value: r.documentID},
		{Key: "packages.path", Value: pkgPath},
	}

	update := bson.D{
		{
			Key: "$push", Value: bson.D{
				{Key: "packages.$.structs", Value: model.MapStructEntity(structEntity)},
			},
		},
		{
			Key: "$set", Value: bson.D{
				{Key: fmt.Sprintf("packages.$.structsIndex.%s", *structEntity.Name), Value: structIndex},
			},
		},
	}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}