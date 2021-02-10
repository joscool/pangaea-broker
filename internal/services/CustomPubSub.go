package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/custom-broker/internal/store"
)

type CustomPubsub struct {
	s store.Broker
}

func NewCustomPubsub() *CustomPubsub {
	return &CustomPubsub{
		s: store.NewBroker(),
	}
}

func (c *CustomPubsub) Subscribe(topic, url string) error {

	c.s.Subscribe(topic, url)
	return nil
}
func (c *CustomPubsub) Publish(topic, message string) error {

	subscriber := c.s.GetTopicSubscribers(topic)
	subscribers := subscriber.GetAll()

	for url, _ := range subscribers {
		forwardData(topic, message, url)
	}
	return nil
}

func forwardData(topic, message, url string) {

	postBody, _ := json.Marshal(map[string]string{
		"topic":   topic,
		"message": message,
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(url, "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

}
