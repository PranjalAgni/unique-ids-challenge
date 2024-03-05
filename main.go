package main

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"github.com/lithammer/shortuuid"
)

func main() {
	node := maelstrom.NewNode()
	prefix := shortuuid.New()

	node.Handle("generate", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		now := time.Now()
		nanosec := now.UnixNano()

		uid := prefix + strconv.Itoa(int(nanosec))

		body["type"] = "generate_ok"
		body["id"] = uid
		body["time"] = now
		return node.Reply(msg, body)
	})

	if err := node.Run(); err != nil {
		log.Fatal(err)
	}
}
