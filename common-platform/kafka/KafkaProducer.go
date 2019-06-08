// Example function-based Apache Kafka producer
package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

//This send message is blocking api and need to be updated with timeout
//In case of failure return error otherwise nil
func SendMessage(topicName string, key string, msg []byte) error {
	//create a producer instance
	//Most important property is kafka server address
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092", "acks": -1})

	//check for errors
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Producer %v\n", p)

	// Optional delivery channel, if not specified the Producer object's
	// .Events channel is used.
	deliveryChan := make(chan kafka.Event)

	//send a message
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topicName},
		Key:            []byte(key),
		Value:          msg,
	}, deliveryChan)

	report := <-deliveryChan
	m := report.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	//close channel after sending the message
	defer func(deliveryChan chan kafka.Event) {
		close(deliveryChan)
	}(deliveryChan)

	return m.TopicPartition.Error
}
