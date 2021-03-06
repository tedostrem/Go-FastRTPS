// +build publisher

package main

// #include "Publisher.h"
// #cgo CXXFLAGS: -std=c++11
// #cgo LDFLAGS: -lfastcdr -lfastrtps -lgofastrtps
import "C"

import (
	"crypto/md5"
	"fmt"
	"go-fastrtps/fastrtps"
	"math/rand"
	"time"
)

func image() []byte {
	rand.Seed(time.Now().UnixNano())
	buffer := make([]byte, 1382400)
	rand.Read(buffer)
	return buffer
}

func main() {
	participant := fastrtps.NewParticipant("participant_publisher")
	topicDataType := fastrtps.NewMediaTopicDataType()
	publisherAttributes := fastrtps.GetAttributes(topicDataType, "topic_multimedia").MultimediaPublisher
	fastrtps.RegisterType(participant, topicDataType)
	publisher := fastrtps.NewPublisher(participant, publisherAttributes)
	img := image()
	for i := 0; i < 10; i += 1 {
		fmt.Printf("Publisher Golang: %x\n", md5.Sum(img))
		C.PublishMedia(publisher, (*C.char)(C.CBytes(img)))
	}
	fastrtps.RemoveParticipant(participant)
}
