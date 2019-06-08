package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

//commit the offset
func CommitMessage(consumer *kafka.Consumer, msg *kafka.Message) {

	//processing is done for this consumer group and commit the offset
	consumer.CommitMessage(msg)
}

//close the consumer
func Close(consumer *kafka.Consumer, msg *kafka.Message) {

	//processing is done for this consumer group and commit the offset
	consumer.CommitMessage(msg)
}

//subscribe to the topic
//it returns the message and consumer
func createConsumer(topic string) *kafka.Consumer {
	//create at least consumer
	//enable.auto.commit is true and sync processing and auto.commit.interval.ms = 10000 -> atleast once/atmost once
	//enable.auto.commit is true and async processing and auto.commit.interval.ms = 0 -> atmost once
	//enable.auto.commit is false -> atleast once

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  "localhost",
		"group.id":           "article_scraper_processor",
		"auto.offset.reset":  "latest",
		"enable.auto.commit": "false",
	})

	//raise the panic in case of error
	if err != nil {
		panic(err)
	}

	consumer.SubscribeTopics([]string{topic}, nil)

	/*//poll to the topic and consume the message
	msg, err := consumer.ReadMessage(-1)
	*/
	return consumer
}
