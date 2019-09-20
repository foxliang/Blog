package main

import (
	"fmt"
        "github.com/mikemintang/go-curl"
        "time"
	"regexp"
        "strconv"
	"database/sql" 
	_ "github.com/go-sql-driver/mysql"
)


func main() {
    ti := time.Now()
    fmt.Println("爬取开始 : ", ti.Format("2006-01-02 15:04:05"))
    http() 
    elapsed := time.Since(ti)
    fmt.Println("爬取结束，总共耗时: ", elapsed)
}

func http() {
    db, err := sql.Open("mysql", "root:@/go?charset=utf8mb4")
    if(err != nil ){
	fmt.Println("连接数据库失败 : ", err)
    }	
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

	    //插入数据
	    stmt, err := db.Prepare("INSERT INTO douban SET name=?,average=?,commentCount=?,careted_at=?")
	    if(err!=nil ){
		fmt.Println("插入数据失败 : ", err)
	    }
	    for i := 0; i < len(txt2); i++ {
		// fmt.Printf("%s %s %s\n", txt4[i][2], txt3[i][1], txt2[i][1])
		res, err := stmt.Exec(txt4[i][2], txt3[i][1], txt2[i][1], time.Now().Format("2006-01-02 15:04:05"))
		if(err!=nil ){
			fmt.Println("插入数据失败2 : ", err)
		}	
		id, err := res.LastInsertId()
		if(err!=nil ){
			fmt.Println("获取记录id失败 : ", err)
		}	
		fmt.Println("获取记录id : ", id)
            }
        } else {
            fmt.Println(resp.Raw)
        }
        
    }
}

