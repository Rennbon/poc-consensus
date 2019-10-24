package sink

import (
	"errors"
	"math/big"
	"reflect"

	"github.com/rennbon/consensus/common"
)

const (
	BigInt byte = 0x01
	Bytes  byte = 0x02

	maxU16 = int(^uint16(0))
)

var (
	ERR_SINK_TYPE_DIFF = errors.New("sink type mismatch.")
)

type ComplexType struct {
	Size  [2]byte
	MType byte
	Data  []byte
}

func ZeroCopySourceRelease(source *ZeroCopySource, typ reflect.Type) (out interface{}, irregular, eof bool) {
	eof = true
	switch typ {
	case common.BType_Bool:
		out, irregular, eof = source.NextBool()
	case common.BType_Uint8:
		out, eof = source.NextUint8()
	case common.BType_Uint16:
		out, eof = source.NextUint16()
	case common.BType_Uint32:
		out, eof = source.NextUint32()
	case common.BType_Uint64:
		out, eof = source.NextUint64()
	case common.BType_BigInt:
		bg := &ComplexType{}
		bg, eof = source.NextComplex()
		out, _ = bg.ComplexToBigInt()
	case common.BType_Bytes:
		bg := &ComplexType{}
		bg, eof = source.NextComplex()
		out, _ = bg.ComplexToBytes()
	}
	return
}

func ZeroCopySinkAppend(sk *ZeroCopySink, value reflect.Value) error {
	switch value.Interface().(type) {
	case bool:
		sk.WriteBool(value.Interface().(bool))
	case uint8:
		sk.WriteUint8(value.Interface().(uint8))
	case uint16:
		sk.WriteUint16(value.Interface().(uint16))
	case uint32:
		sk.WriteUint32(value.Interface().(uint32))
	case uint64:
		sk.WriteUint64(value.Interface().(uint64))
	case *big.Int:
		cpx, err := BigIntToComplex(value.Interface().(*big.Int))
		if err != nil {
			return err
		}
		sk.WriteComplex(cpx)
	/*case common.Hash:
	sk.WriteHash(value.Interface().(common.Hash))*/
	case []byte:
		arr := value.Interface().([]byte)
		cpx, err := BytesToComplex(arr)
		if err != nil {
			return err
		}
		sk.WriteComplex(cpx)
	}
	return nil
}

func DataCheck(data []byte) (size [2]byte, err error) {
	l := len(data)
	if l > int(^uint16(0)) {
		err = errors.New("data overflow uint16.")

	} else {
		u16 := common.Uint16ToBytes(uint16(l))
		size = [2]byte{u16[0], u16[1]}
	}
	return
}
func BigIntToComplex(int *big.Int) (*ComplexType, error) {
	if int == nil {
		int = big.NewInt(0)
	}
	size, err := DataCheck(int.Bytes())
	if err != nil {
		return nil, err
	}
	ct := &ComplexType{
		Size:  size,
		MType: BigInt,
		Data:  int.Bytes(),
	}
	return ct, nil
}

func (cpx *ComplexType) ComplexToBigInt() (*big.Int, error) {
	if cpx.MType != BigInt {
		return nil, ERR_SINK_TYPE_DIFF
	}
	return new(big.Int).SetBytes(cpx.Data), nil
}

func BytesToComplex(bytes []byte) (*ComplexType, error) {
	size, err := DataCheck(bytes)
	if err != nil {
		return nil, err
	}
	ct := &ComplexType{
		Size:  size,
		MType: Bytes,
		Data:  bytes,
	}
	return ct, nil
}
func (cpx *ComplexType) ComplexToBytes() ([]byte, error) {
	if cpx.MType != Bytes {
		return nil, ERR_SINK_TYPE_DIFF
	}
	return cpx.Data, nil
}
