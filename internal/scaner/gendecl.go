package scaner

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/g10z3r/archx/internal/scaner/buffer"
	"github.com/g10z3r/archx/internal/scaner/entity"
)

// func processGenDecl(collection *mongo.Collection, insertedID interface{}, buf *buffer.BufferEventBus, fs *token.FileSet, genDecl *ast.GenDecl) {
// 	if genDecl.Tok != token.TYPE {
// 		return
// 	}

// 	for _, spec := range genDecl.Specs {
// 		typeSpec, ok := spec.(*ast.TypeSpec)
// 		if !ok {
// 			continue
// 		}

// 		structType, ok := typeSpec.Type.(*ast.StructType)
// 		if !ok {
// 			continue
// 		}

// 		sType, err := processStructType(buf, fs, typeSpec, structType)
// 		if err != nil {
// 			errChan <- err
// 			continue
// 		}

// 		filter := bson.D{
// 			{Key: "_id", Value: insertedID.(primitive.ObjectID)},
// 			{Key: "packages.path", Value: "./example/cmd"}, // имя пакета, в который вы хотите добавить новую структуру
// 		}
// 		update := bson.D{
// 			{Key: "$push", Value: bson.D{
// 				{Key: "packages.$.structs", Value: sType},
// 			}},
// 		}

// 		updateResult, err := collection.UpdateOne(context.Background(), filter, update)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		log.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
// 	}
// }

func processStructType(buf *buffer.BufferEventBus, fs *token.FileSet, typeSpec *ast.TypeSpec, structType *ast.StructType) (*entity.Struct, error) {
	sType, usedPackages, err := entity.NewStructType(fs, structType, entity.NotEmbedded)
	if err != nil {
		return nil, fmt.Errorf("failed to create new struct type: %w", err)
	}

	for _, p := range usedPackages {
		if importIndex, exists := buf.ImportBuffer.GetIndexByAlias(p.Alias); exists {
			sType.AddDependency(importIndex, p.Element)
		}
	}

	return sType, nil
}

func notifyStructUpsert(buf *buffer.BufferEventBus, structName string, sType *entity.Struct) {
	buf.SendEvent(
		&buffer.UpsertStructEvent{
			StructInfo: sType,
			StructName: structName,
		},
	)
}
