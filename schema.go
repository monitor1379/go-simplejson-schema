package simplejsonschema

import "reflect"

type FieldSchema struct {
	Key        string
	TypeOfVal  reflect.Kind
	Required   bool
	DefaultVal interface{}
}

type JsonSchema []FieldSchema
