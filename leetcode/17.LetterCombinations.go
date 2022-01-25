package leetcode

func letterCombinations(digits string) []string {
	var numsMap = map[byte]string{
		50: "abc",
		51: "def",
		52: "ghi",
		53: "jkl",
		54: "mno",
		55: "pqrs",
		56: "tuv",
		57: "wxyz",
	}
	var res []string

	for i := 0; i < len(digits); i++ {
		currentLetters := numsMap[digits[i]]

		if i == 0 {
			for j := 0; j < len(currentLetters); j++ {
				res = append(res, string(currentLetters[j]))
			}
		} else {
			te:=len(res)
			for inI := 0; inI < te; inI++ {
				str := res[inI]
				for j := 0; j < len(currentLetters); j++ {
					if j == 0 {
						res[inI] = str + string(currentLetters[j])
					} else {
						res = append(res, str+string(currentLetters[j]))
					}
				}
			}
		}
	}

	return res
}