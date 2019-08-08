package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32:
		fmt.Println("Float")
	case reflect.Int:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown")
	}
}

func TestBasicType(t *testing.T) {
	//var f float64 = 12
	var f float32 = 12
	CheckType(f)
}

// ------------------------------------------------------------------------------

type Employee struct {
	Id   string
	Name string `format:"normal"`
	Age  int
}

func (e *Employee) UpdateAge(newAge int) {
	e.Age = newAge
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"124", "Aheizi", 21}
	t.Logf("Name: value(%[1]v), Type(%[1]T) ", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'name' field")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Update Age:", e)
}

// reflect.ValueOf(*e).FieldByName("Name")
// reflect.ValueOf(*e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
