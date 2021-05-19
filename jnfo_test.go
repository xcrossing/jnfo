package jnfo

import (
	"encoding/json"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	jnfo, err := New(os.Getenv("JNFO_TEST_URL"))
	if err != nil {
		t.Fatal(err)
	}

	bytes, err := json.MarshalIndent(jnfo, "", "  ")
	t.Log(string(bytes))
	t.Log(jnfo.NumCastPicName())
}
