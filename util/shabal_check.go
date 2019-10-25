package util

type shabalCheck struct {
}

func NewShabalCheck() {

}
func (o *shabalCheck) FindBestDeadline(scoops, gensig []byte, numScoops int) int {
	return 0
}
func (o *shabalCheck) FindLowest(gensig, data []byte) int {
	return o.FindBestDeadline(data, gensig, len(data)/SCOOP_SIZE)
}
