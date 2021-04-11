package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	readFile()
}

func readFile() {
	file, err := os.Open("../sea.png")
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
		readBuf := make([]byte, 1024 * 24)
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
			rate := int64(len(fileByte) * 100) / fileSize
			if rate > 100 {
				rate = 100
			}
			rateChan <- rate
		}()
	}
	ioutil.WriteFile("../sea-copy.png", fileByte, 0600)
}

func progress(ch <-chan int64) {
	for rate := range ch {
		fmt.Printf("\rrate:%3d%%", rate)
	}
}
