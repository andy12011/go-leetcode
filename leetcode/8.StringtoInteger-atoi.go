package leetcode

//From https://leetcode.com/problems/string-to-integer-atoi/
import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

const ONE byte = 49
const ZERO byte = 48
const NINE byte = 57
const PLUS byte = 43
const MINUS byte = 45
const SPACE byte = 32
const MAX int = math.MaxInt32

func myAtoi(s string) int {
	var butter bytes.Buffer
	begin := false
	var prefix byte = 0
	var i int
	//hasWord := false
	//數字
	//空格加減號
	//其他
	//   -42
	for i = 0; i < len(s); i++ {
		if s[i] != PLUS && s[i] != MINUS && s[i] != SPACE && (s[i] < ZERO || s[i] > NINE) {
			return 0
		}

		if s[i] == PLUS || s[i] == MINUS {
			if i < len(s) - 1 && s[i+1] >= ZERO && s[i+1] <= NINE {
				continue
			} else {
				return 0
			}
		}

		if s[i] == SPACE {
			continue
		}

		if s[i] >= ZERO && s[i] <= NINE {
			if i > 0 { //數字開頭
				if s[i-1] == SPACE {
					begin = true
					break
				}
				if i == 1 && (s[i-1] == MINUS || s[i-1] == PLUS) {
					begin = true
					prefix = s[i-1]
					break
				}

				fmt.Println(i, string(s[i]), (s[i-1] == MINUS || s[i-1] == PLUS))
				if (s[i-1] == MINUS || s[i-1] == PLUS) && i > 1 && s[i-2] == SPACE {
					begin = true
					prefix = s[i-1]
					break
				}
			} else {

				begin = true
				break
			}
		}
	}
	if begin {
		strConnection(&butter, s[i])
		for true {
			i++
			if i > len(s)-1 {
				break
			}
			if s[i] >= ZERO && s[i] <= NINE {
				strConnection(&butter, s[i])
			} else {
				break
			}
		}

		ans, _ := strconv.Atoi(butter.String())
		if ans > MAX {
			if prefix == MINUS {
				return -(MAX + 1)
			}
			return MAX
		} else {
			if prefix == MINUS {
				return -ans
			}
			return ans
		}
	}

	return 0
}
