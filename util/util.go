package util

func RemoveDuplicates(slice []string) []string {
	var seen = make(map[string]bool)
	var result = []string{}

	var val string
	var ok bool

	for _, val = range slice {
		if _, ok = seen[val]; !ok {
			seen[val] = true
			result = append(result, val)
		}
	}
	return result
}
