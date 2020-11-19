package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"time"
)

type list struct {
	Name  string    `json:"name"`
	Title string    `json:"title"`
	Type  string    `json:"type"`
	Url   string    `json:"url"`
	Time  int       `json:"time"`
	Date  time.Time `json:"date"`
}

type lists []list

type SearchResponse struct {
	Took     uint64 `json:"took"`
	TimedOut bool   `json:"timed_out"`
	Shards   struct {
		Total      uint64 `json:"total"`
		Successful uint64 `json:"successful"`
		Skipped    uint64 `json:"skipped"`
		Failed     uint64 `json:"failed"`
	} `json:"_shards"`
	Hits         *SearchResponseHits        `json:"hits"`
	Aggregations map[string]json.RawMessage `json:"aggregations"`
}

// SearchResponseHits search response hits
type SearchResponseHits struct {
	Total struct {
		Value    uint64 `json:"value"`
		Relation string `json:"relation"`
	} `json:"total"`
	MaxScore float64                   `json:"max_score"`
	Hits     []*SearchResponseHitsHits `json:"hits"`
}

// SearchResponseHitsHits search response hits hits
type SearchResponseHitsHits struct {
	Index  string          `json:"_index"`
	Type   string          `json:"_type"`
	ID     string          `json:"_id"`
	Score  float64         `json:"_score"`
	Source json.RawMessage `json:"_source"`
}

// RangeIDSetter range set id
type RangeIDSetter interface {
	SetID(int, string)
}

type resData struct {
	List  []list `json:"list"`
	Total uint64 `json:"total"`
}

func main() {
	start := time.Now() // 获取当前时间
	addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
	config := elasticsearch.Config{
		Addresses: addresses,
		Username:  "",
		Password:  "",
		CloudID:   "",
		APIKey:    "",
	}
	// new client
	es, err := elasticsearch.NewClient(config)
	if err != nil {
		fmt.Println(err, "Error creating the client")
	}
	res := search(*es, "top1", "博客园")
	fmt.Println("res", res)
	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}

func search(client elasticsearch.Client, index string, query string) resData {
	var buf bytes.Buffer
	queryData := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"name": query,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(queryData); err != nil {
		fmt.Println(err, "Error encoding query")
	}

	resp, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(index),
		client.Search.WithBody(&buf),
		client.Search.WithSort("time:desc"),
		client.Search.WithFrom(10),
		client.Search.WithSize(10),
	)
	if err != nil {
		fmt.Println("err1", err)
	}
	//fmt.Println("resp", resp)
	//fmt.Println("Body", resp.Body)
	//fmt.Println("StatusCode", resp.StatusCode)

	var lists lists
	total, err := DecodeSearch(resp, &lists)
	if err != nil {
		fmt.Println("err2", err)
	}
	return resData{
		List:  lists,
		Total: total,
	}

}

// DecodeSearch decode search response
func DecodeSearch(resp *esapi.Response, v interface{}) (uint64, error) {
	if resp.StatusCode == 404 {
		return 0, nil
	}

	if resp.StatusCode != 200 {
		return 0, errors.New("1")
	}

	var r SearchResponse
	err := json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return 0, err
	}

	ids := make([]string, 0, len(r.Hits.Hits))
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, doc := range r.Hits.Hits {
		ids = append(ids, doc.ID)
		buf.Write(doc.Source)
		if i != len(r.Hits.Hits)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte(']')

	err = json.NewDecoder(&buf).Decode(v)
	if err != nil {
		return 0, err
	}

	if setter, ok := v.(RangeIDSetter); ok {
		for i, id := range ids {
			setter.SetID(i, id)
		}
	}

	return r.Hits.Total.Value, nil
}

