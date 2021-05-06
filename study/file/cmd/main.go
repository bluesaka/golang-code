package main

import (
	"fmt"
	"go-code/study/file/tail"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	tail.TailWatcher()
}

/**
flag int 可用的打开方式有
// Flags to OpenFile wrapping those of the underlying system. Not all
// flags may be implemented on a given system.
const (
    // Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
    // 只读模式
    O_RDONLY int = syscall.O_RDONLY // open the file read-only.
    // 只写模式
    O_WRONLY int = syscall.O_WRONLY // open the file write-only.
    // 可读可写
    O_RDWR   int = syscall.O_RDWR   // open the file read-write.
    // The remaining values may be or'ed in to control behavior.
    // 追加内容
    O_APPEND int = syscall.O_APPEND // append data to the file when writing.
    // 创建文件，如果文件不存在
    O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
    // 与创建文件一同使用，文件必须存在
    O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
    // 打开一个同步的文件流
    O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
    // 如果可能，打开时缩短文件
    O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
)

perm fileMode 打开模式
// The defined file mode bits are the most significant bits of the FileMode.
// The nine least-significant bits are the standard Unix rwxrwxrwx permissions.
// The values of these bits should be considered part of the public API and
// may be used in wire protocols or disk representations: they must not be
// changed, although new bits might be added.
const (
    // The single letters are the abbreviations
    // used by the String method's formatting.
    // 文件夹模式
    ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
    // 追加模式
    ModeAppend                                     // a: append-only
    // 单独使用
    ModeExclusive                                  // l: exclusive use
    // 临时文件
    ModeTemporary                                  // T: temporary file; Plan 9 only
    // 象征性的关联
    ModeSymlink                                    // L: symbolic link
    // 设备文件
    ModeDevice                                     // D: device file
    // 命名管道
    ModeNamedPipe                                  // p: named pipe (FIFO)
    // Unix 主机 socket
    ModeSocket                                     // S: Unix domain socket
    // 设置uid
    ModeSetuid                                     // u: setuid
    // 设置gid
    ModeSetgid                                     // g: setgid
    // UNIX 字符串设备，当设备模式是设置unix
    ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
    // 粘性的
    ModeSticky                                     // t: sticky
    // 非常规文件；对该文件一无所知
    ModeIrregular                                  // ?: non-regular file; nothing else is known about this file

    // bit位遮盖，不变的文件设置为none
    // Mask for the type bits. For regular files, none will be set.
    ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeCharDevice | ModeIrregular
    // 权限位
    ModePerm FileMode = 0777 // Unix permission bits
)

*/
func read1() {
	flag := os.O_CREATE | os.O_RDWR | os.O_APPEND
	perm := os.ModeAppend | os.ModePerm
	file, err := os.OpenFile("../file/test.log", flag, perm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("test\n")
}

func readFile() {
	file, err := os.Open("../file/sea.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	rateChan := make(chan int64)
	defer close(rateChan)
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	fileByte := make([]byte, 0)

	go progress(rateChan)

	for {
		readBuf := make([]byte, 1024*24)
		n, err := file.Read(readBuf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		fileByte = append(fileByte, readBuf...)
		time.Sleep(time.Second)
		go func() {
			rate := int64(len(fileByte)*100) / fileSize
			if rate > 100 {
				rate = 100
			}
			rateChan <- rate
		}()
	}
	ioutil.WriteFile("../file/sea-copy.png", fileByte, 0600)
}

func progress(ch <-chan int64) {
	for rate := range ch {
		fmt.Printf("\rrate:%3d%%", rate)
	}
}
