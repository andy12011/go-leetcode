package dynamicprograming

import (
	"fmt"
	"testing"
)

func Test_rodCutting(t *testing.T) {

	price := [1000]int{1, 5, 8, 9, 10, 17, 17, 20}
	fmt.Println(cutRod(price, 800))
}
