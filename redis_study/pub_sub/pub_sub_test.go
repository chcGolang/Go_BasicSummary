package pub_sub

import (
	"fmt"
	"testing"
)

func TestPublish(t *testing.T) {
	subscribeNum, err := Publish("pub_channl", "消息发布")
	if err != nil {
		panic(err)
	}
	fmt.Println(subscribeNum)
}

func TestSubscribe(t *testing.T) {
	err := Subscribe("pub_channl")
	if err != nil {
		panic(err)
	}
}
