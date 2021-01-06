package main

import (
	easyjson "easyjson/proto"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	//
	s := easyjson.Student{
		Id:   11,
		Name: "qq",
		School: easyjson.School{
			Name: "CUMT",
			Addr: "xz",
		},
		Birthday: time.Now(),
	}
	bt, err := s.MarshalJSON()
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("bt", string(bt))
	//json := `{"id":11,"name":"s_name","s_chool":{"name":"CUMT","addr":"xz"},"birthday":"2017-08-04T20:58:07.9894603+08:00"}`
	ss := easyjson.Student{}
	ss.UnmarshalJSON(bt)
	fmt.Println("ss", ss)

	end := time.Since(start)
	fmt.Println("Since", end)
}


