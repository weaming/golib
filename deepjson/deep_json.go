package deepjson

import (
	"fmt"
	"reflect"
	"strconv"
)

type Node struct {
	Path          []string
	Key           string
	Value         interface{}
	ContainerType string
}

func NewRootNode(value interface{}) Node {
	return Node{
		Path:          []string{},
		Key:           ".",
		Value:         value,
		ContainerType: "nil",
	}
}

func (r Node) String() string {
	return fmt.Sprintf("%v (%v)", r.Value, r.ContainerType)
}

func (r *Node) Get(key string) *Node {
	switch reflect.ValueOf(r.Value).Kind() {
	case reflect.Map:
		vv := r.Value.(map[string]interface{})
		if v, ok := vv[key]; ok {
			return &Node{
				Path:          append(r.Path, key),
				Key:           key,
				Value:         v,
				ContainerType: "map",
			}
		}
		return r.Nil(key, "map")

	case reflect.Slice:
		vv := r.Value.([]interface{})
		if i, e := strconv.Atoi(key); e == nil && i < len(vv) {
			return &Node{
				Path:          append(r.Path, key),
				Key:           key,
				Value:         vv[i],
				ContainerType: "list",
			}
		}
		return r.Nil(key, "list")
	}
	return r.Nil(key, "unknown")
}

func (r *Node) Nil(key string, ct string) *Node {
	return &Node{
		Path:          append(r.Path, key),
		Key:           key,
		Value:         nil,
		ContainerType: ct,
	}
}

func (r *Node) GetByPath(path []string) (rv *Node) {
	for _, i := range path {
		if rv != nil {
			rv = rv.Get(i)
		} else {
			rv = r.Get(i)
		}
	}
	return
}
