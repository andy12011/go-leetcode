package leetcode

// definit in 12.IntegertoRoman
// var numberSymbolMap = [][]byte{
// 	{77},     // M
// 	{67, 77}, // CM
// 	{68},     // D
// 	{67, 68}, // CD
// 	{67},     // C
// 	{88, 67}, // XC
// 	{76},     // L
// 	{88, 76}, // XL
// 	{88},     // X
// 	{73, 88}, // IX
// 	{86},     // V
// 	{73, 86}, // IV
// 	{73},     // I
// }

// var number = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
//MMM
func romanToInt(s string) int {
	result := 0
	inputBytes := []byte(s)
	inputLength := len(inputBytes)

	compareIdx := 0
	compareBytes := numberSymbolMap[compareIdx]

	inputFlag := 0
	var inputVal byte

	for {
		inputVal = inputBytes[inputFlag]

		if inputVal == compareBytes[0] {
			if len(compareBytes) == 1 {
				result += number[compareIdx]
				inputFlag++
			} else {
				if inputFlag == inputLength-1 { //當我沒有下一位數的時候 num++
					compareIdx++
					compareBytes = numberSymbolMap[compareIdx]
				} else {
					nextVal := inputBytes[inputFlag+1]

					if nextVal == compareBytes[1] {
						result += number[compareIdx]
						inputFlag += 2
					} else {
						compareIdx++
						compareBytes = numberSymbolMap[compareIdx]
					}
				}
			}
		} else {
			compareIdx++
			compareBytes = numberSymbolMap[compareIdx]
		}

		if inputFlag == inputLength {
			break
		}
	}

	return result
}
