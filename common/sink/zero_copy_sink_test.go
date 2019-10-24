/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */
package sink

import (
	"testing"

	"bytes"

	"math/big"

	"github.com/stretchr/testify/assert"
	//"github.com/rennbon/common"
)

func TestSourceSink(t *testing.T) {
	/*	addressFs1, _ := common.ToAddress("bx1qK96vAkK6E8S7JgYUY3YY28Qhj6cmfdy")
		addressFs2, _ := common.ToAddress("bx1qK96vAkK6E8S7JgYUY3YY28Qhj6cmfdz")
		bint, _ := big.NewInt(0).SetString("1232222222222222222222222222222222222222222222", 10)

		txamounts := &common.TxOuts{
			Tos: []*common.TxOut{{
				Address: addressFs1,
				Amount:  bint,
			},
				{
					Address: addressFs2,
					Amount:  big.NewInt(300),
				},
			},
		}*/
	bigInt := big.NewInt(1000)

	a3 := uint8(100)
	a4 := uint16(65535)
	a5 := uint32(4294967295)
	a6 := uint64(18446744073709551615)
	a7 := uint64(18446744073709551615)
	a8 := []byte{10, 11, 12}
	a9 := "hello onchain."
	/*a10, _ := TxOutsToComplex(txamounts)*/
	a11, _ := BigIntToComplex(bigInt)

	sink := NewZeroCopySink(nil)
	sink.WriteByte(a3)
	sink.WriteUint16(a4)
	sink.WriteUint32(a5)
	sink.WriteUint64(a6)
	sink.WriteVarUint(a7)
	sink.WriteVarBytes(a8)
	sink.WriteString(a9)
	/*	sink.WriteComplex(a10)
		sink.WriteComplex(a11)*/

	source := NewZeroCopySource(sink.Bytes())
	b3, _ := source.NextByte()
	assert.Equal(t, a3, b3)
	b4, _ := source.NextUint16()
	assert.Equal(t, a4, b4)
	b5, _ := source.NextUint32()
	assert.Equal(t, a5, b5)
	b6, _ := source.NextUint64()
	assert.Equal(t, a6, b6)
	b7, _, _, _ := source.NextVarUint()
	assert.Equal(t, a7, b7)
	b8, _, _, _ := source.NextVarBytes()
	assert.Equal(t, a8, b8)
	b9, _, _, _ := source.NextString()
	assert.Equal(t, a9, b9)
	/*b10, _ := source.NextComplex()
	assert.Equal(t, a10, b10)*/
	b11, _ := source.NextComplex()
	assert.Equal(t, a11, b11)
}

func BenchmarkNewZeroCopySink_Serialize(b *testing.B) {
	N := 1000
	a3 := uint8(100)
	a4 := uint16(65535)
	a5 := uint32(4294967295)
	a6 := uint64(18446744073709551615)
	a7 := uint64(18446744073709551615)
	a8 := []byte{10, 11, 12}
	a9 := "hello onchain."
	buf := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		for j := 0; j < N; j++ {
			WriteVarUint(buf, uint64(a3))
			WriteVarUint(buf, uint64(a4))
			WriteVarUint(buf, uint64(a5))
			WriteVarUint(buf, uint64(a6))
			WriteVarUint(buf, uint64(a7))
			WriteVarBytes(buf, a8)
			WriteString(buf, a9)

			buf.WriteByte(20)
			buf.WriteByte(21)
			buf.WriteByte(22)
		}
	}
}

func BenchmarkZeroCopySink(ben *testing.B) {
	N := 1000
	a3 := uint8(100)
	a4 := uint16(65535)
	a5 := uint32(4294967295)
	a6 := uint64(18446744073709551615)
	a7 := uint64(18446744073709551615)
	a8 := []byte{10, 11, 12}
	a9 := "hello onchain."
	sink := NewZeroCopySink(nil)
	for i := 0; i < ben.N; i++ {
		sink.Reset()
		for j := 0; j < N; j++ {
			sink.WriteVarUint(uint64(a3))
			sink.WriteVarUint(uint64(a4))
			sink.WriteVarUint(uint64(a5))
			sink.WriteVarUint(uint64(a6))
			sink.WriteVarUint(uint64(a7))
			sink.WriteVarBytes(a8)
			sink.WriteString(a9)
			sink.WriteByte(20)
			sink.WriteByte(21)
			sink.WriteByte(22)
		}
	}

}
