package default_map

import (
	"sort"
)

type DefaultMap struct {
	Map map[string]interface{}
	fn  func(index string) interface{}
}

func NewDefaultMap(fn func(index string) interface{}) *DefaultMap {
	if fn == nil {
		fn = func(index string) interface{} { return nil }
	}
	return &DefaultMap{fn: fn, Map: map[string]interface{}{}}
}

func (p *DefaultMap) Set(index string, value interface{}) (old interface{}) {
	if v, ok := p.Map[index]; ok {
		old = v
	}
	p.Map[index] = value
	return
}

func (p *DefaultMap) Get(index string) (rv interface{}) {
	if v, ok := p.Map[index]; ok {
		rv = v
	}
	return
}

func (p *DefaultMap) GetOrDefault(index string) (rv interface{}) {
	if v, ok := p.Map[index]; ok {
		rv = v
	} else {
		rv = p.fn(index)
	}
	return
}

func (p *DefaultMap) GetOrSetDefault(index string) (rv interface{}) {
	if v, ok := p.Map[index]; ok {
		rv = v
	} else {
		rv = p.fn(index)
		p.Map[index] = rv
	}
	return
}

func (p *DefaultMap) GetOrSet(index string, value interface{}) (rv interface{}) {
	if v, ok := p.Map[index]; ok {
		rv = v
	} else {
		rv = value
		p.Map[index] = value
	}
	return
}

func (p *DefaultMap) Sorted(reverse bool) []interface{} {
	keys := p.SortedKeys(reverse)
	var values []interface{}
	for _, k := range keys {
		values = append(values, p.Map[k])
	}
	return values
}

func (p *DefaultMap) SortedKeys(reverse bool) []string {
	var keys []string
	for k := range p.Map {
		keys = append(keys, k)
	}
	if reverse {
		sort.Slice(keys, func(i, j int) bool { return keys[i] > keys[j] })
	} else {
		sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	}
	return keys
}
