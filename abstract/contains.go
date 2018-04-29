package abstract

func Contains(list []interface{}, other interface{}) int {
	for i, a := range list {
		switch a.(type) {
		case int:
			if v, ok := other.(int); ok {
				if a == v {
					return i
				}
			}
		case int32:
			if v, ok := other.(int32); ok {
				if a == v {
					return i
				}
			}
		case int64:
			if v, ok := other.(int64); ok {
				if a == v {
					return i
				}
			}
		case float32:
			if v, ok := other.(float32); ok {
				if a == v {
					return i
				}
			}
		case float64:
			if v, ok := other.(float64); ok {
				if a == v {
					return i
				}
			}
		case string:
			if v, ok := other.(string); ok {
				if a == v {
					return i
				}
			}
		}
	}
	return -1
}
