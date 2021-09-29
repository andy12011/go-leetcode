package leetcode

import "bytes"

func longestCommonPrefix(strs []string) string {
	var buffer bytes.Buffer

	if len(strs) == 0 {
		return buffer.String()
	}

	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		same := true
		for _, s := range strs {
			if i >= len(s) || s[i] != char {
				same = false
				break
			}
		}

		if !same {
			break
		}

		buffer.WriteByte(char)
	}

	return buffer.String()
}
