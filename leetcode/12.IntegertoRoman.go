package leetcode

var numberSymbolMap = [][]byte{
	{77},     // M
	{67, 77}, // CM
	{68},     // D
	{67, 68}, // CD
	{67},     // C
	{88, 67}, // XC
	{76},     // L
	{88, 76}, // XL
	{88},     // X
	{73, 88}, // IX
	{86},     // V
	{73, 86}, // IV
	{73},     // I
}

var number = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var numLen = len(number)

func intToRoman(num int) string {
	var resultByte []byte
	flag := 0
	val := number[flag]

	for {
		if num >= val {
			num -= val
			resultByte = append(resultByte, numberSymbolMap[flag]...)
		} else  {
			if num == 0 {
				break
			}
			flag++
			val = number[flag]
		}
	}

	return string(resultByte)
}
