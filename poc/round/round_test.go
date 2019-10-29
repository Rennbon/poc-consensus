package round

import "testing"

func Test_calcScoopNumber(t *testing.T) {
	r := &Round{}
	signature := make([]byte, 32)
	signature[0] = 1
	signature[1] = 2
	signature[2] = 3
	signature[3] = 4
	signature[4] = 5
	scoop := r.calcScoopNumber(1, signature)
	t.Log(scoop)
}

func Test_calculateResult(t *testing.T) {
	r := &Round{}
	r.calculateResult()

}
