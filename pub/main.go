package main

import (
	"context"
	"flag"

	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()

	projectid := flag.String("projectid", "", "project id")
	topicName := flag.String("topic", "", "topic name")
	msg := flag.String("msg", "", "message")
	flag.Parse()

	c, err := pubsub.NewClient(ctx, *projectid)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	topic := c.Topic(*topicName)
	exist, err := topic.Exists(ctx)
	if err != nil {
		panic(err)
	}

	if !exist {
		topic, err = c.CreateTopic(ctx, *topicName)
		if err != nil {
			panic(err)
		}
	}

	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(*msg),
	})

	_, err = result.Get(ctx)
	if err != nil {
		panic(err)
	}
}
