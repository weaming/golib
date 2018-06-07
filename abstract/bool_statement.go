package abstract

func Or(a, b interface{}) interface{} {
	if IsTrue(a) {
		return a
	} else {
		return b
	}
}

func And(a, b interface{}) interface{} {
	if IsTrue(a) && IsTrue(b) {
		return b
	} else {
		return a
	}
}

func OrString(a, b string) string {
	return Or(a, b).(string)
}

func AndString(a, b string) string {
	return And(a, b).(string)
}
