
func consumerTest() {
	fmt.Printf("consumer_test")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	// consumer
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}

	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("kafka_es_test", 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
		case err := <-partitionConsumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}

}


/*接收内容
➜  kafka git:(master) ✗ go run consumer.go
consumer_testmsg offset: 0, partition: 0, timestamp: 2020-11-23 11:30:37.356 +0800 CST, value: this is message
msg offset: 1, partition: 0, timestamp: 2020-11-23 11:30:39.309 +0800 CST, value: this is message
msg offset: 2, partition: 0, timestamp: 2020-11-23 11:30:40.591 +0800 CST, value: this is message
msg offset: 3, partition: 0, timestamp: 2020-11-23 11:30:41.382 +0800 CST, value: this is message
*/
