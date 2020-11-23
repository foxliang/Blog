
func producerTest() {
	fmt.Printf("producer_test\n")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Printf("producer_test create producer error :%s\n", err.Error())
		return
	}

	defer producer.AsyncClose()

	// send message
	msg := &sarama.ProducerMessage{
		Topic: "kafka_go_test",
		Key:   sarama.StringEncoder("go_test"),
	}

	value := "this is message"
	for {
		fmt.Scanln(&value)
		msg.Value = sarama.ByteEncoder(value)
		fmt.Printf("input [%s]\n", value)

		// send to chain
		producer.Input() <- msg

		select {
		case suc := <-producer.Successes():
			fmt.Printf("offset: %d,  timestamp: %s", suc.Offset, suc.Timestamp.String())
		case fail := <-producer.Errors():
			fmt.Printf("err: %s\n", fail.Err.Error())
		}
	}
}

/*输出内容
➜  kafka git:(master) ✗ go run main.go
producer_test

input [this is message]
offset: 1,  timestamp: 0001-01-01 00:00:00 +0000 UTC
input [this is message]
offset: 12,  timestamp: 0001-01-01 00:00:00 +0000 UTC
input [this is message]
offset: 13,  timestamp: 0001-01-01 00:00:00 +0000 UTC
input [this is message]
*/
