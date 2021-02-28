package serializer_test

import (
	"testing"

	"github.com/MohamedNazir/TheCompleteGRPC/pb/github.com/MohamedNazir/TheCompleteGRPC/proto/pb"
	"github.com/MohamedNazir/TheCompleteGRPC/sample"
	"github.com/MohamedNazir/TheCompleteGRPC/serializer"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryfile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()

	err := serializer.WriteProtoBufTobiaryFile(laptop1, binaryfile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtoBufFrombinaryFile(binaryfile, laptop2)
	require.NoError(t, err)

	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtoBufToJson(laptop1, jsonFile)
	require.NoError(t, err)

}
