package main

import (
	"fmt"
        "time"
	"os"
	"io"
	"bufio"
	"path/filepath"
	"math"
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
	var one_size int 
	sum := math.Ceil(float64(total_size)/float64(bufSize))
	i := 0
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
		i++
		// fmt.Println("单次大小: ", one_size)
		per := (one_size * 100 / int(total_size))
		// fmt.Printf("%d%%\r", per)     //只输出百分比100%
		fmt.Printf("%d%% [%s]\r",per,getS(i,"#") + getS(int(sum)-i," "))  //100% [########]		
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

func getS(n int,char string) (s string) {
    for i:=1;i<=n;i++{
        s+=char
    }
    return
}
