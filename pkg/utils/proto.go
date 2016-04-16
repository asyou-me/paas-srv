package utils

import (
	"github.com/golang/protobuf/proto"
	"io"
)

func RecvFrame(r io.Reader, data proto.Message) (err error) {
	return nil
}

func SendFrame(w io.Writer, data proto.Message) (err error) {
	// marshal response
	pbData := []byte{}
	if data != nil {
		pbData, err = proto.Marshal(data)
		if err != nil {
			return err
		}
	}

	_, err = w.Write(pbData)
	return
}
