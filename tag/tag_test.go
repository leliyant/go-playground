package tag

import (
	"reflect"
	"testing"
)

func TestStructTag(t *testing.T) {
	type S struct {
		F string `species:"gopher" color:"blue"`
	}

	s := S{}
	st := reflect.TypeOf(s)
	field := st.Field(0)

	if field.Tag.Get("color") != "blue" {
		t.Fail()
	}

	if field.Tag.Get("species") != "gopher" {
		t.Fail()
	}
}
