package abstract

func IsTrue(i interface{}) bool {
	switch i.(type) {
	case int, int32, int64, float32, float64:
		if i == 0 {
			return false
		}
		return true
	case string:
		if i == "" {
			return false
		}
		return true
	}
	return false
}

func IsFalse(i interface{}) bool {
	return !IsTrue(i)
}
