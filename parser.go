package simplejsonschema

import (
	"github.com/bitly/go-simplejson"
	"fmt"
	"errors"
	"reflect"
)

type JsonFieldGetter struct {
	j     *simplejson.Json
	fsmap map[string]FieldSchema
}

func NewJsonFieldGetter(j *simplejson.Json, js *JsonSchema) *JsonFieldGetter {
	fsmap := make(map[string]FieldSchema)
	for _, fieldSchema := range *js {
		fsmap[fieldSchema.Key] = fieldSchema
	}
	return &JsonFieldGetter{j, fsmap}
	return nil
}

func (this *JsonFieldGetter) Get(key string) (interface{}, error, Response) {
	fieldSchema, ok := this.fsmap[key]
	if !ok {
		panic(fmt.Sprintf("Can not find key=%s in JsonSchema", key))
		//return nil, errors.New(fmt.Sprintf("Can not find key=%s in JsonSchema", key)), NewMissFieldInSchemaResponse(key)
	}
	valJson, ok := this.j.CheckGet(key)
	if !ok {
		if fieldSchema.Required {
			return nil,
				errors.New(fmt.Sprintf("Need field: '%s' (type: '%s')!", key, fieldSchema.TypeOfVal)),
				NewNeedArgErrorResponse(key, fieldSchema.TypeOfVal)
		} else {
			if reflect.TypeOf(fieldSchema.DefaultVal).Kind() != fieldSchema.TypeOfVal {
				panic(fmt.Sprintf("Mismatch type of DefaultVal and TypeOfVal!"))
			}
			return fieldSchema.DefaultVal,
				nil,
				NewSuccessResponse()
		}
	} else {
		var val interface{} = nil
		var err error = nil
		switch fieldSchema.TypeOfVal {
		case reflect.String:
			val, err = valJson.String()
			break
		case reflect.Int:
			val, err = valJson.Int()
			break
		case reflect.Float32:
			val, err = valJson.Float64()
			break
		case reflect.Float64:
			val, err = valJson.Float64()
			break
		default:
			err = errors.New(fmt.Sprintf("Type of value of key=%s assertion to %s failed", key, fieldSchema.TypeOfVal.String()))
		}

		if err != nil {
			return nil, err, NewArgTypeErrorResponse(key, fieldSchema.TypeOfVal.String())
		} else {
			return val, nil, NewSuccessResponse()
		}
	}

}

func (this *JsonFieldGetter) GetString(key string) (*string, error, Response) {
	fieldSchema, ok := this.fsmap[key]
	if !ok {
		panic(fmt.Sprintf("Can not find key=%s in JsonSchema", key))
		//return nil, errors.New(fmt.Sprintf("Can not find key=%s in JsonSchema", key)), NewMissFieldInSchemaResponse(key)
	}
	valJson, ok := this.j.CheckGet(key)
	if !ok {
		if fieldSchema.Required {
			return nil,
				errors.New(fmt.Sprintf("Need field: '%s' (type: '%s')!", key, fieldSchema.TypeOfVal)),
				NewNeedArgErrorResponse(key, fieldSchema.TypeOfVal)
		} else {
			if reflect.TypeOf(fieldSchema.DefaultVal).Kind() != fieldSchema.TypeOfVal {
				panic(fmt.Sprintf("Mismatch type of DefaultVal and TypeOfVal!"))
			}
			val := fieldSchema.DefaultVal.(string)
			return &val,
				nil,
				NewSuccessResponse()
		}
	} else {
		val, err := valJson.String()
		if err != nil {
			return nil, err, NewArgTypeErrorResponse(key, fieldSchema.TypeOfVal.String())
		} else {
			return &val, nil, NewSuccessResponse()
		}
	}
}

func (this *JsonFieldGetter) GetInt(key string) (*int, error, Response) {
	fieldSchema, ok := this.fsmap[key]
	if !ok {
		panic(fmt.Sprintf("Can not find key=%s in JsonSchema", key))
		//return nil, errors.New(fmt.Sprintf("Can not find key=%s in JsonSchema", key)), NewMissFieldInSchemaResponse(key)
	}
	valJson, ok := this.j.CheckGet(key)
	if !ok {
		if fieldSchema.Required {
			return nil,
				errors.New(fmt.Sprintf("Need field: '%s' (type: '%s')!", key, fieldSchema.TypeOfVal)),
				NewNeedArgErrorResponse(key, fieldSchema.TypeOfVal)
		} else {
			if reflect.TypeOf(fieldSchema.DefaultVal).Kind() != fieldSchema.TypeOfVal {
				panic(fmt.Sprintf("Mismatch type of DefaultVal and TypeOfVal!"))
			}
			val := fieldSchema.DefaultVal.(int)
			return &val,
				nil,
				NewSuccessResponse()
		}
	} else {
		val, err := valJson.Int()
		if err != nil {
			return nil, err, NewArgTypeErrorResponse(key, fieldSchema.TypeOfVal.String())
		} else {
			return &val, nil, NewSuccessResponse()
		}
	}
}

func (this *JsonFieldGetter) GetFloat32(key string) (*float32, error, Response) {
	fieldSchema, ok := this.fsmap[key]
	if !ok {
		panic(fmt.Sprintf("Can not find key=%s in JsonSchema", key))
		//return nil, errors.New(fmt.Sprintf("Can not find key=%s in JsonSchema", key)), NewMissFieldInSchemaResponse(key)
	}
	valJson, ok := this.j.CheckGet(key)
	if !ok {
		if fieldSchema.Required {
			return nil,
				errors.New(fmt.Sprintf("Need field: '%s' (type: '%s')!", key, fieldSchema.TypeOfVal)),
				NewNeedArgErrorResponse(key, fieldSchema.TypeOfVal)
		} else {
			if reflect.TypeOf(fieldSchema.DefaultVal).Kind() != fieldSchema.TypeOfVal {
				panic(fmt.Sprintf("Mismatch type of DefaultVal and TypeOfVal!"))
			}
			val := fieldSchema.DefaultVal.(float32)
			return &val,
				nil,
				NewSuccessResponse()
		}
	} else {
		val, err := valJson.Float64()
		if err != nil {
			return nil, err, NewArgTypeErrorResponse(key, fieldSchema.TypeOfVal.String())
		} else {
			_val := float32(val)
			return &_val, nil, NewSuccessResponse()
		}
	}
}

func (this *JsonFieldGetter) GetFloat64(key string) (*float64, error, Response) {
	fieldSchema, ok := this.fsmap[key]
	if !ok {
		panic(fmt.Sprintf("Can not find key=%s in JsonSchema", key))
		//return nil, errors.New(fmt.Sprintf("Can not find key=%s in JsonSchema", key)), NewMissFieldInSchemaResponse(key)
	}
	valJson, ok := this.j.CheckGet(key)
	if !ok {
		if fieldSchema.Required {
			return nil,
				errors.New(fmt.Sprintf("Need field: '%s' (type: '%s')!", key, fieldSchema.TypeOfVal)),
				NewNeedArgErrorResponse(key, fieldSchema.TypeOfVal)
		} else {
			if reflect.TypeOf(fieldSchema.DefaultVal).Kind() != fieldSchema.TypeOfVal {
				panic(fmt.Sprintf("Mismatch type of DefaultVal and TypeOfVal!"))
			}
			val := fieldSchema.DefaultVal.(float64)
			return &val,
				nil,
				NewSuccessResponse()
		}
	} else {
		val, err := valJson.Float64()
		if err != nil {
			return nil, err, NewArgTypeErrorResponse(key, fieldSchema.TypeOfVal.String())
		} else {
			return &val, nil, NewSuccessResponse()
		}
	}
}
