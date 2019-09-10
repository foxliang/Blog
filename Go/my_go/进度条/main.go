package main

import (
	"fmt"
        "time"
	"os"
	"io"
	"bufio"
	"path/filepath"
	"unsafe"
)

func main() {
	ti := time.Now()
	fmt.Println("开始 : ", ti.Format("2006-01-02 15:04:05"))
	read("test.txt",1024)
	elapsed := time.Since(ti)
	fmt.Println("\n结束，总共耗时: ", elapsed)
}

func read (filePth string, bufSize int){
	total_size := getFileSize(filePth)
        // fmt.Println("文件总大小: ", total_size) // 总字节数
	file,err := os.Open(filePth)
	if err != nil{
		fmt.Println(err)
		return 
	}
	defer file.Close()
	//按照字节数读取
	buf := make([] byte,bufSize)
	bfRd := bufio.NewReader(file)
	var int_size int = *(*int)(unsafe.Pointer(&total_size))
	var one_size int 
	fmt.Printf("%d%%", 0)     // 输出一行结果
	for {
		n, err := bfRd.Read(buf)  
		if err != nil { 
			if err == io.EOF {
				return 
		    }
			fmt.Println(err)
			return
		} 
		one_size += n
		// fmt.Println("单次大小: ", one_size)
		per := (one_size * 100 / int_size)
		// fmt.Println("百分比: ", per)
		fmt.Printf("\r%d%%", per)     // 输出第二行结果
		time.Sleep(time.Duration(100) * time.Millisecond)  //延迟输出
	}
}

func getFileSize(filename string) int64 {
    var result int64
    filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
        result = f.Size()
        return nil
    })
    return result
}

