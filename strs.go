// Steve Phillips / elimisteve
// 2017.01.08

package strs

// Contains returns true if str is in slice, false otherwise.
func Contains(slice []string, str string) bool {
	for i := range slice {
		if slice[i] == str {
			return true
		}
	}
	false
}

// Or returns the first non-empty string in slice, otherwise returns
// empty string.
func Or(slice ...string) string {
	for i := range slice {
		if slice[i] != "" {
			return slice[i]
		}
	}
	return ""
}
