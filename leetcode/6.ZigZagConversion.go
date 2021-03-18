package leetcode

import "bytes"
//
//From https://leetcode.com/problems/zigzag-conversion/
func convert(s string, numRows int) string {
	var butter bytes.Buffer
	length := len(s)
	var distance int
	if numRows > 1 {
		distance = (numRows - 1) * 2
	} else {
		distance = 1
	}
	for i := 0; i < numRows; i++ {
		start := i
		if start >= length {
			break
		}
		strConnection(&butter, s[start])
		even := distance - i*2
		odd := distance - even

		first := true
		for true {
			if first {
				first = false
				if start = start + even; start < length {
					if even == 0 {
						continue
					}
					strConnection(&butter, s[start])
					continue
				}
				break
			} else {
				first = true
				if start = start + odd; start < length {
					if odd == 0 {
						continue
					}
					strConnection(&butter, s[start])
					continue
				}
				break
			}
		}
	}
	return butter.String()
}

func strConnection(buffer *bytes.Buffer, inffix uint8) {
	buffer.WriteByte(inffix)
}
