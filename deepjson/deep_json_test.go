package deepjson

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func Test1(t *testing.T) {
	data := "example.json"
	v := map[string]interface{}{}
	if c, e := ioutil.ReadFile(data); e == nil {
		json.Unmarshal(c, &v)
		node := NewRootNode(v)
		t.Log(node)

		t.Log(node.Get("a").Get("a"))
		t.Log(node.Get("b").Get("a"))
		t.Log(node.Get("b").Get("c"))

		t.Log(node.GetByPath([]string{"a", "a"}))
		t.Log(node.GetByPath([]string{"b", "a"}))
		t.Log(node.GetByPath([]string{"b", "c"}))

		t.Log(node.GetByPath([]string{"b", "a", "2"}))
	} else {
		t.Error(e)
	}
}
