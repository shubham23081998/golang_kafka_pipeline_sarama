package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/IBM/sarama"
)

var continents = []string{"North America", "Asia", "South America", "Europe", "Africa", "Australia"}
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randomAddress() string {
	return fmt.Sprintf("%d %s %s", rand.Intn(9999), randomString(5), randomString(5))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	brokers := []string{os.Getenv("KAFKA_ENDPOINT")} // Kafka broker(s)
	topic := "source"

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal  // The producer waits for the leader broker to acknowledge the message
	config.Producer.Return.Successes = true             // The producer reports when a message is successfully sent.
	config.Producer.Partitioner = sarama.NewRandomPartitioner  // Randomly sends messages to any partition (helps load balancing).

	producer, err := sarama.NewSyncProducer(brokers, config)    //Creates a synchronous Kafka producer
	if err != nil {
		log.Printf("Err while creating producer instance %v \n", err)
	}
	defer producer.Close()

	total1 := os.Getenv("TOTAL_MESSAGE")
	total,_ := strconv.Atoi(total1)
	start := time.Now()

	for i := 0; i < total; i++ {
		record := fmt.Sprintf("%d,%s,%s,%s",  // create a csv string
			rand.Int31(),
			randomString(rand.Intn(6)+10),
			randomAddress(),
			continents[rand.Intn(len(continents))],
		)
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(record),
		}
		_, _, err := producer.SendMessage(msg)    //send message 
		fmt.Printf("msg %v : %v\n", i, msg)
		if err != nil {
			fmt.Println("Failed to produce message:", err)
		}
	}

	fmt.Println("Produced", total, "records to Kafka topic", topic, "in", time.Since(start))
}
