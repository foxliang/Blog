package main

import (
    "fmt"
    "os"
    "github.com/mikemintang/go-curl"
    "time"
    "regexp"
    "strconv"
    "encoding/csv"
)

func http() {
    //创建csv文件
    f, err := os.Create("./douban.csv")
    if err != nil {
       panic(err)
    }
    defer f.Close()

    f.WriteString("\xEF\xBB\xBF")// 写入UTF-8 BOM
    //写入
    w := csv.NewWriter(f)//创建一个新的写入文件流
    data := [][]string{{"电影名称","评分","评价数量"}}
    w.WriteAll(data)//写入数据
    w.Flush()

    // 读取每页数据并写入
    for i := 0; i < 10; i++ {
    	url := "https://movie.douban.com/top250?start="+strconv.Itoa(i*25)
    	// 链式操作
    	req := curl.NewRequest()
    	resp, err := req.
    	SetUrl(url).
    	Get()
    	// 获取结果进行处理    
    	if err != nil {
	    fmt.Println(err)
    	}
    	if resp.IsOk() {
	    // fmt.Println(resp.Body)
	    html:=resp.Body
	    //评价人数
	    commentCount := `<span>(.*?)评价</span>`
	    rp2 := regexp.MustCompile(commentCount)
	    txt2 := rp2.FindAllStringSubmatch(html, -1)

	    //评分
	    pattern3 := `property="v:average">(.*?)</span>`
	    rp3 := regexp.MustCompile(pattern3)
	    txt3 := rp3.FindAllStringSubmatch(html, -1)

	    //电影名称
	    pattern4 := `img width="(.*?)" alt="(.*?)" src=`
	    rp4 := regexp.MustCompile(pattern4)
	    txt4 := rp4.FindAllStringSubmatch(html, -1)

	    w := csv.NewWriter(f)//创建一个新的写入文件流
	    data := [][]string{}
	    for i := 0; i < len(txt2); i++ {
	        // fmt.Printf("%s %s %s\n", txt4[i][2], txt3[i][1], txt2[i][1])
	        data = append(data, []string{txt4[i][2],txt3[i][1],txt2[i][1]})
	    }
	    w.WriteAll(data)//写入数据
	    w.Flush()
	} else {
	    fmt.Println(resp.Raw)
	}
    }
}


func main() {
    ti := time.Now()
    fmt.Println("爬取开始 : ", ti.Format("2006-01-02 15:04:05"))
    http() 
    elapsed := time.Since(ti)
    fmt.Println("爬取结束，总共耗时: ", elapsed)
}
