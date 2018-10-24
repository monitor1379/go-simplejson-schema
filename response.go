package simplejsonschema

import (
	"fmt"
	"reflect"
)

type Response map[string]interface{}

func NewBaseResponse(code int, message string) Response {
	return Response{
		"code":    code,
		"message": message,
	}
}

var (
	SuccessResponseCode           = 0
	MissFieldInSchemaResponseCode = -1
	ParseErrorResponseCode        = -2
	NeedArgErrorResponseCode      = -3
	ArgTypeErrorResponseCode      = -4
)

func NewSuccessResponse() Response {
	return NewBaseResponse(SuccessResponseCode, "Success")
}

func NewMissFieldInSchemaResponse(key string) Response {
	return NewBaseResponse(MissFieldInSchemaResponseCode, fmt.Sprintf("Can not find key=%s in JsonSchema!", key))
}

func NewParseErrorResponse() Response {
	return NewBaseResponse(ParseErrorResponseCode, "Can not parse request json!")
}

func NewNeedArgErrorResponse(key string, typeOfVal reflect.Kind) Response {
	var msg = fmt.Sprintf("Need field: '%s' (type: '%s')!", key, typeOfVal)
	return NewBaseResponse(NeedArgErrorResponseCode, msg)
}

func NewArgTypeErrorResponse(key string, typeOfField string) Response {
	var msg = fmt.Sprintf("Field type of field '%s' must be '%s'", key, typeOfField)
	return NewBaseResponse(ArgTypeErrorResponseCode, msg)
}
