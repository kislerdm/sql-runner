package helper

// InArrayStr function to check if a string element is present in an array of string
func InArrayStr(array []string, testElement string) bool {
	for _, s := range array {
		if testElement == s {
			return true
		}
	}
	return false
}
