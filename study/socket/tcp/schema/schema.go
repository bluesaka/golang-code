/**
定义一个协议，比如数据包的前4个字节为包头，里面存储的是发送的数据的长度。
解决TCP粘包问题
*/

package schema

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// Encode 将消息编码
func Encode(message string) ([]byte, error) {
	// 读取消息的长度，转换成int32类型（占4个字节）
	length := int32(len(message))
	buf := new(bytes.Buffer)
	// 写入消息头
	if err := binary.Write(buf, binary.LittleEndian, length); err != nil {
		return nil, err
	}

	// 写入消息实体
	if err := binary.Write(buf, binary.LittleEndian, []byte(message)); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error) {
	// 读取前4个字节的数据
	peekByte, _ := reader.Peek(4)
	peekBuf := bytes.NewBuffer(peekByte)
	var length int32
	if err := binary.Read(peekBuf, binary.LittleEndian, &length); err != nil {
		return "", err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered()) < length+4 {
		return "", nil
	}
	// 读取真正的消息数据
	pack := make([]byte, int(4+length))
	if _, err := reader.Read(pack); err != nil {
		return "", err
	}

	return string(pack[4:]), nil
}
