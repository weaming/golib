package abstract

func OrString(first string, two string) string {
	if IsTrue(first) {
		return first
	} else {
		return two
	}
}
