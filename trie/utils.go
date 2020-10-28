package trie

import "math"

// getCommonPrefix returns the longest commong prefix between strings a and b
func getCommonPrefix(a, b string) string {
	prefix := make([]byte, 0, int(math.Min(float64(len(a)), float64(len(b)))))

	for i := 0; float64(i) < math.Min(float64(len(a)), float64(len(b))); i++ {
		// if the characters at index i don't match, return the prefix so far
		if a[i] != b[i] {
			return string(prefix)
		}

		prefix = append(prefix, a[i])
	}

	return string(prefix)
}
