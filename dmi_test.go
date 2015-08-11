package dmijson

import (
	"io/ioutil"
	"testing"
)

func TestJSON(t *testing.T) {
	data, err := ioutil.ReadFile("test.json")
	if err != nil {
		t.Fatal(err)
	}
	j, err := FromJSON(string(data))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(j.ToJSON())
}

func TestScript(t *testing.T) {
	script, err := Script()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(script)
}
