package tail

import (
	"bufio"
	"github.com/spf13/cast"
	"gopkg.in/fsnotify.v1"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hpcloud/tail"
)

/**
github.com/hpcloud/tail
*/
func Tailf2() {
	file, err := os.OpenFile("../../tailf.log", os.O_CREATE|os.O_RDWR|os.O_SYNC|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	log.Println("tailf start...")

	go func() {
		WriteString(file)
	}()

	t, err := tail.TailFile("../../tailf.log", tail.Config{Follow: true})
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		log.Println(line.Text)
	}
}

/**
最简单的实现，循环读取文件，没有新内容就休眠一段时间
但此种方法无效，新写入的内容并不能读取到
*/
func Tailf() {
	file, err := os.OpenFile("../file/tailf.log", os.O_CREATE|os.O_RDWR|os.O_SYNC|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	log.Println("tailf start...")

	go func() {
		WriteString(file)
	}()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Println("read error:", err)
			break
		}

		if err == io.EOF {
			time.Sleep(time.Second)
		} else {
			log.Print(line)
		}
	}
}

/**
linux inotify
*/
func TailWatcher() {
	file, err := os.OpenFile("../file/tailf.log", os.O_CREATE|os.O_RDWR|os.O_SYNC|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	log.Println("tailf start...")

	go func() {
		WriteString(file)
	}()

	offset := SeekTo(file, 0, io.SeekEnd)
	reader := bufio.NewReader(file)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()
	err = watcher.Add("../file/tailf.log")
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println(event)
				//file, _ := os.OpenFile("../file/tailf.log", os.O_CREATE|os.O_RDWR|os.O_SYNC|os.O_APPEND, os.ModePerm)
				offset = SeekTo(file, offset, io.SeekCurrent)
				line, _ := reader.ReadString('\n')
				log.Println(offset)
				log.Println(line)

			case err := <-watcher.Errors:
				log.Println("watcher error:", err)
			}
		}
	}()

	done := make(chan struct{}, 1)
	<-done
}

func readLines(reader *bufio.Reader) {
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Println("read error:", err)
			break
		}

		if err == io.EOF {
			time.Sleep(time.Second)
		} else {
			log.Print(line)
		}
	}
}

// Seek设置下一次读/写的位置。offset为相对偏移量，
// 而whence决定相对位置：0为相对文件开头，1为相对当前位置，2为相对文件结尾。它返回新的偏移量（相对开头）和可能的错误

func WriteString(file *os.File) {
	for {
		rand.Seed(time.Now().UnixNano())
		if _, err := io.WriteString(file, "str "+cast.ToString(rand.Int31n(10000))+"\n"); err != nil {
			log.Println("write error:", err)
		}
		time.Sleep(2000 * time.Millisecond)
	}

}
