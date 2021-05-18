package jnfo

import (
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	jnfo, err := New(os.Getenv("JNFO_TEST_URL"))
	if err != nil {
		t.Fatal(err)
	}

	v := reflect.ValueOf(*jnfo)
	ty := v.Type()
	kv := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		kv[ty.Field(i).Name] = v.Field(i).Interface()
	}

	t.Log(kv)
}
