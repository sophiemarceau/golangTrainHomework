package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"github.com/pkg/errors"
)

/**
2.实现一个从 socket connection 中解码出 goim 协议的解码器。
go im
4 bytes packetLen 包长度，在数据流传输过程中，先写入整个包的长度，方便整个包的数据读取
2 bytes HeaderLen 头长度，在处理数据时，会先解析头部，可以知道具体业务操作
2 bytes Version 协议版本号，朱勇用于上行和下行数据包按版本号进行解析
4 bytes Operation 业务操作码，可以按操作码进行分发数据包到具体业务当中
4 bytes Sequence 序列号 数据包的唯一标记，可以做具体业务处理，或者数据表去重
PacketLen-HeaderLen Body 实际业务数据，在业务层中会进行数据解码和编码
*/
func main() {
	data := encoder("Hello, goLang!")
	decoder(data)
}

func decoder(data []byte) {
	if len(data) <= 16 {
		fmt.Println("data len < 16.")
		return
	}
	packetLen := binary.BigEndian.Uint32(data[:4])
	fmt.Printf("packetLen:%v\n", packetLen)
	headerLen := binary.BigEndian.Uint16(data[4:6])
	fmt.Printf("headerLen:%v\n", headerLen)
	version := binary.BigEndian.Uint32(data[6:8])
	fmt.Printf("version:%v\n", version)
	operation := binary.BigEndian.Uint32(data[8:12])
	fmt.Printf("operation:%v\n", operation)
	sequence := binary.BigEndian.Uint32(data[12:16])
	fmt.Printf("sequence:%v\n", sequence)
	body := string(data[16:])
	fmt.Printf("body:%v\n", body)
}

func encoder(body string) []byte {
	headerLen := 16
	packetLen := len(body) + headerLen
	ret := make([]byte, packetLen)
	binary.BigEndian.PutUint32(ret[:4], uint32(packetLen))
	binary.BigEndian.PutUint16(ret[4:6], uint16(headerLen))
	version := 5
	binary.BigEndian.PutUint16(ret[6:8], uint16(version))
	operation := 6
	binary.BigEndian.PutUint32(ret[8:12], uint32(operation))
	sequence := 7
	binary.BigEndian.PutUint32(ret[12:16], uint32(sequence))
	byteBody := []byte(body)
	copy(ret[16:], byteBody)
	return ret
}

const MaxSize = int32(1 << 12)

const (
	PackageLengthSize = 4
	HeaderLengthSize  = 2
	VersionSize       = 4
	OperationSize     = 2
	SequenceIdSize    = 4
)

type Message struct {
	HeaderLen  int16
	Version    int16
	Operation  int32
	SequenceId int32
	Body       string
}

func Pack(ctx context.Context, m *Message) []byte {
	packLen := HeaderLengthSize + len(m.Body)
	buf := make([]byte, 1000)
	binary.BigEndian.PutUint32(buf, uint32(packLen))
	binary.BigEndian.PutUint16(buf, uint16(m.HeaderLen))
	binary.BigEndian.PutUint16(buf, uint16(m.Version))
	binary.BigEndian.PutUint32(buf, uint32(m.Operation))
	binary.BigEndian.PutUint32(buf, uint32(m.SequenceId))
	if len(m.Body) > 0 {
		buf = append(buf, []byte(m.Body)...)
	}
	return buf
}

func ReadSocket(ctx context.Context, buffer *bytes.Buffer) (*Message, error) {
	lengthSize := PackageLengthSize + HeaderLengthSize + VersionSize + OperationSize + SequenceIdSize
	buf := buffer.Next(lengthSize)
	if len(buf) != lengthSize*8 {
		return nil, errors.New("invalid message")
	}
	message := &Message{
		0,
		0,
		0,
		0,
		"",
	}
	packLenOffset := PackageLengthSize
	packLen := binary.BigEndian.Uint32(buf[0:PackageLengthSize])
	if packLen > uint32(MaxSize) {
		return nil, errors.New("max size over limit")
	}
	headerLenOffset := packLenOffset + HeaderLengthSize
	message.HeaderLen = int16(binary.BigEndian.Uint16(buf[packLenOffset:headerLenOffset]))
	if message.HeaderLen != HeaderLengthSize {
		return nil, errors.New("invalid header length")
	}
	versionOffset := headerLenOffset + VersionSize
	message.Version = int16(binary.BigEndian.Uint16(buf[headerLenOffset:versionOffset]))
	OperationOffset := versionOffset + OperationSize
	message.Operation = int32(binary.BigEndian.Uint32(buf[versionOffset:OperationOffset]))
	sequenceOffset := OperationOffset + SequenceIdSize
	message.SequenceId = int32(binary.BigEndian.Uint32(buf[OperationOffset:sequenceOffset]))

	bodyLen := int32(packLen) - int32(message.HeaderLen)
	if bodyLen > 0 {
		message.Body = string(buffer.Next(int(bodyLen)))
	}
	return message, nil
}
