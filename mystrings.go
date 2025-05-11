package mystrings

func Reverse(s string) string { // We need to name the function starts with capital letter since it needs to be exported outside in go
	// Lowercase function name will not be exported outside in go that's we should use capital letter as starting letter

	reversedStr := ""
	for _, str := range s {
		reversedStr = string(str) + reversedStr
	}

	return reversedStr
}
