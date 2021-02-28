package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

func WriteProtoBufTobiaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to Binary: %w", err)
	}
	err = ioutil.WriteFile(filename, data, 0644)

	if err != nil {
		return fmt.Errorf("cannot marshal proto message to Binary: %w", err)
	}
	return nil
}

func ReadProtoBufFrombinaryFile(filename string, message proto.Message) error {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary data from file: %w", err)
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal binary file to protomessage: %w", err)
	}
	return nil
}

func WriteProtoBufToJson(message proto.Message, filename string) error {
	data, err := ProtoBufToJson(message)
	if err != nil {
		return fmt.Errorf("cannot marshal binary file to proto message: %w", err)
	}
	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write  binary data to file: %w", err)
	}
	return nil
}
