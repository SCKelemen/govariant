// +build amd64

// This file contains Variant implementation specific to GOARCH=amd64
package variant

import (
	"reflect"
	"unsafe"
)

type Variant struct {
	// Pointer to the slice start for slice-based types.
	ptr unsafe.Pointer

	// Len and Type fields.
	// Type uses `TypeFieldBitCount` least significant bits, Len uses the rest.
	// Len is used only for the slice-based types.
	lenAndType int

	// Capacity for slice-based types, or the value for other types. For Float64Val type
	// contains the 64 bits of the floating point value.
	capOrVal int
}

// NewInt creates a Variant of VTypeInt type.
func NewInt(v int) Variant {
	return Variant{
		lenAndType: int(VTypeInt),
		capOrVal:   v,
	}
}

// NewFloat64 creates a Variant of VTypeFloat64 type.
func NewFloat64(v float64) Variant {
	return Variant{
		lenAndType: int(VTypeFloat64),
		capOrVal:   *(*int)(unsafe.Pointer(&v)),
	}
}

// NewBytes creates a Variant of VTypeBytes type and initializes it with the specified
// slice of bytes.
func NewBytes(v []byte) Variant {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&v))
	if hdr.Len > MaxSliceLen {
		panic("maximum len exceeded")
	}

	return Variant{
		ptr:        unsafe.Pointer(hdr.Data),
		lenAndType: (hdr.Len << TypeFieldBitCount) | int(VTypeBytes),
		capOrVal:   hdr.Cap,
	}
}

// NewValueList creates a Variant of VTypeValueList type and initializes it with the
// specified slice of Variants.
func NewValueList(v []Variant) Variant {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&v))
	if hdr.Len > MaxSliceLen {
		panic("maximum len exceeded")
	}

	return Variant{
		ptr:        unsafe.Pointer(hdr.Data),
		lenAndType: (hdr.Len << TypeFieldBitCount) | int(VTypeValueList),
		capOrVal:   hdr.Cap,
	}
}

// NewKeyValueList creates a Variant of VTypeKeyValueList type and initializes it with the
// specified slice of KeyValues.
func NewKeyValueList(v []KeyValue) Variant {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&v))

	return Variant{
		ptr:        unsafe.Pointer(hdr.Data),
		lenAndType: (hdr.Len << TypeFieldBitCount) | int(VTypeKeyValueList),
		capOrVal:   hdr.Cap,
	}
}