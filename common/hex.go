package common

func Uint16ToBytes(n uint16) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
	}
}

func BytesToUint16(array []byte) uint16 {
	var data uint16 = 0
	for i := 0; i < len(array); i++ {
		data = data + uint16(uint(array[i])<<uint(8*i))
	}
	return data
}
