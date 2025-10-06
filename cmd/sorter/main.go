package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/IBM/sarama"
)

// Record represents one data entry from Kafka
type Record struct {
	ID        int
	Name      string
	Address   string
	Continent string
	Raw       string
}

// Record represents one data entry from Kafka
func parseRecord(line string) Record {
	parts := strings.Split(line, ",")
	id, _ := strconv.Atoi(parts[0])
	return Record{ID: id, Name: parts[1], Address: parts[2], Continent: parts[3], Raw: line}
}

func main() {

	brokers := []string{os.Getenv("KAFKA_ENDPOINT")}    // Kafka broker(s)
	// Topic names for source and sorted outputs
	sourceTopic := "source"
	idTopic := "id"
	nameTopic := "name"
	continentTopic := "continent"

	// Consumer config
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V2_1_0_0

	consumer, err := sarama.NewConsumer(brokers, config)    // create a consumer instance
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	partConsumer, err := consumer.ConsumePartition(sourceTopic, 0, sarama.OffsetOldest)     //consume message from the 0 partition
	if err != nil {
		panic(err)
	}
	defer partConsumer.Close()
    
	var records []Record
	start := time.Now()
	fmt.Println("Consuming records...")

	//main consuming logic
consumeLoop:
	for {
		select {
		case msg := <-partConsumer.Messages():
			records = append(records, parseRecord(string(msg.Value)))
		case err := <-partConsumer.Errors():
			fmt.Println("Consumer error:", err)
		case <-time.After(5 * time.Second): // assume no more messages
			break consumeLoop
		}
	}
	fmt.Println("Consumed", len(records), "records in", time.Since(start))

	produceSorted(records, idTopic, func(i, j int) bool { return records[i].ID < records[j].ID }, brokers)
	produceSorted(records, nameTopic, func(i, j int) bool { return records[i].Name < records[j].Name }, brokers)
	produceSorted(records, continentTopic, func(i, j int) bool { return records[i].Continent < records[j].Continent }, brokers)
}

//again produce sorted message to topic on kafka

func produceSorted(records []Record, topic string, less func(i, j int) bool, brokers []string) {
	sort.Slice(records, less)

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	for _, r := range records {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(r.Raw),
		}
		_, _, err := producer.SendMessage(msg)
        log.Printf("r %v msg %v\n",r,msg)
		if err != nil {
			fmt.Println("Failed to produce message:", err)
		}
	}
	fmt.Println("Produced sorted records to topic", topic)
}
