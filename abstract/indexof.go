package abstract

func IndexOf(list []interface{}, other interface{}) int {
	for i, a := range list {
		if a == other {
			return i
		}
	}
	return -1
}
