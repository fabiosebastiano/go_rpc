package serializer

import (
	"fmt"
	"testing"

	"github.com/fabiosebastiano/go_grpc/pb"
	"github.com/fabiosebastiano/go_grpc/sample"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

// WriteProtobufToBinaryFile
func TestFileSerializer(t *testing.T) {

	t.Parallel()
	fmt.Println("EHI")
	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()

	err := WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}

	err = ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}
