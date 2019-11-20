package main

import (
	"fmt"
	"strings"
	"time"
)

type LogProcess struct {
	rc chan string
	wc chan  string
	read Reader
	write Writer

	
}
type Reader interface {
	Read(rc chan string)
}

type Writer interface {
	Write(wc chan string)
}

type ReaderFromFile struct {
	path string //读取文件的路径
}
type WriteToInfluxDB struct {
	influxDBDsn string // influx data source
}

func (r *ReaderFromFile) Read(rc chan string)  {
	//读取模块
	line := "message"
	rc <- line
}

func (w *WriteToInfluxDB) Write(wc chan string)  {
	// 写入模块
	fmt.Print(<-wc)

}


func (l *LogProcess) Process()  {
	//解析模块
	data := <- l.rc
	l.wc <- strings.ToUpper(data)
}


func main()  {

	r := &ReaderFromFile{path:        "/tmp/access.log"}
	w := &WriteToInfluxDB{influxDBDsn:"username&password"}
	lp := &LogProcess{
		rc:make(chan string),
		wc: make(chan string),
		read:r,
		write:w,
	}

	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)

	time.Sleep(1*time.Second)
}
