package route

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

var (
	host  = "192.168.79.2:31359"
	topic = "kafka_es_test_4"
)

type Msg struct {
	Id         string `json:"id"`
	Content    string `json:"content"`
	Date       int64  `json:"date"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	Title      string `json:"title"`
}

func (s Server) Producer(c *gin.Context) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	producer, err := sarama.NewAsyncProducer([]string{host}, config)
	if err != nil {
		fmt.Printf("producer_test create producer error :%s\n", err.Error())
		return
	}

	defer producer.AsyncClose()

	// send message
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder("go_test_1"),
	}
	nowTime := time.Now()
	for i := 0; i < 10000; i++ {
		k := strconv.Itoa(i)
		id := c.Query("id")
		if id == "" {
			id = "1"
		}
		Content := c.Query("Content")
		if Content == "" {
			Content = "this is Content"
		}
		Title := c.Query("Title")
		if Title == "" {
			Title = "this is Title"
		}
		doc := Msg{
			Id:         id + "_" + k,
			Content:    Content + "_" + k,
			Date:       nowTime.UnixNano() / 1e6,
			CreateTime: nowTime.Unix(),
			Title:      Title + "_" + k,
		}
		res, _ := json.Marshal(doc)
		msg.Value = sarama.ByteEncoder(res)
		fmt.Printf("input [%s]\n", &res)
		// send to chain
		producer.Input() <- msg
		select {
		case suc := <-producer.Successes():
			fmt.Printf("offset: %d,  timestamp: %s", suc.Offset, suc.Timestamp.String())
		case fail := <-producer.Errors():
			fmt.Printf("err: %s\n", fail.Err.Error())
		}
	}

	c.JSON(http.StatusOK, "")
}

func (s Server) Consumer(c *gin.Context) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	// consumer
	consumer, err := sarama.NewConsumer([]string{host}, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}

	defer consumer.Close()

	//获取 kafka 主题
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Println("get partitions failed, err:", err)
		return
	}

	for _, p := range partitions {
		partitionConsumer, err := consumer.ConsumePartition(topic, p, sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("get partition consumer failed, error[%v]", err.Error())
			continue
		}
		var (
			count = 0
			rows  = ""
		)
		for message := range partitionConsumer.Messages() {
			count++
			//fmt.Printf("message:[%v], key:[%v], offset:[%v]\n", string(message.Value), string(message.Key), string(message.Offset))
			var getVal Msg
			errMsg := json.Unmarshal(message.Value, &getVal)
			if errMsg != nil {
				fmt.Println("err", errMsg)
			}

			getVal.UpdateTime = time.Now().Unix()
			rows += esBulkData(topic, getVal)
			if count == 100 {
				go CreateEsForKafkaBulk(rows)
				count = 0
				rows = ""
			}
		}
	}
}

func esBulkData(index string, msg Msg) string {
	rowstr := `{ "index" : { "_index" : "` + index + `" }}`
	body, err := json.Marshal(&msg)
	if err != nil {
		return ""
	}
	rowstr += "\n" + string(body) + "\n"
	//fmt.Println("rowstr", rowstr)
	return rowstr
}

/**
新增es Time Filter field name: date
PUT kafka_es_test_4
{
"mappings": {

"properties": {
"date": {
"type":   "date",
"format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
}
}
}

}
*/
