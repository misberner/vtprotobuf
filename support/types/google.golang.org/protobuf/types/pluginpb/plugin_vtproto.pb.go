// Code generated by protoc-gen-go-vtproto-builtins. DO NOT EDIT.
// protoc-gen-go-vtproto-builtins version: (devel)
// source: google/protobuf/compiler/plugin.proto

package pluginpb

import (
	fmt "fmt"
	descriptorpb1 "github.com/planetscale/vtprotobuf/support/types/google.golang.org/protobuf/types/descriptorpb"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	pluginpb "google.golang.org/protobuf/types/pluginpb"
	io "io"
	bits "math/bits"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func CloneVT_Version(m *pluginpb.Version) *pluginpb.Version {
	if m == nil {
		return (*pluginpb.Version)(nil)
	}
	r := &pluginpb.Version{}
	if rhs := m.Major; rhs != nil {
		tmpVal := *rhs
		r.Major = &tmpVal
	}
	if rhs := m.Minor; rhs != nil {
		tmpVal := *rhs
		r.Minor = &tmpVal
	}
	if rhs := m.Patch; rhs != nil {
		tmpVal := *rhs
		r.Patch = &tmpVal
	}
	if rhs := m.Suffix; rhs != nil {
		tmpVal := *rhs
		r.Suffix = &tmpVal
	}
	if uf := m.ProtoReflect().GetUnknown(); len(uf) > 0 {
		ufc := make([]byte, len(uf))
		copy(ufc, uf)
		r.ProtoReflect().SetUnknown(ufc)
	}
	return r
}

func CloneVT_CodeGeneratorRequest(m *pluginpb.CodeGeneratorRequest) *pluginpb.CodeGeneratorRequest {
	if m == nil {
		return (*pluginpb.CodeGeneratorRequest)(nil)
	}
	r := &pluginpb.CodeGeneratorRequest{
		CompilerVersion: CloneVT_Version(m.CompilerVersion),
	}
	if rhs := m.FileToGenerate; rhs != nil {
		tmpContainer := make([]string, len(rhs))
		copy(tmpContainer, rhs)
		r.FileToGenerate = tmpContainer
	}
	if rhs := m.Parameter; rhs != nil {
		tmpVal := *rhs
		r.Parameter = &tmpVal
	}
	if rhs := m.ProtoFile; rhs != nil {
		tmpContainer := make([]*descriptorpb.FileDescriptorProto, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = descriptorpb1.CloneVT_FileDescriptorProto(v)
		}
		r.ProtoFile = tmpContainer
	}
	if uf := m.ProtoReflect().GetUnknown(); len(uf) > 0 {
		ufc := make([]byte, len(uf))
		copy(ufc, uf)
		r.ProtoReflect().SetUnknown(ufc)
	}
	return r
}

func CloneVT_CodeGeneratorResponse_File(m *pluginpb.CodeGeneratorResponse_File) *pluginpb.CodeGeneratorResponse_File {
	if m == nil {
		return (*pluginpb.CodeGeneratorResponse_File)(nil)
	}
	r := &pluginpb.CodeGeneratorResponse_File{
		GeneratedCodeInfo: descriptorpb1.CloneVT_GeneratedCodeInfo(m.GeneratedCodeInfo),
	}
	if rhs := m.Name; rhs != nil {
		tmpVal := *rhs
		r.Name = &tmpVal
	}
	if rhs := m.InsertionPoint; rhs != nil {
		tmpVal := *rhs
		r.InsertionPoint = &tmpVal
	}
	if rhs := m.Content; rhs != nil {
		tmpVal := *rhs
		r.Content = &tmpVal
	}
	if uf := m.ProtoReflect().GetUnknown(); len(uf) > 0 {
		ufc := make([]byte, len(uf))
		copy(ufc, uf)
		r.ProtoReflect().SetUnknown(ufc)
	}
	return r
}

func CloneVT_CodeGeneratorResponse(m *pluginpb.CodeGeneratorResponse) *pluginpb.CodeGeneratorResponse {
	if m == nil {
		return (*pluginpb.CodeGeneratorResponse)(nil)
	}
	r := &pluginpb.CodeGeneratorResponse{}
	if rhs := m.Error; rhs != nil {
		tmpVal := *rhs
		r.Error = &tmpVal
	}
	if rhs := m.SupportedFeatures; rhs != nil {
		tmpVal := *rhs
		r.SupportedFeatures = &tmpVal
	}
	if rhs := m.File; rhs != nil {
		tmpContainer := make([]*pluginpb.CodeGeneratorResponse_File, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = CloneVT_CodeGeneratorResponse_File(v)
		}
		r.File = tmpContainer
	}
	if uf := m.ProtoReflect().GetUnknown(); len(uf) > 0 {
		ufc := make([]byte, len(uf))
		copy(ufc, uf)
		r.ProtoReflect().SetUnknown(ufc)
	}
	return r
}

func EqualVT_Version(this *pluginpb.Version, that *pluginpb.Version) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if p, q := this.Major, that.Major; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.Minor, that.Minor; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.Patch, that.Patch; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.Suffix, that.Suffix; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	return string(this.ProtoReflect().GetUnknown()) == string(that.ProtoReflect().GetUnknown())
}
func EqualVT_CodeGeneratorRequest(this *pluginpb.CodeGeneratorRequest, that *pluginpb.CodeGeneratorRequest) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if len(this.FileToGenerate) != len(that.FileToGenerate) {
		return false
	}
	for i, vx := range this.FileToGenerate {
		vy := that.FileToGenerate[i]
		if vx != vy {
			return false
		}
	}
	if p, q := this.Parameter, that.Parameter; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if !EqualVT_Version(this.CompilerVersion, that.CompilerVersion) {
		return false
	}
	if len(this.ProtoFile) != len(that.ProtoFile) {
		return false
	}
	for i, vx := range this.ProtoFile {
		vy := that.ProtoFile[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &descriptorpb.FileDescriptorProto{}
			}
			if q == nil {
				q = &descriptorpb.FileDescriptorProto{}
			}
			if !descriptorpb1.EqualVT_FileDescriptorProto(p, q) {
				return false
			}
		}
	}
	return string(this.ProtoReflect().GetUnknown()) == string(that.ProtoReflect().GetUnknown())
}
func EqualVT_CodeGeneratorResponse_File(this *pluginpb.CodeGeneratorResponse_File, that *pluginpb.CodeGeneratorResponse_File) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if p, q := this.Name, that.Name; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.InsertionPoint, that.InsertionPoint; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.Content, that.Content; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if !descriptorpb1.EqualVT_GeneratedCodeInfo(this.GeneratedCodeInfo, that.GeneratedCodeInfo) {
		return false
	}
	return string(this.ProtoReflect().GetUnknown()) == string(that.ProtoReflect().GetUnknown())
}
func EqualVT_CodeGeneratorResponse(this *pluginpb.CodeGeneratorResponse, that *pluginpb.CodeGeneratorResponse) bool {
	if this == nil {
		return that == nil
	} else if that == nil {
		return false
	}
	if p, q := this.Error, that.Error; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.SupportedFeatures, that.SupportedFeatures; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if len(this.File) != len(that.File) {
		return false
	}
	for i, vx := range this.File {
		vy := that.File[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &pluginpb.CodeGeneratorResponse_File{}
			}
			if q == nil {
				q = &pluginpb.CodeGeneratorResponse_File{}
			}
			if !EqualVT_CodeGeneratorResponse_File(p, q) {
				return false
			}
		}
	}
	return string(this.ProtoReflect().GetUnknown()) == string(that.ProtoReflect().GetUnknown())
}
func MarshalVT_Version(m *pluginpb.Version) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := SizeVT_Version(m)
	dAtA = make([]byte, size)
	n, err := MarshalToSizedBufferVT_Version(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func MarshalToVT_Version(m *pluginpb.Version, dAtA []byte) (int, error) {
	size := SizeVT_Version(m)
	return MarshalToSizedBufferVT_Version(m, dAtA[:size])
}

func MarshalToSizedBufferVT_Version(m *pluginpb.Version, dAtA []byte) (int, error) {
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
	if m.Suffix != nil {
		i -= len(*m.Suffix)
		copy(dAtA[i:], *m.Suffix)
		i = encodeVarint(dAtA, i, uint64(len(*m.Suffix)))
		i--
		dAtA[i] = 0x22
	}
	if m.Patch != nil {
		i = encodeVarint(dAtA, i, uint64(*m.Patch))
		i--
		dAtA[i] = 0x18
	}
	if m.Minor != nil {
		i = encodeVarint(dAtA, i, uint64(*m.Minor))
		i--
		dAtA[i] = 0x10
	}
	if m.Major != nil {
		i = encodeVarint(dAtA, i, uint64(*m.Major))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func MarshalVT_CodeGeneratorRequest(m *pluginpb.CodeGeneratorRequest) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := SizeVT_CodeGeneratorRequest(m)
	dAtA = make([]byte, size)
	n, err := MarshalToSizedBufferVT_CodeGeneratorRequest(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func MarshalToVT_CodeGeneratorRequest(m *pluginpb.CodeGeneratorRequest, dAtA []byte) (int, error) {
	size := SizeVT_CodeGeneratorRequest(m)
	return MarshalToSizedBufferVT_CodeGeneratorRequest(m, dAtA[:size])
}

func MarshalToSizedBufferVT_CodeGeneratorRequest(m *pluginpb.CodeGeneratorRequest, dAtA []byte) (int, error) {
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
	if len(m.ProtoFile) > 0 {
		for iNdEx := len(m.ProtoFile) - 1; iNdEx >= 0; iNdEx-- {
			size, err := descriptorpb1.MarshalToSizedBufferVT_FileDescriptorProto(m.ProtoFile[iNdEx], dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x7a
		}
	}
	if m.CompilerVersion != nil {
		size, err := MarshalToSizedBufferVT_Version(m.CompilerVersion, dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if m.Parameter != nil {
		i -= len(*m.Parameter)
		copy(dAtA[i:], *m.Parameter)
		i = encodeVarint(dAtA, i, uint64(len(*m.Parameter)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FileToGenerate) > 0 {
		for iNdEx := len(m.FileToGenerate) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.FileToGenerate[iNdEx])
			copy(dAtA[i:], m.FileToGenerate[iNdEx])
			i = encodeVarint(dAtA, i, uint64(len(m.FileToGenerate[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func MarshalVT_CodeGeneratorResponse_File(m *pluginpb.CodeGeneratorResponse_File) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := SizeVT_CodeGeneratorResponse_File(m)
	dAtA = make([]byte, size)
	n, err := MarshalToSizedBufferVT_CodeGeneratorResponse_File(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func MarshalToVT_CodeGeneratorResponse_File(m *pluginpb.CodeGeneratorResponse_File, dAtA []byte) (int, error) {
	size := SizeVT_CodeGeneratorResponse_File(m)
	return MarshalToSizedBufferVT_CodeGeneratorResponse_File(m, dAtA[:size])
}

func MarshalToSizedBufferVT_CodeGeneratorResponse_File(m *pluginpb.CodeGeneratorResponse_File, dAtA []byte) (int, error) {
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
	if m.GeneratedCodeInfo != nil {
		size, err := descriptorpb1.MarshalToSizedBufferVT_GeneratedCodeInfo(m.GeneratedCodeInfo, dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if m.Content != nil {
		i -= len(*m.Content)
		copy(dAtA[i:], *m.Content)
		i = encodeVarint(dAtA, i, uint64(len(*m.Content)))
		i--
		dAtA[i] = 0x7a
	}
	if m.InsertionPoint != nil {
		i -= len(*m.InsertionPoint)
		copy(dAtA[i:], *m.InsertionPoint)
		i = encodeVarint(dAtA, i, uint64(len(*m.InsertionPoint)))
		i--
		dAtA[i] = 0x12
	}
	if m.Name != nil {
		i -= len(*m.Name)
		copy(dAtA[i:], *m.Name)
		i = encodeVarint(dAtA, i, uint64(len(*m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func MarshalVT_CodeGeneratorResponse(m *pluginpb.CodeGeneratorResponse) (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := SizeVT_CodeGeneratorResponse(m)
	dAtA = make([]byte, size)
	n, err := MarshalToSizedBufferVT_CodeGeneratorResponse(m, dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func MarshalToVT_CodeGeneratorResponse(m *pluginpb.CodeGeneratorResponse, dAtA []byte) (int, error) {
	size := SizeVT_CodeGeneratorResponse(m)
	return MarshalToSizedBufferVT_CodeGeneratorResponse(m, dAtA[:size])
}

func MarshalToSizedBufferVT_CodeGeneratorResponse(m *pluginpb.CodeGeneratorResponse, dAtA []byte) (int, error) {
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
	if len(m.File) > 0 {
		for iNdEx := len(m.File) - 1; iNdEx >= 0; iNdEx-- {
			size, err := MarshalToSizedBufferVT_CodeGeneratorResponse_File(m.File[iNdEx], dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x7a
		}
	}
	if m.SupportedFeatures != nil {
		i = encodeVarint(dAtA, i, uint64(*m.SupportedFeatures))
		i--
		dAtA[i] = 0x10
	}
	if m.Error != nil {
		i -= len(*m.Error)
		copy(dAtA[i:], *m.Error)
		i = encodeVarint(dAtA, i, uint64(len(*m.Error)))
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
func SizeVT_Version(m *pluginpb.Version) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Major != nil {
		n += 1 + sov(uint64(*m.Major))
	}
	if m.Minor != nil {
		n += 1 + sov(uint64(*m.Minor))
	}
	if m.Patch != nil {
		n += 1 + sov(uint64(*m.Patch))
	}
	if m.Suffix != nil {
		l = len(*m.Suffix)
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.ProtoReflect().GetUnknown())
	return n
}

func SizeVT_CodeGeneratorRequest(m *pluginpb.CodeGeneratorRequest) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FileToGenerate) > 0 {
		for _, s := range m.FileToGenerate {
			l = len(s)
			n += 1 + l + sov(uint64(l))
		}
	}
	if m.Parameter != nil {
		l = len(*m.Parameter)
		n += 1 + l + sov(uint64(l))
	}
	if m.CompilerVersion != nil {
		l = SizeVT_Version(m.CompilerVersion)
		n += 1 + l + sov(uint64(l))
	}
	if len(m.ProtoFile) > 0 {
		for _, e := range m.ProtoFile {
			l = descriptorpb1.SizeVT_FileDescriptorProto(e)
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.ProtoReflect().GetUnknown())
	return n
}

func SizeVT_CodeGeneratorResponse_File(m *pluginpb.CodeGeneratorResponse_File) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Name != nil {
		l = len(*m.Name)
		n += 1 + l + sov(uint64(l))
	}
	if m.InsertionPoint != nil {
		l = len(*m.InsertionPoint)
		n += 1 + l + sov(uint64(l))
	}
	if m.Content != nil {
		l = len(*m.Content)
		n += 1 + l + sov(uint64(l))
	}
	if m.GeneratedCodeInfo != nil {
		l = descriptorpb1.SizeVT_GeneratedCodeInfo(m.GeneratedCodeInfo)
		n += 2 + l + sov(uint64(l))
	}
	n += len(m.ProtoReflect().GetUnknown())
	return n
}

func SizeVT_CodeGeneratorResponse(m *pluginpb.CodeGeneratorResponse) (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Error != nil {
		l = len(*m.Error)
		n += 1 + l + sov(uint64(l))
	}
	if m.SupportedFeatures != nil {
		n += 1 + sov(uint64(*m.SupportedFeatures))
	}
	if len(m.File) > 0 {
		for _, e := range m.File {
			l = SizeVT_CodeGeneratorResponse_File(e)
			n += 1 + l + sov(uint64(l))
		}
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
func UnmarshalVT_Version(m *pluginpb.Version, dAtA []byte) error {
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
			return fmt.Errorf("proto: Version: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Version: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Major", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Major = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minor", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Minor = &v
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Patch", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Patch = &v
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Suffix", wireType)
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
			s := string(dAtA[iNdEx:postIndex])
			m.Suffix = &s
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
func UnmarshalVT_CodeGeneratorRequest(m *pluginpb.CodeGeneratorRequest, dAtA []byte) error {
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
			return fmt.Errorf("proto: CodeGeneratorRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CodeGeneratorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileToGenerate", wireType)
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
			m.FileToGenerate = append(m.FileToGenerate, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parameter", wireType)
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
			s := string(dAtA[iNdEx:postIndex])
			m.Parameter = &s
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompilerVersion", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CompilerVersion == nil {
				m.CompilerVersion = &pluginpb.Version{}
			}
			if err := UnmarshalVT_Version(m.CompilerVersion, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtoFile", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProtoFile = append(m.ProtoFile, &descriptorpb.FileDescriptorProto{})
			if err := descriptorpb1.UnmarshalVT_FileDescriptorProto(m.ProtoFile[len(m.ProtoFile)-1], dAtA[iNdEx:postIndex]); err != nil {
				return err
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
func UnmarshalVT_CodeGeneratorResponse_File(m *pluginpb.CodeGeneratorResponse_File, dAtA []byte) error {
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
			return fmt.Errorf("proto: CodeGeneratorResponse_File: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CodeGeneratorResponse_File: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			s := string(dAtA[iNdEx:postIndex])
			m.Name = &s
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InsertionPoint", wireType)
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
			s := string(dAtA[iNdEx:postIndex])
			m.InsertionPoint = &s
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
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
			s := string(dAtA[iNdEx:postIndex])
			m.Content = &s
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GeneratedCodeInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GeneratedCodeInfo == nil {
				m.GeneratedCodeInfo = &descriptorpb.GeneratedCodeInfo{}
			}
			if err := descriptorpb1.UnmarshalVT_GeneratedCodeInfo(m.GeneratedCodeInfo, dAtA[iNdEx:postIndex]); err != nil {
				return err
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
func UnmarshalVT_CodeGeneratorResponse(m *pluginpb.CodeGeneratorResponse, dAtA []byte) error {
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
			return fmt.Errorf("proto: CodeGeneratorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CodeGeneratorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
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
			s := string(dAtA[iNdEx:postIndex])
			m.Error = &s
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupportedFeatures", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SupportedFeatures = &v
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field File", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.File = append(m.File, &pluginpb.CodeGeneratorResponse_File{})
			if err := UnmarshalVT_CodeGeneratorResponse_File(m.File[len(m.File)-1], dAtA[iNdEx:postIndex]); err != nil {
				return err
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