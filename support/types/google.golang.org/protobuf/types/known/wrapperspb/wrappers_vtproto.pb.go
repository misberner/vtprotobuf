// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: (devel)
// source: google/protobuf/wrappers.proto

package wrapperspb

import (
	binary "encoding/binary"
	fmt "fmt"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	io "io"
	math "math"
	bits "math/bits"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func CloneVT_DoubleValue(m *wrapperspb.DoubleValue) *wrapperspb.DoubleValue {
	if m == nil {
		return (*wrapperspb.DoubleValue)(nil)
	}
	r := &wrapperspb.DoubleValue{
		Value: m.Value,
	}
	if uf := m.ProtoReflect().GetUnknown(); len(uf) > 0 {
		ufc := make([]byte, len(uf))
		copy(ufc, uf)
		r.ProtoReflect().SetUnknown(ufc)
	}
	return r
}

func CloneVT_FloatValue(m *wrapperspb.FloatValue) *wrapperspb.FloatValue {
	if m == nil {
		return (*wrapperspb.FloatValue)(nil)
	}
	r := &wrapperspb.FloatValue{
		Value: m.Value,
	}
	if uf := m.ProtoReflect().GetUnknown(); len(uf) > 0 {
		ufc := make([]byte, len(uf))
		copy(ufc, uf)
		r.ProtoReflect().SetUnknown(ufc)
	}
	return r
}

func CloneVT_Int64Value(m *wrapperspb.Int64Value) *wrapperspb.Int64Value {
	if m == nil {
		return (*wrapperspb.Int64Value)(nil)
	}
	r := &wrapperspb.Int64Value{
		Value: m.Value,
	}
	if uf := m.ProtoReflect().GetUnknown(); len(uf) > 0 {
		ufc := make([]byte, len(uf))
		copy(ufc, uf)
		r.ProtoReflect().SetUnknown(ufc)
	}
	return r
}

func CloneVT_UInt64Value(m *wrapperspb.UInt64Value) *wrapperspb.UInt64Value {
	if m == nil {
		return (*wrapperspb.UInt64Value)(nil)
	}
	r := &wrapperspb.UInt64Value{
		Value: m.Value,
	}
	if uf := m.ProtoReflect().GetUnknown(); len(uf) > 0 {
		ufc := make([]byte, len(uf))
		copy(ufc, uf)
		r.ProtoReflect().SetUnknown(ufc)
	}
	return r
}

func CloneVT_Int32Value(m *wrapperspb.Int32Value) *wrapperspb.Int32Value {
	if m == nil {
		return (*wrapperspb.Int32Value)(nil)
	}
	r := &wrapperspb.Int32Value{
		Value: m.Value,
	}
	if uf := m.ProtoReflect().GetUnknown(); len(uf) > 0 {
		ufc := make([]byte, len(uf))
		copy(ufc, uf)
		r.ProtoReflect().SetUnknown(ufc)
	}
	return r
}

func CloneVT_UInt32Value(m *wrapperspb.UInt32Value) *wrapperspb.UInt32Value {
	if m == nil {
		return (*wrapperspb.UInt32Value)(nil)
	}
	r := &wrapperspb.UInt32Value{
		Value: m.Value,
	}
	if uf := m.ProtoReflect().GetUnknown(); len(uf) > 0 {
		ufc := make([]byte, len(uf))
		copy(ufc, uf)
		r.ProtoReflect().SetUnknown(ufc)
	}
	return r
}

func CloneVT_BoolValue(m *wrapperspb.BoolValue) *wrapperspb.BoolValue {
	if m == nil {
		return (*wrapperspb.BoolValue)(nil)
	}
	r := &wrapperspb.BoolValue{
		Value: m.Value,
	}
	if uf := m.ProtoReflect().GetUnknown(); len(uf) > 0 {
		ufc := make([]byte, len(uf))
		copy(ufc, uf)
		r.ProtoReflect().SetUnknown(ufc)
	}
	return r
}

func CloneVT_StringValue(m *wrapperspb.StringValue) *wrapperspb.StringValue {
	if m == nil {
		return (*wrapperspb.StringValue)(nil)
	}
	r := &wrapperspb.StringValue{
		Value: m.Value,
	}
	if uf := m.ProtoReflect().GetUnknown(); len(uf) > 0 {
		ufc := make([]byte, len(uf))
		copy(ufc, uf)
		r.ProtoReflect().SetUnknown(ufc)
	}
	return r
}

func CloneVT_BytesValue(m *wrapperspb.BytesValue) *wrapperspb.BytesValue {
	if m == nil {
		return (*wrapperspb.BytesValue)(nil)
	}
	r := &wrapperspb.BytesValue{}
	if rhs := m.Value; rhs != nil {
		tmpBytes := make([]byte, len(rhs))
		copy(tmpBytes, rhs)
		r.Value = tmpBytes
	}
	if uf := m.ProtoReflect().GetUnknown(); len(uf) > 0 {
		ufc := make([]byte, len(uf))
		copy(ufc, uf)
		r.ProtoReflect().SetUnknown(ufc)
	}
	return r
}

func EqualVT_DoubleValue(this *wrapperspb.DoubleValue, that *wrapperspb.DoubleValue) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if this.Value != that.Value {
		return false
	}
	return string(this.ProtoReflect().GetUnknown()) == string(that.ProtoReflect().GetUnknown())
}
func EqualVT_FloatValue(this *wrapperspb.FloatValue, that *wrapperspb.FloatValue) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if this.Value != that.Value {
		return false
	}
	return string(this.ProtoReflect().GetUnknown()) == string(that.ProtoReflect().GetUnknown())
}
func EqualVT_Int64Value(this *wrapperspb.Int64Value, that *wrapperspb.Int64Value) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if this.Value != that.Value {
		return false
	}
	return string(this.ProtoReflect().GetUnknown()) == string(that.ProtoReflect().GetUnknown())
}
func EqualVT_UInt64Value(this *wrapperspb.UInt64Value, that *wrapperspb.UInt64Value) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if this.Value != that.Value {
		return false
	}
	return string(this.ProtoReflect().GetUnknown()) == string(that.ProtoReflect().GetUnknown())
}
func EqualVT_Int32Value(this *wrapperspb.Int32Value, that *wrapperspb.Int32Value) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if this.Value != that.Value {
		return false
	}
	return string(this.ProtoReflect().GetUnknown()) == string(that.ProtoReflect().GetUnknown())
}
func EqualVT_UInt32Value(this *wrapperspb.UInt32Value, that *wrapperspb.UInt32Value) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if this.Value != that.Value {
		return false
	}
	return string(this.ProtoReflect().GetUnknown()) == string(that.ProtoReflect().GetUnknown())
}
func EqualVT_BoolValue(this *wrapperspb.BoolValue, that *wrapperspb.BoolValue) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if this.Value != that.Value {
		return false
	}
	return string(this.ProtoReflect().GetUnknown()) == string(that.ProtoReflect().GetUnknown())
}
func EqualVT_StringValue(this *wrapperspb.StringValue, that *wrapperspb.StringValue) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if this.Value != that.Value {
		return false
	}
	return string(this.ProtoReflect().GetUnknown()) == string(that.ProtoReflect().GetUnknown())
}
func EqualVT_BytesValue(this *wrapperspb.BytesValue, that *wrapperspb.BytesValue) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if string(this.Value) != string(that.Value) {
		return false
	}
	return string(this.ProtoReflect().GetUnknown()) == string(that.ProtoReflect().GetUnknown())
}
func MarshalVT_DoubleValue(m *wrapperspb.DoubleValue) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := SizeVT_DoubleValue(m)
	dAtA = make([]byte, size)
	n, err := MarshalToSizedBufferVT_DoubleValue(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func MarshalToVT_DoubleValue(m *wrapperspb.DoubleValue, dAtA []byte) (int, error) {
	size := SizeVT_DoubleValue(m)
	return MarshalToSizedBufferVT_DoubleValue(m, dAtA[:size])
}

func MarshalToSizedBufferVT_DoubleValue(m *wrapperspb.DoubleValue, dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if uf := m.ProtoReflect().GetUnknown(); uf != nil {
		i -= len(uf)
		copy(dAtA[i:], uf)
	}
	if m.Value != 0 {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func MarshalVT_FloatValue(m *wrapperspb.FloatValue) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := SizeVT_FloatValue(m)
	dAtA = make([]byte, size)
	n, err := MarshalToSizedBufferVT_FloatValue(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func MarshalToVT_FloatValue(m *wrapperspb.FloatValue, dAtA []byte) (int, error) {
	size := SizeVT_FloatValue(m)
	return MarshalToSizedBufferVT_FloatValue(m, dAtA[:size])
}

func MarshalToSizedBufferVT_FloatValue(m *wrapperspb.FloatValue, dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if uf := m.ProtoReflect().GetUnknown(); uf != nil {
		i -= len(uf)
		copy(dAtA[i:], uf)
	}
	if m.Value != 0 {
		i -= 4
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Value))))
		i--
		dAtA[i] = 0xd
	}
	return len(dAtA) - i, nil
}

func MarshalVT_Int64Value(m *wrapperspb.Int64Value) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := SizeVT_Int64Value(m)
	dAtA = make([]byte, size)
	n, err := MarshalToSizedBufferVT_Int64Value(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func MarshalToVT_Int64Value(m *wrapperspb.Int64Value, dAtA []byte) (int, error) {
	size := SizeVT_Int64Value(m)
	return MarshalToSizedBufferVT_Int64Value(m, dAtA[:size])
}

func MarshalToSizedBufferVT_Int64Value(m *wrapperspb.Int64Value, dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if uf := m.ProtoReflect().GetUnknown(); uf != nil {
		i -= len(uf)
		copy(dAtA[i:], uf)
	}
	if m.Value != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func MarshalVT_UInt64Value(m *wrapperspb.UInt64Value) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := SizeVT_UInt64Value(m)
	dAtA = make([]byte, size)
	n, err := MarshalToSizedBufferVT_UInt64Value(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func MarshalToVT_UInt64Value(m *wrapperspb.UInt64Value, dAtA []byte) (int, error) {
	size := SizeVT_UInt64Value(m)
	return MarshalToSizedBufferVT_UInt64Value(m, dAtA[:size])
}

func MarshalToSizedBufferVT_UInt64Value(m *wrapperspb.UInt64Value, dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if uf := m.ProtoReflect().GetUnknown(); uf != nil {
		i -= len(uf)
		copy(dAtA[i:], uf)
	}
	if m.Value != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func MarshalVT_Int32Value(m *wrapperspb.Int32Value) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := SizeVT_Int32Value(m)
	dAtA = make([]byte, size)
	n, err := MarshalToSizedBufferVT_Int32Value(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func MarshalToVT_Int32Value(m *wrapperspb.Int32Value, dAtA []byte) (int, error) {
	size := SizeVT_Int32Value(m)
	return MarshalToSizedBufferVT_Int32Value(m, dAtA[:size])
}

func MarshalToSizedBufferVT_Int32Value(m *wrapperspb.Int32Value, dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if uf := m.ProtoReflect().GetUnknown(); uf != nil {
		i -= len(uf)
		copy(dAtA[i:], uf)
	}
	if m.Value != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func MarshalVT_UInt32Value(m *wrapperspb.UInt32Value) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := SizeVT_UInt32Value(m)
	dAtA = make([]byte, size)
	n, err := MarshalToSizedBufferVT_UInt32Value(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func MarshalToVT_UInt32Value(m *wrapperspb.UInt32Value, dAtA []byte) (int, error) {
	size := SizeVT_UInt32Value(m)
	return MarshalToSizedBufferVT_UInt32Value(m, dAtA[:size])
}

func MarshalToSizedBufferVT_UInt32Value(m *wrapperspb.UInt32Value, dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if uf := m.ProtoReflect().GetUnknown(); uf != nil {
		i -= len(uf)
		copy(dAtA[i:], uf)
	}
	if m.Value != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func MarshalVT_BoolValue(m *wrapperspb.BoolValue) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := SizeVT_BoolValue(m)
	dAtA = make([]byte, size)
	n, err := MarshalToSizedBufferVT_BoolValue(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func MarshalToVT_BoolValue(m *wrapperspb.BoolValue, dAtA []byte) (int, error) {
	size := SizeVT_BoolValue(m)
	return MarshalToSizedBufferVT_BoolValue(m, dAtA[:size])
}

func MarshalToSizedBufferVT_BoolValue(m *wrapperspb.BoolValue, dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if uf := m.ProtoReflect().GetUnknown(); uf != nil {
		i -= len(uf)
		copy(dAtA[i:], uf)
	}
	if m.Value {
		i--
		if m.Value {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func MarshalVT_StringValue(m *wrapperspb.StringValue) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := SizeVT_StringValue(m)
	dAtA = make([]byte, size)
	n, err := MarshalToSizedBufferVT_StringValue(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func MarshalToVT_StringValue(m *wrapperspb.StringValue, dAtA []byte) (int, error) {
	size := SizeVT_StringValue(m)
	return MarshalToSizedBufferVT_StringValue(m, dAtA[:size])
}

func MarshalToSizedBufferVT_StringValue(m *wrapperspb.StringValue, dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if uf := m.ProtoReflect().GetUnknown(); uf != nil {
		i -= len(uf)
		copy(dAtA[i:], uf)
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarint(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func MarshalVT_BytesValue(m *wrapperspb.BytesValue) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := SizeVT_BytesValue(m)
	dAtA = make([]byte, size)
	n, err := MarshalToSizedBufferVT_BytesValue(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func MarshalToVT_BytesValue(m *wrapperspb.BytesValue, dAtA []byte) (int, error) {
	size := SizeVT_BytesValue(m)
	return MarshalToSizedBufferVT_BytesValue(m, dAtA[:size])
}

func MarshalToSizedBufferVT_BytesValue(m *wrapperspb.BytesValue, dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if uf := m.ProtoReflect().GetUnknown(); uf != nil {
		i -= len(uf)
		copy(dAtA[i:], uf)
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarint(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarint(dAtA []byte, offset int, v uint64) int {
	offset -= sov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func SizeVT_DoubleValue(m *wrapperspb.DoubleValue) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 9
	}
	n += len(m.ProtoReflect().GetUnknown())
	return n
}

func SizeVT_FloatValue(m *wrapperspb.FloatValue) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 5
	}
	n += len(m.ProtoReflect().GetUnknown())
	return n
}

func SizeVT_Int64Value(m *wrapperspb.Int64Value) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sov(uint64(m.Value))
	}
	n += len(m.ProtoReflect().GetUnknown())
	return n
}

func SizeVT_UInt64Value(m *wrapperspb.UInt64Value) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sov(uint64(m.Value))
	}
	n += len(m.ProtoReflect().GetUnknown())
	return n
}

func SizeVT_Int32Value(m *wrapperspb.Int32Value) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sov(uint64(m.Value))
	}
	n += len(m.ProtoReflect().GetUnknown())
	return n
}

func SizeVT_UInt32Value(m *wrapperspb.UInt32Value) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sov(uint64(m.Value))
	}
	n += len(m.ProtoReflect().GetUnknown())
	return n
}

func SizeVT_BoolValue(m *wrapperspb.BoolValue) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value {
		n += 2
	}
	n += len(m.ProtoReflect().GetUnknown())
	return n
}

func SizeVT_StringValue(m *wrapperspb.StringValue) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.ProtoReflect().GetUnknown())
	return n
}

func SizeVT_BytesValue(m *wrapperspb.BytesValue) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.ProtoReflect().GetUnknown())
	return n
}

func sov(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}
func soz(x uint64) (n int) {
	return sov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func UnmarshalVT_DoubleValue(m *wrapperspb.DoubleValue, dAtA []byte) error {
	unknownFields := m.ProtoReflect().GetUnknown()
	unknownFieldsPreLen := len(unknownFields)
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DoubleValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DoubleValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Value = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			unknownFields = append(unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if len(unknownFields) > unknownFieldsPreLen {
		m.ProtoReflect().SetUnknown(unknownFields)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func UnmarshalVT_FloatValue(m *wrapperspb.FloatValue, dAtA []byte) error {
	unknownFields := m.ProtoReflect().GetUnknown()
	unknownFieldsPreLen := len(unknownFields)
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FloatValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FloatValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Value = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			unknownFields = append(unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if len(unknownFields) > unknownFieldsPreLen {
		m.ProtoReflect().SetUnknown(unknownFields)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func UnmarshalVT_Int64Value(m *wrapperspb.Int64Value, dAtA []byte) error {
	unknownFields := m.ProtoReflect().GetUnknown()
	unknownFieldsPreLen := len(unknownFields)
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Int64Value: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Int64Value: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			unknownFields = append(unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if len(unknownFields) > unknownFieldsPreLen {
		m.ProtoReflect().SetUnknown(unknownFields)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func UnmarshalVT_UInt64Value(m *wrapperspb.UInt64Value, dAtA []byte) error {
	unknownFields := m.ProtoReflect().GetUnknown()
	unknownFieldsPreLen := len(unknownFields)
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UInt64Value: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UInt64Value: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			unknownFields = append(unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if len(unknownFields) > unknownFieldsPreLen {
		m.ProtoReflect().SetUnknown(unknownFields)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func UnmarshalVT_Int32Value(m *wrapperspb.Int32Value, dAtA []byte) error {
	unknownFields := m.ProtoReflect().GetUnknown()
	unknownFieldsPreLen := len(unknownFields)
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Int32Value: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Int32Value: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			unknownFields = append(unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if len(unknownFields) > unknownFieldsPreLen {
		m.ProtoReflect().SetUnknown(unknownFields)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func UnmarshalVT_UInt32Value(m *wrapperspb.UInt32Value, dAtA []byte) error {
	unknownFields := m.ProtoReflect().GetUnknown()
	unknownFieldsPreLen := len(unknownFields)
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UInt32Value: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UInt32Value: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			unknownFields = append(unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if len(unknownFields) > unknownFieldsPreLen {
		m.ProtoReflect().SetUnknown(unknownFields)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func UnmarshalVT_BoolValue(m *wrapperspb.BoolValue, dAtA []byte) error {
	unknownFields := m.ProtoReflect().GetUnknown()
	unknownFieldsPreLen := len(unknownFields)
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BoolValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BoolValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			unknownFields = append(unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if len(unknownFields) > unknownFieldsPreLen {
		m.ProtoReflect().SetUnknown(unknownFields)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func UnmarshalVT_StringValue(m *wrapperspb.StringValue, dAtA []byte) error {
	unknownFields := m.ProtoReflect().GetUnknown()
	unknownFieldsPreLen := len(unknownFields)
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StringValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StringValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			unknownFields = append(unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if len(unknownFields) > unknownFieldsPreLen {
		m.ProtoReflect().SetUnknown(unknownFields)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func UnmarshalVT_BytesValue(m *wrapperspb.BytesValue, dAtA []byte) error {
	unknownFields := m.ProtoReflect().GetUnknown()
	unknownFieldsPreLen := len(unknownFields)
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BytesValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BytesValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			unknownFields = append(unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if len(unknownFields) > unknownFieldsPreLen {
		m.ProtoReflect().SetUnknown(unknownFields)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skip(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflow
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflow
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflow
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLength
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroup
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLength
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLength        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflow          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroup = fmt.Errorf("proto: unexpected end of group")
)
