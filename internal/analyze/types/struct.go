package types

import (
	"go/ast"
	"go/token"
)

type StructType struct {
	_      [0]int
	pos    token.Pos
	end    token.Pos
	Field  map[string]*FieldType
	Method map[string]map[string]struct{}
}

func MakeStructType(res *ast.StructType) *StructType {
	return &StructType{
		pos:    res.Pos(),
		end:    res.End(),
		Field:  extractFieldMap(res.Fields.List),
		Method: make(map[string]map[string]struct{}),
	}
}

func extractFieldMap(fieldList []*ast.Field) map[string]*FieldType {
	fieldMap := make(map[string]*FieldType)

	for _, field := range fieldList {
		for _, name := range field.Names {
			fieldMap[name.Name] = &FieldType{
				pos:    name.Pos(),
				end:    name.End(),
				Public: name.IsExported(),
			}
		}
	}

	return fieldMap
}

type FieldType struct {
	_      [0]int
	pos    token.Pos
	end    token.Pos
	Public bool
}

func MakeFieldType(res *ast.Ident) *FieldType {
	return &FieldType{
		pos:    res.Pos(),
		end:    res.End(),
		Public: res.IsExported(),
	}
}
