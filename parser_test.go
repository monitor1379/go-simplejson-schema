package simplejsonschema

import (
	"testing"
	"github.com/bitly/go-simplejson"
	"reflect"
	"errors"
)

func TestJsonFieldGetter_Get1(t *testing.T) {
	jsonStr := "{\"name\":\"dengzhenpeng\"}"
	sjson, _ := simplejson.NewJson([]byte(jsonStr))
	var personJsonSchema = JsonSchema{{Key: "name", TypeOfVal: reflect.String, Required: true}}
	jsonGetter := NewJsonFieldGetter(sjson, &personJsonSchema)
	val, err, resp := jsonGetter.Get("name")
	if val != "dengzhenpeng" {
		t.Errorf("Actually: [%v],  Expected: [%v]", val, "dengzhenpeng")
	}
	if err != nil {
		t.Errorf("Actually: [%v],  Expected: [%v]", err, nil)
	}
	if resp == nil || resp["code"] != SuccessResponseCode {
		t.Errorf("Actually: [%v],  Expected: [%v]", resp, NewSuccessResponse())
	}

}

func TestJsonFieldGetter_Get2(t *testing.T) {
	jsonStr := "{\"age\": 23}"
	sjson, _ := simplejson.NewJson([]byte(jsonStr))
	var personJsonSchema = JsonSchema{{Key: "age", TypeOfVal: reflect.Int, Required: true}}
	jsonGetter := NewJsonFieldGetter(sjson, &personJsonSchema)
	val, err, resp := jsonGetter.Get("age")
	if val != 23 {
		t.Errorf("Actually: [%v],  Expected: [%v]", val, 23)
	}
	if err != nil {
		t.Errorf("Actually: [%v],  Expected: [%v]", err, nil)
	}
	if resp == nil || resp["code"] != SuccessResponseCode {
		t.Errorf("Actually: [%v],  Expected: [%v]", resp, NewSuccessResponse())
	}

}

func TestJsonFieldGetter_Get3(t *testing.T) {
	jsonStr := "{\"name\":\"dengzhenpeng\"}"
	sjson, _ := simplejson.NewJson([]byte(jsonStr))
	var personJsonSchema = JsonSchema{{Key: "age", TypeOfVal: reflect.Int, Required: false, DefaultVal: 23}}
	jsonGetter := NewJsonFieldGetter(sjson, &personJsonSchema)
	val, err, resp := jsonGetter.Get("age")
	if val != 23 {
		t.Errorf("Actually: [%v],  Expected: [%v]", val, "dengzhenpeng")
	}
	if err != nil {
		t.Errorf("Actually: [%v],  Expected: [%v]", err, nil)
	}
	if resp == nil || resp["code"] != SuccessResponseCode {
		t.Errorf("Actually: [%v],  Expected: [%v]", resp, NewSuccessResponse())
	}

}

func TestJsonFieldGetter_Get4(t *testing.T) {
	jsonStr := "{}"
	sjson, _ := simplejson.NewJson([]byte(jsonStr))
	var personJsonSchema = JsonSchema{{Key: "name", TypeOfVal: reflect.String, Required: false, DefaultVal: "dengzhenpeng"}}
	jsonGetter := NewJsonFieldGetter(sjson, &personJsonSchema)
	val, err, resp := jsonGetter.Get("name")
	if val != "dengzhenpeng" {
		t.Errorf("Actually: [%v],  Expected: [%v]", val, 23)
	}
	if err != nil {
		t.Errorf("Actually: [%v],  Expected: [%v]", err, nil)
	}
	if resp == nil || resp["code"] != SuccessResponseCode {
		t.Errorf("Actually: [%v],  Expected: [%v]", resp, NewSuccessResponse())
	}

}

func TestJsonFieldGetter_Get5(t *testing.T) {
	jsonStr := "{\"age\":\"23\"}"
	sjson, _ := simplejson.NewJson([]byte(jsonStr))
	var personJsonSchema = JsonSchema{{Key: "age", TypeOfVal: reflect.Int, Required: true, DefaultVal: 23}}
	jsonGetter := NewJsonFieldGetter(sjson, &personJsonSchema)
	val, err, resp := jsonGetter.Get("age")
	if val != nil {
		t.Errorf("Actually: [%v],  Expected: [%v]", val, nil)
	}
	if err.Error() != errors.New("invalid value type").Error() {
		t.Errorf("Actually: [%s],  Expected: [%s]", err, errors.New("invalid value type"))
	}
	if resp == nil || resp["code"] != ArgTypeErrorResponseCode {
		t.Errorf("Actually: [%v],  Expected: [%v]", resp, NewArgTypeErrorResponse("age", reflect.Int.String()))
	}

}

func TestJsonFieldGetter_GetString1(t *testing.T) {
	jsonStr := "{\"name\":\"dengzhenpeng\"}"
	sjson, _ := simplejson.NewJson([]byte(jsonStr))
	var personJsonSchema = JsonSchema{{Key: "name", TypeOfVal: reflect.String, Required: true}}
	jsonGetter := NewJsonFieldGetter(sjson, &personJsonSchema)
	val, err, resp := jsonGetter.GetString("name")
	if *val != "dengzhenpeng" {
		t.Errorf("Actually: [%v],  Expected: [%v]", val, nil)
	}
	if err != nil {
		t.Errorf("Actually: [%s],  Expected: [%v]", err, nil)
	}
	if resp == nil || resp["code"] != SuccessResponseCode {
		t.Errorf("Actually: [%v],  Expected: [%v]", resp, NewSuccessResponse())
	}

}

func TestJsonFieldGetter_GetString2(t *testing.T) {
	jsonStr := "{}"
	sjson, _ := simplejson.NewJson([]byte(jsonStr))
	var personJsonSchema = JsonSchema{{Key: "name", TypeOfVal: reflect.String, Required: false, DefaultVal: "dengzhenpeng"}}
	jsonGetter := NewJsonFieldGetter(sjson, &personJsonSchema)
	val, err, resp := jsonGetter.GetString("name")
	if *val != "dengzhenpeng" {
		t.Errorf("Actually: [%v],  Expected: [%v]", val, nil)
	}
	if err != nil {
		t.Errorf("Actually: [%s],  Expected: [%v]", err, nil)
	}
	if resp == nil || resp["code"] != SuccessResponseCode {
		t.Errorf("Actually: [%v],  Expected: [%v]", resp, NewSuccessResponse())
	}

}

func TestJsonFieldGetter_GetInt1(t *testing.T) {
	jsonStr := "{\"age\":23}"
	sjson, _ := simplejson.NewJson([]byte(jsonStr))
	var personJsonSchema = JsonSchema{{Key: "age", TypeOfVal: reflect.Int, Required: true}}
	jsonGetter := NewJsonFieldGetter(sjson, &personJsonSchema)
	val, err, resp := jsonGetter.GetInt("age")
	if *val != 23 {
		t.Errorf("Actually: [%v],  Expected: [%v]", val, nil)
	}
	if err != nil {
		t.Errorf("Actually: [%s],  Expected: [%v]", err, nil)
	}
	if resp == nil || resp["code"] != SuccessResponseCode {
		t.Errorf("Actually: [%v],  Expected: [%v]", resp, NewSuccessResponse())
	}

}

func TestJsonFieldGetter_GetInt2(t *testing.T) {
	jsonStr := "{}"
	sjson, _ := simplejson.NewJson([]byte(jsonStr))
	var personJsonSchema = JsonSchema{{Key: "age", TypeOfVal: reflect.Int, Required: false, DefaultVal: 23}}
	jsonGetter := NewJsonFieldGetter(sjson, &personJsonSchema)
	val, err, resp := jsonGetter.GetInt("age")
	if *val != 23 {
		t.Errorf("Actually: [%v],  Expected: [%v]", val, nil)
	}
	if err != nil {
		t.Errorf("Actually: [%s],  Expected: [%v]", err, nil)
	}
	if resp == nil || resp["code"] != SuccessResponseCode {
		t.Errorf("Actually: [%v],  Expected: [%v]", resp, NewSuccessResponse())
	}

}

func TestJsonFieldGetter_GetBool1(t *testing.T) {
	jsonStr := "{\"isMale\":true}"
	sjson, _ := simplejson.NewJson([]byte(jsonStr))
	var personJsonSchema = JsonSchema{{Key: "isMale", TypeOfVal: reflect.Bool, Required: true}}
	jsonGetter := NewJsonFieldGetter(sjson, &personJsonSchema)
	val, err, resp := jsonGetter.GetBool("isMale")
	if *val != true {
		t.Errorf("Actually: [%v],  Expected: [%v]", val, nil)
	}
	if err != nil {
		t.Errorf("Actually: [%s],  Expected: [%v]", err, nil)
	}
	if resp == nil || resp["code"] != SuccessResponseCode {
		t.Errorf("Actually: [%v],  Expected: [%v]", resp, NewSuccessResponse())
	}

}

func TestJsonFieldGetter_GetBool2(t *testing.T) {
	jsonStr := "{}"
	sjson, _ := simplejson.NewJson([]byte(jsonStr))
	var personJsonSchema = JsonSchema{{Key: "isMale", TypeOfVal: reflect.Bool, Required: true}}
	jsonGetter := NewJsonFieldGetter(sjson, &personJsonSchema)
	val, err, resp := jsonGetter.GetBool("isMale")
	if err == nil {
		t.Errorf("Actually: [%s],  Expected: [%v]", err, nil)
	}
	if val != nil {
		t.Errorf("Actually: [%v],  Expected: [%v]", val, nil)
	}
	if resp == nil || resp["code"] != NeedArgErrorResponseCode {
		t.Errorf("Actually: [%v],  Expected: [%v]", resp, NewNeedArgErrorResponse("isMale", reflect.Bool))
	}
	t.Errorf("Test Error for Travis-CI\n")
}
