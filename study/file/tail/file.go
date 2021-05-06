package tail

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

/**
Seek设置下一次读/写的位置
offset 偏移量
whence，从哪开始：0从头，1当前，2末尾
*/
func Seek() {
	file, err := os.OpenFile("../file/seek.log", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	//SeekReader(file, reader, io.SeekStart)
	SeekReader(file, reader, io.SeekEnd)
	SeekReader(file, reader, io.SeekCurrent)
}

func SeekReader(file *os.File, reader *bufio.Reader, whence int) {
	offset, err := file.Seek(0, whence)
	//offset, err = file.Seek(offset, whence)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("offset is:", offset)
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Println("read error:", err)
	}

	if err == nil {
		log.Println(strings.TrimRight(line, "\n"))
	}
}

func SeekTo(file *os.File, offset int64, whence int) int64 {
	offset, err := file.Seek(offset, whence)
	if err != nil {
		log.Fatal(err)
	}
	return offset
}
