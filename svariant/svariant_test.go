package svariant

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/tigrannajaryan/govariant/testutil"

	"github.com/tigrannajaryan/govariant/uvariant"

	"github.com/stretchr/testify/assert"
)

func TestUSVariant(t *testing.T) {
	fmt.Printf("SVariant size=%v\n", unsafe.Sizeof(SVariant{}))

	b1 := []byte{1, 2, 3}
	v := BytesSVariant(b1)
	b2 := v.Bytes()
	assert.EqualValues(t, b1, b2)
	assert.EqualValues(t, uvariant.VariantTypeBytes, v.Type())

	s1 := "abcdef"
	v = StringSVariant(s1)
	s2 := v.String()
	assert.EqualValues(t, s1, s2)
	assert.EqualValues(t, uvariant.VariantTypeString, v.Type())

	i1 := 1234
	v = IntSVariant(i1)
	i2 := v.Int()
	assert.EqualValues(t, i1, i2)
	assert.EqualValues(t, uvariant.VariantTypeInt, v.Type())

	f1 := 1234.567
	v = Float64SVariant(f1)
	f2 := v.Float64()
	assert.EqualValues(t, f1, f2)
	assert.EqualValues(t, uvariant.VariantTypeFloat64, v.Type())

	//assert.EqualValues(t, 8, unsafe.Sizeof(int(123)))
}

func createSVariantInt() SVariant {
	for i := 0; i < 1; i++ {
		return IntSVariant(testutil.IntMagicVal)
	}
	return IntSVariant(0)
}

func createSVariantFloat64() SVariant {
	for i := 0; i < 1; i++ {
		return Float64SVariant(testutil.Float64MagicVal)
	}
	return Float64SVariant(0)
}

func createSVariantString() SVariant {
	for i := 0; i < 1; i++ {
		return StringSVariant(testutil.StrMagicVal)
	}
	return StringSVariant("def")
}

func createSVariantBytes() SVariant {
	for i := 0; i < 1; i++ {
		return BytesSVariant(testutil.BytesMagicVal)
	}
	return BytesSVariant(testutil.BytesMagicVal)
}

func BenchmarkSimpleVariantIntGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := createSVariantInt()
		vi := v.Int()
		if vi != vi {
			panic("invalid value")
		}
	}
}

func BenchmarkSimpleVariantFloat64Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := createSVariantFloat64()
		vf := v.Float64()
		if vf != testutil.Float64MagicVal {
			panic("invalid value")
		}
	}
}

func BenchmarkSimpleVariantStringTypeAndGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := createSVariantString()
		if v.Type() == uvariant.VariantTypeString {
			if v.String() == "" {
				panic("empty string")
			}
		} else {
			panic("invalid type")
		}
	}
}

func BenchmarkSimpleVariantBytesTypeAndGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := createSVariantBytes()
		if v.Type() == uvariant.VariantTypeBytes {
			if v.Bytes() == nil {
				panic("nil bytes")
			}
		} else {
			panic("invalid type")
		}
	}
}

func BenchmarkSimpleVariantIntTypeAndGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := createSVariantInt()
		if v.Type() == uvariant.VariantTypeInt {
			vi := v.Int()
			if vi != vi {
				panic("invalid value")
			}
		} else {
			panic("invalid type")
		}
	}
}

func createSVariantIntSlice(n int) []SVariant {
	v := make([]SVariant, n)
	for i := 0; i < n; i++ {
		v[i] = IntSVariant(i)
	}
	return v
}

func createSVariantStringSlice(n int) []SVariant {
	v := make([]SVariant, n)
	for i := 0; i < n; i++ {
		v[i] = StringSVariant("abc")
	}
	return v
}

func BenchmarkSimpleVariantSliceIntGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vv := createSVariantIntSlice(testutil.VariantSliceSize)
		for _, v := range vv {
			if v.Int() < 0 {
				panic("zero int")
			}
		}
	}
}

func BenchmarkSimpleVariantSliceIntTypeAndGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vv := createSVariantIntSlice(testutil.VariantSliceSize)
		for _, v := range vv {
			if v.Type() == uvariant.VariantTypeInt {
				if v.Int() < 0 {
					panic("zero int")
				}
			} else {
				panic("invalid type")
			}
		}
	}
}

func BenchmarkSimpleVariantSliceStringTypeAndGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vv := createSVariantStringSlice(testutil.VariantSliceSize)
		for _, v := range vv {
			if v.Type() == uvariant.VariantTypeString {
				if v.String() == "" {
					panic("empty string")
				}
			} else {
				panic("invalid type")
			}
		}
	}
}