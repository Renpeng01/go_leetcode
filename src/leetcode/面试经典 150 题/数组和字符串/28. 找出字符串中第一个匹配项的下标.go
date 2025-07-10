package main

func strStr(haystack string, needle string) int {
	subLen := len(needle)

	for i := subLen - 1; i < len(haystack); i++ {
		if haystack[i-(subLen-1):i+1] == needle {
			return i - (subLen - 1)
		}
	}
	return -1
}
