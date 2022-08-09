package conformance

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"testing"
)

var (
	testMsg = &TestAllTypesProto3{
		//OneofField: &TestAllTypesProto3_OneofNullValue{OneofNullValue: structpb.NullValue_NULL_VALUE},
		//
		//OptionalBoolWrapper:   wrapperspb.Bool(true),
		//OptionalInt32Wrapper:  wrapperspb.Int32(1),
		//OptionalInt64Wrapper:  wrapperspb.Int64(1),
		//OptionalUint32Wrapper: wrapperspb.UInt32(1),
		//OptionalUint64Wrapper: wrapperspb.UInt64(1),
		//OptionalFloatWrapper:  wrapperspb.Float(1),
		//OptionalDoubleWrapper: wrapperspb.Double(1),
		//OptionalStringWrapper: wrapperspb.String("blip"),
		//OptionalBytesWrapper:  wrapperspb.Bytes([]byte("blop")),
		//
		//RepeatedBoolWrapper:   []*wrapperspb.BoolValue{wrapperspb.Bool(true)},
		//RepeatedInt32Wrapper:  []*wrapperspb.Int32Value{wrapperspb.Int32(1)},
		//RepeatedInt64Wrapper:  []*wrapperspb.Int64Value{wrapperspb.Int64(1)},
		//RepeatedUint32Wrapper: []*wrapperspb.UInt32Value{wrapperspb.UInt32(1)},
		//RepeatedUint64Wrapper: []*wrapperspb.UInt64Value{wrapperspb.UInt64(1)},
		//RepeatedFloatWrapper:  []*wrapperspb.FloatValue{wrapperspb.Float(1)},
		//RepeatedDoubleWrapper: []*wrapperspb.DoubleValue{wrapperspb.Double(1)},
		//RepeatedStringWrapper: []*wrapperspb.StringValue{wrapperspb.String("blip")},
		//RepeatedBytesWrapper:  []*wrapperspb.BytesValue{wrapperspb.Bytes([]byte("blop"))},
		//
		//// OptionalDuration:      *durationpb.Duration
		//// OptionalTimestamp:     *timestamppb.Timestamp
		//// OptionalFieldMask:     *fieldmaskpb.FieldMask
		//// OptionalStruct:        *structpb.Struct
		//// OptionalAny:           *anypb.Any
		//OptionalValue: structpb.NewNumberValue(42),
		//// OptionalNullValue:     structpb.NullValue
		//
		//// repeated google.protobuf.Duration repeated_duration
		//// repeated google.protobuf.Timestamp repeated_timestamp
		//// repeated google.protobuf.FieldMask repeated_fieldmask
		//// repeated google.protobuf.Struct repeated_struct
		//// repeated google.protobuf.Any repeated_any
		//// repeated google.protobuf.Value repeated_value
		//RepeatedValue: []*structpb.Value{structpb.NewNumberValue(42)},
		// repeated google.protobuf.ListValue repeated_list_value
	}
)

func init() {
	fmt.Println("Size before:", testMsg.SizeVT())
	return
	PopulateStruct(testMsg)
	testMsg.OptionalBoolWrapper = nil
	testMsg.OptionalInt32Wrapper = nil
	testMsg.OptionalInt64Wrapper = nil
	testMsg.OptionalUint32Wrapper = nil
	testMsg.OptionalUint64Wrapper = nil
	testMsg.OptionalFloatWrapper = nil
	testMsg.OptionalDoubleWrapper = nil
	testMsg.OptionalStringWrapper = nil
	testMsg.OptionalBytesWrapper = nil
	testMsg.RepeatedBoolWrapper = nil
	testMsg.RepeatedInt32Wrapper = nil
	testMsg.RepeatedInt64Wrapper = nil
	testMsg.RepeatedUint32Wrapper = nil
	testMsg.RepeatedUint64Wrapper = nil
	testMsg.RepeatedFloatWrapper = nil
	testMsg.RepeatedDoubleWrapper = nil
	testMsg.RepeatedStringWrapper = nil
	testMsg.RepeatedBytesWrapper = nil
	testMsg.OptionalDuration = nil
	testMsg.OptionalTimestamp = nil
	testMsg.OptionalFieldMask = nil
	testMsg.OptionalStruct = nil
	testMsg.OptionalAny = nil
	testMsg.OptionalValue = nil
	testMsg.RepeatedDuration = nil
	testMsg.RepeatedTimestamp = nil
	testMsg.RepeatedFieldmask = nil
	testMsg.RepeatedStruct = nil
	testMsg.RepeatedAny = nil
	testMsg.RepeatedValue = nil
	testMsg.RepeatedListValue = nil
	testMsg.RecursiveMessage = nil
	fmt.Println("Size after:", testMsg.SizeVT())
}

func BenchmarkCloneVT(b *testing.B) {
	for i := 0; i < b.N*10; i++ {
		cloned := testMsg.CloneVT()
		b.StopTimer()
		assert.True(b, cloned.EqualVT(testMsg))
		b.StartTimer()
	}
}

func BenchmarkProtoClone(b *testing.B) {
	for i := 0; i < b.N*10; i++ {
		cloned := proto.Clone(testMsg).(*TestAllTypesProto3)
		b.StopTimer()
		assert.True(b, cloned.EqualVT(testMsg))
		b.StartTimer()
	}
}

//
//func BenchmarkMarshalVTUnmarshalVT(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		buf, err := testMsg.MarshalVT()
//		require.NoError(b, err)
//		var cloned TestAllTypesProto3
//		err = cloned.UnmarshalVT(buf)
//		require.NoError(b, err)
//	}
//}
//
//func BenchmarkProtoMarshalProtoUnmarshal(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		buf, err := proto.Marshal(testMsg)
//		require.NoError(b, err)
//		var cloned TestAllTypesProto3
//		err = proto.Unmarshal(buf, &cloned)
//		require.NoError(b, err)
//	}
//}
