package abstract

func IfElse(a, b, c interface{}) interface{} {
	if IsTrue(a) {
		return b
	} else {
		return c
	}
}
