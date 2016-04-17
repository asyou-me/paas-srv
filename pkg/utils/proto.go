package utils

import (
	"encoding/binary"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/golang/snappy"
	"io"
	"net"
)

func RecvFrame(r io.Reader, data proto.Message) (err error) {
	result, err := recvFrame(r)
	if err != nil {
		return err
	}

	// decode the compressed data
	pb, err := snappy.Decode(nil, result)
	if err != nil {
		return err
	}

	if data != nil {
		err = proto.Unmarshal(pb, data)
		if err != nil {
			return err
		}
	}

	return nil
}

func SendFrame(w io.Writer, data proto.Message) (err error) {
	// marshal response
	pbData := []byte{}
	if data == nil {
		return errors.New("data is nil")
	}

	pbData, err = proto.Marshal(data)
	if err != nil {
		return err
	}

	compressedPb := snappy.Encode(nil, pbData)

	err = sendFrame(w, compressedPb)

	if err != nil {
		return err
	}

	if data == nil {
		return errors.New("data is nil")
	}

	err = proto.Unmarshal(pbData, data)

	return
}

func sendFrame(w io.Writer, data []byte) (err error) {
	// Allocate enough space for the biggest uvarint
	var size [binary.MaxVarintLen64]byte

	if data == nil || len(data) == 0 {
		n := binary.PutUvarint(size[:], uint64(0))
		if err = write(w, size[:n], false); err != nil {
			return
		}
		return
	}

	// Write the size and data
	n := binary.PutUvarint(size[:], uint64(len(data)))
	if err = write(w, size[:n], false); err != nil {
		return
	}
	if err = write(w, data, false); err != nil {
		return
	}
	return
}

func recvFrame(r io.Reader) (data []byte, err error) {
	size, err := readUvarint(r)
	if err != nil {
		return nil, err
	}
	if size != 0 {
		data = make([]byte, size)
		if err = read(r, data); err != nil {
			return nil, err
		}
	}
	return data, nil
}

// ReadUvarint reads an encoded unsigned integer from r and returns it as a uint64.
func readUvarint(r io.Reader) (uint64, error) {
	var x uint64
	var s uint
	for i := 0; ; i++ {
		var b byte
		b, err := readByte(r)
		if err != nil {
			return 0, err
		}
		if b < 0x80 {
			if i > 9 || i == 9 && b > 1 {
				return x, errors.New("protorpc: varint overflows a 64-bit integer")
			}
			return x | uint64(b)<<s, nil
		}
		x |= uint64(b&0x7f) << s
		s += 7
	}
}

func write(w io.Writer, data []byte, onePacket bool) error {
	if onePacket {
		if _, err := w.Write(data); err != nil {
			return err
		}
		return nil
	}
	for index := 0; index < len(data); {
		n, err := w.Write(data[index:])
		if err != nil {
			if nerr, ok := err.(net.Error); !ok || !nerr.Temporary() {
				return err
			}
		}
		index += n
	}
	return nil
}

func read(r io.Reader, data []byte) error {
	for index := 0; index < len(data); {
		n, err := r.Read(data[index:])
		if err != nil {
			if nerr, ok := err.(net.Error); !ok || !nerr.Temporary() {
				return err
			}
		}
		index += n
	}
	return nil
}

func readByte(r io.Reader) (c byte, err error) {
	data := make([]byte, 1)
	if err = read(r, data); err != nil {
		return 0, err
	}
	c = data[0]
	return
}
