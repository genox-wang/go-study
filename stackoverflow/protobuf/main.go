package main

import (
	"fmt"
	msg "go-study/stackoverflow/protobuf/proto_test/proto_msg"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
)

func main() {
	m := &msg.Message{}
	MessageData := `{"jobId": "abc", "sentOnce": true}`

	err := protojson.Unmarshal([]byte(MessageData), m)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(m.String())
}
