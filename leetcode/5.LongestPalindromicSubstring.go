package leetcode


func longestPalindrome(s string) string {
	count := len(s)
	maxIndex := count - 1
	if count == 0 {
		return ""
	}
	if count == 1 {
		return string(s[0])
	}

	p := 0
	l := p
	r := p + 1
	ans := string(s[0])
	tmp := ""
	for true {
		if (maxIndex - p) * 2 < len(ans) {
			break
		}
		if s[l] == s[r] {
			for true {
				if hasRight(&r, &maxIndex) && s[l] == s[r+1] {
					r++
					continue
				} else {
					break
				}
			}

		} else if hasRight(&r, &maxIndex) && s[l] == s[r+1] {
			r++

		} else {

			if p+2 < maxIndex {
				p++
				l = p
				r = p + 1
				continue
			} else {
				if hasRight(&r, &maxIndex) && s[r] == s[r+1] && len(ans) < 2 {
					ans = s[r:r+2]
				}
				break
			}

		}

		for true {
			if hasLeft(&l) && hasRight(&r, &maxIndex) {
				if s[l-1] == s[r+1] {
					l--
					r++
				} else {
					break;
				}
				continue
			} else {
				break
			}
		}
		tmp = s[l : r+1]

		if len(tmp) > len(ans) {
			ans = tmp
		}

		if p+2 < maxIndex {
			p++
			l = p
			r = p + 1
			continue
		} else {
			break
		}

	}

	return ans
}


func hasLeft(l *int) bool {
	if *l > 0 {
		return true
	}
	return false
}

func hasRight(r *int, maxIndex *int) bool {
	if *r < *maxIndex {
		return true
	}
	return false
}

