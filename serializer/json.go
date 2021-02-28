package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// ProtoBufToJson converts protocol buffer message to JSON string
func ProtoBufToJson(message proto.Message) (string, error) {
	marshaller := jsonpb.Marshaler{
		EnumsAsInts:  false,
		EmitDefaults: true,
		OrigName:     false,
		Indent: "	",
	}
	return marshaller.MarshalToString(message)
}

// JSONToProtobufMessage converts JSON string to protocol buffer message
func JSONToProtobufMessage(data string, message proto.Message) error {
	return jsonpb.UnmarshalString(data, message)
}
